package main

import (
	"apple-health-data-workflow/internal/backend"
	"apple-health-data-workflow/pkg/database"
	"apple-health-data-workflow/pkg/queue"
	"apple-health-data-workflow/pkg/storage"
	"os"
)

func main() {

	storageConfig := storage.StorageConfig{
		Directory: os.Getenv("INGESTER_SOURCE_DIRECTORY"),
	}
	queueConfig := queue.QueueConfig{
		Server: os.Getenv("QUEUE_SERVER"),
		Topic:  os.Getenv("QUEUE_TOPIC"),
	}
	databaseConfig := database.DatabaseConfig{
		User:     os.Getenv("INGESTER_DESTINATION_USER"),
		Password: os.Getenv("INGESTER_DESTINATION_PASSWORD"),
		Host:     os.Getenv("INGESTER_DESTINATION_HOST"),
		Port:     os.Getenv("INGESTER_DESTINATION_PORT"),
		Database: os.Getenv("INGESTER_DESTINATION_DATABASE"),
	}

	backend.ReadAppleHealthSummaryDataFromCSVFilesAndSendToQueue(storageConfig, queueConfig)
	backend.ReadAppleHealthSummaryDataFromQueueAndIngestIntoDatabase(queueConfig, databaseConfig)
}
