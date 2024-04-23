package frontend

import (
	"apple-health-data-workflow/pkg/database"
	"apple-health-data-workflow/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
	"runtime"
	"text/template"
)

type DropdownOption struct {
	Text string
	X    []string
	Y    []float32
}

func DisplaySummaryChart(w http.ResponseWriter, databaseConfig database.DatabaseConfig) {

	summaries, err := database.GetSummaryDataFromDatabase(databaseConfig)
	if err != nil {
		message := fmt.Sprintf("Failed to get summary data from database: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	dropdownOptions := createDropdownOptions(summaries)
	data, err := createTemplateData(dropdownOptions)
	if err != nil {
		message := fmt.Sprintf("Failed to create template data: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	err = executeTemplate(w, data)
	if err != nil {
		message := fmt.Sprintf("Failed to execute template: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}
}

func createDropdownOptions(summaries []models.Summary) []DropdownOption {

	if len(summaries) == 0 {
		return []DropdownOption{}
	}

	xValues := []string{}
	for _, summary := range summaries {
		xValues = append(xValues, summary.Date)
	}

	dropdownOptions := []DropdownOption{}
	summaryReflectValue := reflect.ValueOf(summaries[0])
	summaryReflectType := summaryReflectValue.Type()
	for i := 0; i < summaryReflectValue.NumField(); i++ {

		fieldName := summaryReflectType.Field(i).Name
		if fieldName == "Date" {
			continue
		}

		yValues := []float32{}
		for _, summary := range summaries {
			summaryReflectValue := reflect.ValueOf(summary)
			summaryReflectField := summaryReflectValue.FieldByName(fieldName)
			yValues = append(yValues, summaryReflectField.Interface().(float32))
		}

		dropdownOptions = append(
			dropdownOptions,
			DropdownOption{
				Text: fieldName,
				X:    xValues,
				Y:    yValues,
			},
		)
	}

	return dropdownOptions
}

func createTemplateData(dropdownOptions []DropdownOption) (map[string]interface{}, error) {

	dropdownOptionsJSON, err := json.Marshal(dropdownOptions)
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"dropdownOptions": string(dropdownOptionsJSON),
	}

	return data, nil
}

func executeTemplate(w http.ResponseWriter, data map[string]interface{}) error {

	_, currentFilePath, _, _ := runtime.Caller(0)
	currentFileDir := filepath.Dir(currentFilePath)

	tmpl, err := template.ParseFiles(currentFileDir + "/summary-chart.html")
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}
