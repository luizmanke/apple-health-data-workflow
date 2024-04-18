package queue

import (
	"apple-health-data-workflow/internal/controller"
	"encoding/json"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type QueueConfig struct {
	Server string
	Topic  string
}

func SendAppleHealthSummaryDataToQueue(
	queueConfig QueueConfig,
	summaries []controller.Summary,
) error {

	producer, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": queueConfig.Server,
		},
	)
	if err != nil {
		return err
	}
	defer producer.Close()

	err = publishSummaryMessages(producer, queueConfig.Topic, summaries)
	if err != nil {
		return err
	}

	return nil
}

func ReadAppleHealthSummaryDataFromQueue(
	queueConfig QueueConfig,
) ([]controller.Summary, error) {

	consumer, err := kafka.NewConsumer(
		&kafka.ConfigMap{
			"bootstrap.servers": queueConfig.Server,
			"group.id":          "apple-health.summary",
			"auto.offset.reset": "earliest",
		},
	)
	if err != nil {
		return nil, err
	}
	defer consumer.Close()

	err = consumer.SubscribeTopics([]string{queueConfig.Topic}, nil)
	if err != nil {
		return nil, err
	}

	// Wait consumer assignment
	consumer.Poll(10 * 1000)
	time.Sleep(1 * time.Second)

	summaries, err := readSummaryMessages(consumer)
	if err != nil {
		return nil, err
	}

	return summaries, nil
}

func publishSummaryMessages(kafkaProducer *kafka.Producer, topic string, summaries []controller.Summary) error {

	deliveryChannel := make(chan kafka.Event, 1)

	for _, summary := range summaries {

		topicPartition := kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		}

		summarySerialized, err := json.Marshal(summary)
		if err != nil {
			return err
		}

		message := kafka.Message{
			TopicPartition: topicPartition,
			Value:          summarySerialized,
		}

		kafkaProducer.Produce(&message, deliveryChannel)

		event := <-deliveryChannel
		err = event.(*kafka.Message).TopicPartition.Error
		if err != nil {
			log.Printf("Failed to publish message: %v\n", err)
		}
	}

	close(deliveryChannel)

	return nil
}

func readSummaryMessages(consumer *kafka.Consumer) ([]controller.Summary, error) {

	numberOfAvailableMessages, err := getNumberOfAvailableMessages(consumer)
	if err != nil {
		return nil, err
	}

	summaries := []controller.Summary{}
	for i := 0; i < numberOfAvailableMessages; i++ {
		message, err := consumer.ReadMessage(100 * time.Millisecond)
		if err == nil {

			summary := controller.Summary{}
			err := json.Unmarshal(message.Value, &summary)
			if err != nil {
				return nil, err
			}

			summaries = append(summaries, summary)

		} else if !err.(kafka.Error).IsTimeout() {
			return nil, err
		}
	}

	return summaries, nil
}

func getNumberOfAvailableMessages(consumer *kafka.Consumer) (int, error) {

	topicPartitions, err := consumer.Assignment()
	if err != nil {
		return 0, err
	}

	numberOfAvailableMessages := 0
	for _, partition := range topicPartitions {

		low, high, err := consumer.GetWatermarkOffsets(*partition.Topic, partition.Partition)
		if err != nil {
			return 0, err
		}

		numberOfAvailableMessages += int(high - low)
	}

	return numberOfAvailableMessages, nil
}
