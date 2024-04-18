package ingester

import (
	"apple-health-data-workflow/internal/controller"
	"apple-health-data-workflow/internal/queue"
)

func ReadAppleHealthSummaryDataFromQueueAndIngestIntoDatabase(
	queueConfig queue.QueueConfig,
	destinationDatabase controller.Database,
) {

	summaries, err := queue.ReadAppleHealthSummaryDataFromQueue(queueConfig)
	if err != nil {
		panic(err)
	}

	err = controller.InsertAppleHealthSummaryDataIntoDatabase(destinationDatabase, summaries)
	if err != nil {
		panic(err)
	}
}
