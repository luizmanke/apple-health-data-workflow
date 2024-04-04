package controller

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type Storage struct {
	Directory string
}

func ListCSVFiles(storage Storage) ([]string, error) {
	fileNames := []string{}

	directory := storage.Directory
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

func ReadAppleHealthSummaryDataFromCSVFile(storage Storage, fileName string) ([]Summary, error) {

	filePath := filepath.Join(storage.Directory, fileName)
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

func convertCSVRowsToSummaryStructs(csvRows [][]string) ([]Summary, error) {

	headers := csvRows[0]
	rows := csvRows[1:]
	summaries := []Summary{}

	for _, row := range rows {

		rowMap := map[string]string{}
		for i := range headers {
			rowMap[headers[i]] = row[i]
		}

		summary := Summary{}
		decoderConfig := &mapstructure.DecoderConfig{Result: &summary, WeaklyTypedInput: true}
		decoder, err := mapstructure.NewDecoder(decoderConfig)
		if err != nil {
			return nil, err
		}

		decoder.Decode(rowMap)
		summaries = append(summaries, summary)
	}

	return summaries, nil
}
