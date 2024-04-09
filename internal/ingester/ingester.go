package ingester

import (
	"apple-health-data-workflow/internal/controller"
	"log"
)

func IngestAppleHealthSummaryData(
	sourceStorage controller.Storage,
	destinationDatabase controller.Database,
) {
	fileNames, err := controller.ListCSVFiles(sourceStorage)
	if err != nil {
		panic(err)
	}

	for _, fileName := range fileNames {
		log.Printf("Ingesting file '%s'", fileName)

		summaries, err := controller.ReadAppleHealthSummaryDataFromCSVFile(sourceStorage, fileName)
		if err != nil {
			panic(err)
		}

		err = controller.InsertAppleHealthSummaryDataIntoDatabase(destinationDatabase, summaries)
		if err != nil {
			panic(err)
		}
	}
}
