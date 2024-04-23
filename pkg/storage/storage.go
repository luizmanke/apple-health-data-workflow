package storage

import (
	"apple-health-data-workflow/pkg/models"
	"encoding/csv"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
)

type StorageConfig struct {
	Directory string
}

func ListCSVFiles(storageConfig StorageConfig) ([]string, error) {
	fileNames := []string{}

	directory := storageConfig.Directory
	if !strings.HasSuffix(directory, "/") {
		directory = directory + "/"
	}

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".csv") {
			fileName := strings.Replace(path, directory, "", 1)
			fileNames = append(fileNames, fileName)
		}

		return nil
	})

	return fileNames, err
}

func ReadAppleHealthSummaryDataFromCSVFile(storageConfig StorageConfig, fileName string) ([]models.Summary, error) {

	filePath := filepath.Join(storageConfig.Directory, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = -1
	csvRows, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	summaries, err := convertCSVRowsToSummaryStructs(csvRows)
	if err != nil {
		return nil, err
	}

	return summaries, nil
}

func convertCSVRowsToSummaryStructs(csvRows [][]string) ([]models.Summary, error) {

	headers := csvRows[0]
	rows := csvRows[1:]
	summaries := []models.Summary{}

	for _, row := range rows {

		rowMap := map[string]string{}
		for i := range headers {
			rowMap[headers[i]] = row[i]
		}

		summary := models.Summary{}
		decoderConfig := &mapstructure.DecoderConfig{Result: &summary, WeaklyTypedInput: true}
		decoder, err := mapstructure.NewDecoder(decoderConfig)
		if err != nil {
			return nil, err
		}

		decoder.Decode(rowMap)
		err = convertTimestampFormat(&summary)
		if err != nil {
			return nil, err
		}

		summaries = append(summaries, summary)
	}

	return summaries, nil
}

func convertTimestampFormat(summary *models.Summary) error {

	inputFormat := "1/2/2006 15:04:05"
	outputFormat := "2006-01-02T15:04:05Z"

	parsedTime, err := time.Parse(inputFormat, summary.Date)
	if err != nil {
		return err
	}

	summary.Date = parsedTime.UTC().Format(outputFormat)
	return nil
}
