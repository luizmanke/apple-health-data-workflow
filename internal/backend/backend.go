package backend

import (
	"apple-health-data-workflow/pkg/database"
	"apple-health-data-workflow/pkg/queue"
	"apple-health-data-workflow/pkg/storage"
)

func ReadAppleHealthSummaryDataFromCSVFilesAndSendToQueue(
	storageConfig storage.StorageConfig,
	queueConfig queue.QueueConfig,
) {

	fileNames, err := storage.ListCSVFiles(storageConfig)
	if err != nil {
		panic(err)
	}

	for _, fileName := range fileNames {

		summaries, err := storage.ReadAppleHealthSummaryDataFromCSVFile(storageConfig, fileName)
		if err != nil {
			panic(err)
		}

		err = queue.SendAppleHealthSummaryDataToQueue(queueConfig, summaries)
		if err != nil {
			panic(err)
		}
	}
}

func ReadAppleHealthSummaryDataFromQueueAndIngestIntoDatabase(
	queueConfig queue.QueueConfig,
	databaseConfig database.DatabaseConfig,
) {

	summaries, err := queue.ReadAppleHealthSummaryDataFromQueue(queueConfig)
	if err != nil {
		panic(err)
	}

	err = database.InsertAppleHealthSummaryDataIntoDatabase(databaseConfig, summaries)
	if err != nil {
		panic(err)
	}
}
