package main

import (
	"apple-health-data-workflow/internal/controller"
	"apple-health-data-workflow/internal/ingester"
	"apple-health-data-workflow/internal/queue"
	"apple-health-data-workflow/internal/reader"
	"os"
)

func main() {

	sourceStorage := controller.Storage{
		Directory: os.Getenv("INGESTER_SOURCE_DIRECTORY"),
	}
	queueConfig := queue.QueueConfig{
		Server: os.Getenv("QUEUE_SERVER"),
		Topic:  os.Getenv("QUEUE_TOPIC"),
	}
	destinationDatabase := controller.Database{
		User:     os.Getenv("INGESTER_DESTINATION_USER"),
		Password: os.Getenv("INGESTER_DESTINATION_PASSWORD"),
		Host:     os.Getenv("INGESTER_DESTINATION_HOST"),
		Port:     os.Getenv("INGESTER_DESTINATION_PORT"),
		Database: os.Getenv("INGESTER_DESTINATION_DATABASE"),
	}

	reader.ReadAppleHealthSummaryDataFromCSVFilesAndSendToQueue(sourceStorage, queueConfig)
	ingester.ReadAppleHealthSummaryDataFromQueueAndIngestIntoDatabase(queueConfig, destinationDatabase)
}
