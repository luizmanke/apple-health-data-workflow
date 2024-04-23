package main

import (
	"apple-health-data-workflow/internal/frontend"
	"apple-health-data-workflow/pkg/database"
	"log"
	"net/http"
	"os"
)

func main() {

	databaseConfig := database.DatabaseConfig{
		User:     os.Getenv("WAREHOUSE_USER"),
		Password: os.Getenv("WAREHOUSE_PASSWORD"),
		Host:     os.Getenv("WAREHOUSE_HOST"),
		Port:     os.Getenv("WAREHOUSE_PORT"),
		Database: os.Getenv("WAREHOUSE_DATABASE"),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		frontend.DisplaySummaryChart(w, databaseConfig)
	})

	addr := ":8080"
	log.Printf("Starting web application on %s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
