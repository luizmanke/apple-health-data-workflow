package main

import (
	"apple-health-data-workflow/internal/ingester"
	"fmt"
	"os"
)

func main() {
	fmt.Println("App started!")

	sourceStorage := ingester.SourceStorage{
		Directory: os.Getenv("INGESTER_SOURCE_DIRECTORY"),
	}
	destinationDatabase := ingester.DestinationDatabase{
		User:     os.Getenv("INGESTER_DESTINATION_USER"),
		Password: os.Getenv("INGESTER_DESTINATION_PASSWORD"),
		Host:     os.Getenv("INGESTER_DESTINATION_HOST"),
		Port:     os.Getenv("INGESTER_DESTINATION_PORT"),
		Database: os.Getenv("INGESTER_DESTINATION_DATABASE"),
	}

	fmt.Println("Ingesting Apple Health summary data...")
	ingester.IngestAppleHealthSummaryData(sourceStorage, destinationDatabase)
}
