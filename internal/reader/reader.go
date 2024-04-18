package reader

import (
	"apple-health-data-workflow/internal/controller"
	"apple-health-data-workflow/internal/queue"
)

func ReadAppleHealthSummaryDataFromCSVFilesAndSendToQueue(
	sourceStorage controller.Storage,
	queueConfig queue.QueueConfig,
) {

	fileNames, err := controller.ListCSVFiles(sourceStorage)
	if err != nil {
		panic(err)
	}

	for _, fileName := range fileNames {

		summaries, err := controller.ReadAppleHealthSummaryDataFromCSVFile(sourceStorage, fileName)
		if err != nil {
			panic(err)
		}

		err = queue.SendAppleHealthSummaryDataToQueue(queueConfig, summaries)
		if err != nil {
			panic(err)
		}
	}
}
