package backend_test

import (
	"apple-health-data-workflow/internal/backend"
	"apple-health-data-workflow/pkg/models"
	"apple-health-data-workflow/pkg/queue"
	"apple-health-data-workflow/pkg/storage"
	"apple-health-data-workflow/pkg/testkit"
	"encoding/csv"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/lib/pq"
)

type csvFile struct {
	name string
	data []map[string]string
}

type storageToQueueFixtures struct {
	storageConfig storage.StorageConfig
	queueConfig   queue.QueueConfig
}

func TestEveryAppleHealthSummaryColumnMustBeSentToQueue(t *testing.T) {

	fixtures := storageToQueueSetUp(
		t,
		[]csvFile{
			appleHealthSummaryFileWithAllColumns(),
		},
	)

	backend.ReadAppleHealthSummaryDataFromCSVFilesAndSendToQueue(
		fixtures.storageConfig,
		fixtures.queueConfig,
	)

	testkit.AssertEqual(
		t,
		getSummaryMessagesFromQueue(t, fixtures.queueConfig),
		[]models.Summary{appleHealthSummaryDataWithAllColumns()},
	)
}

func TestMultipleAppleHealthSummaryFilesMustBeSentToQueue(t *testing.T) {

	fixtures := storageToQueueSetUp(
		t,
		[]csvFile{
			{
				name: "file1.csv",
				data: []map[string]string{
					{"Date": "1/2/2024 0:00:00"},
					{"Date": "1/3/2024 0:00:00"},
				},
			},
			{
				name: "file2.csv",
				data: []map[string]string{
					{"Date": "1/4/2024 0:00:00"},
				},
			},
		},
	)

	backend.ReadAppleHealthSummaryDataFromCSVFilesAndSendToQueue(
		fixtures.storageConfig,
		fixtures.queueConfig,
	)

	testkit.AssertEqual(
		t,
		getSummaryMessagesFromQueue(t, fixtures.queueConfig),
		[]models.Summary{
			{Date: "2024-01-02T00:00:00Z"},
			{Date: "2024-01-03T00:00:00Z"},
			{Date: "2024-01-04T00:00:00Z"},
		},
	)
}

func storageToQueueSetUp(t *testing.T, csvFiles []csvFile) storageToQueueFixtures {

	testDir := createTestDirectory(t)
	createCSVFilesInStorage(t, testDir, csvFiles)

	storageConfig := storage.StorageConfig{
		Directory: testDir,
	}
	queueConfig := queue.QueueConfig{
		Server: "datahub:9092",
		Topic:  "reader-test.summary",
	}

	return storageToQueueFixtures{
		storageConfig,
		queueConfig,
	}
}

func createTestDirectory(t *testing.T) string {
	testDir, err := os.MkdirTemp("", "")
	testkit.AssertNoError(t, err)
	return testDir
}

func createCSVFilesInStorage(t *testing.T, directory string, csvFiles []csvFile) {
	for _, file := range csvFiles {

		filePath := filepath.Join(directory, file.name)
		fileObject, err := os.Create(filePath)
		testkit.AssertNoError(t, err)
		defer fileObject.Close()

		csvWriter := csv.NewWriter(fileObject)
		defer csvWriter.Flush()

		data := convertListOfMapsToNestedLists(file.data)
		csvWriter.WriteAll(data)
		testkit.AssertNoError(t, err)
	}
}

func convertListOfMapsToNestedLists(listOfMaps []map[string]string) [][]string {

	headers := []string{}
	for key := range listOfMaps[0] {
		headers = append(headers, key)
	}

	nestedLists := [][]string{headers}
	for _, item := range listOfMaps {

		newList := []string{}
		for _, key := range headers {
			newList = append(newList, item[key])
		}

		nestedLists = append(nestedLists, newList)
	}

	return nestedLists
}

func appleHealthSummaryFileWithAllColumns() csvFile {
	return csvFile{
		name: "file.csv",
		data: []map[string]string{
			{
				"Date":                                             "1/1/2024 0:00:00",
				"Active Energy (kJ)":                               "1.0",
				"Alcohol Consumption (count)":                      "1.0",
				"Apple Exercise Time (min)":                        "1.0",
				"Apple Move Time (min)":                            "1.0",
				"Apple Sleeping Wrist Temperature (ºC)":            "1.0",
				"Apple Stand Hour (hours)":                         "1.0",
				"Apple Stand Time (min)":                           "1.0",
				"Atrial Fibrillation Burden (%)":                   "1.0",
				"Basal Body Temperature (ºC)":                      "1.0",
				"Blood Alcohol Content (%)":                        "1.0",
				"Blood Glucose (mmol/L)":                           "1.0",
				"Blood Oxygen Saturation (%)":                      "1.0",
				"Blood Pressure [Systolic] (mmHg)":                 "1.0",
				"Blood Pressure [Diastolic] (mmHg)":                "1.0",
				"Body Fat Percentage (%)":                          "1.0",
				"Body Mass Index (count)":                          "1.0",
				"Body Temperature (ºC)":                            "1.0",
				"Caffeine (mg)":                                    "1.0",
				"Calcium (mg)":                                     "1.0",
				"Carbohydrates (g)":                                "1.0",
				"Cardio Recovery (count/min)":                      "1.0",
				"Chloride (mg)":                                    "1.0",
				"Chromium (mcg)":                                   "1.0",
				"Copper (mg)":                                      "1.0",
				"Cycling Cadence (count/min)":                      "1.0",
				"Cycling Distance (km)":                            "1.0",
				"Cycling Functional Threshold Power (watts)":       "1.0",
				"Cycling Power (watts)":                            "1.0",
				"Cycling Speed (km/hr)":                            "1.0",
				"Dietary Biotin (mcg)":                             "1.0",
				"Dietary Cholesterol (mg)":                         "1.0",
				"Dietary Energy (kJ)":                              "1.0",
				"Dietary Sugar (g)":                                "1.0",
				"Distance Downhill Snow Sports (km)":               "1.0",
				"Electrodermal Activity (S)":                       "1.0",
				"Environmental Audio Exposure (dBASPL)":            "1.0",
				"Fiber (g)":                                        "1.0",
				"Flights Climbed (count)":                          "1.0",
				"Folate (mcg)":                                     "1.0",
				"Forced Expiratory Volume 1 (L)":                   "1.0",
				"Forced Vital Capacity (L)":                        "1.0",
				"Handwashing (s)":                                  "1.0",
				"Headphone Audio Exposure (dBASPL)":                "1.0",
				"Heart Rate [Min] (bpm)":                           "1.0",
				"Heart Rate [Max] (bpm)":                           "1.0",
				"Heart Rate [Avg] (bpm)":                           "1.0",
				"Heart Rate Variability (ms)":                      "1.0",
				"Height (m)":                                       "1.0",
				"High Heart Rate Notifications [Min] (count)":      "1.0",
				"High Heart Rate Notifications [Max] (count)":      "1.0",
				"High Heart Rate Notifications [Avg] (count)":      "1.0",
				"Inhaler Usage (count)":                            "1.0",
				"Insulin Delivery (IU)":                            "1.0",
				"Iodine (mcg)":                                     "1.0",
				"Iron (mg)":                                        "1.0",
				"Irregular Heart Rate Notifications [Min] (count)": "1.0",
				"Irregular Heart Rate Notifications [Max] (count)": "1.0",
				"Irregular Heart Rate Notifications [Avg] (count)": "1.0",
				"Lean Body Mass (kg)":                              "1.0",
				"Low Heart Rate Notifications [Min] (count)":       "1.0",
				"Low Heart Rate Notifications [Max] (count)":       "1.0",
				"Low Heart Rate Notifications [Avg] (count)":       "1.0",
				"Magnesium (mg)":                                   "1.0",
				"Manganese (mg)":                                   "1.0",
				"Mindful Minutes (min)":                            "1.0",
				"Molybdenum (mcg)":                                 "1.0",
				"Monounsaturated Fat (g)":                          "1.0",
				"Niacin (mg)":                                      "1.0",
				"Number of Times Fallen (falls)":                   "1.0",
				"Pantothenic Acid (mg)":                            "1.0",
				"Peak Expiratory Flow Rate (L/min)":                "1.0",
				"Peripheral Perfusion Index (%)":                   "1.0",
				"Physical Effort (MET)":                            "1.0",
				"Polyunsaturated Fat (g)":                          "1.0",
				"Potassium (mg)":                                   "1.0",
				"Protein (g)":                                      "1.0",
				"Push Count (count)":                               "1.0",
				"Respiratory Rate (count/min)":                     "1.0",
				"Resting Energy (kJ)":                              "1.0",
				"Resting Heart Rate (bpm)":                         "1.0",
				"Riboflavin (mg)":                                  "1.0",
				"Running Ground Contact Time (ms)":                 "1.0",
				"Running Power (watts)":                            "1.0",
				"Running Speed (km/hr)":                            "1.0",
				"Running Stride Length (m)":                        "1.0",
				"Running Vertical Oscillation (cm)":                "1.0",
				"Saturated Fat (g)":                                "1.0",
				"Selenium (mcg)":                                   "1.0",
				"Sexual Activity [Unspecified] (times)":            "1.0",
				"Sexual Activity [Protection Used] (times)":        "1.0",
				"Sexual Activity [Protection Not Used] (times)":    "1.0",
				"Six-Minute Walking Test Distance (m)":             "1.0",
				"Sleep Analysis [Asleep] (hr)":                     "1.0",
				"Sleep Analysis [In Bed] (hr)":                     "1.0",
				"Sleep Analysis [Core] (hr)":                       "1.0",
				"Sleep Analysis [Deep] (hr)":                       "1.0",
				"Sleep Analysis [REM] (hr)":                        "1.0",
				"Sleep Analysis [Awake] (hr)":                      "1.0",
				"Sodium (mg)":                                      "1.0",
				"Stair Speed: Down (m/s)":                          "1.0",
				"Stair Speed: Up (m/s)":                            "1.0",
				"Step Count (steps)":                               "1.0",
				"Swimming Distance (m)":                            "1.0",
				"Swimming Stroke Count (count)":                    "1.0",
				"Thiamin (mg)":                                     "1.0",
				"Time in Daylight (min)":                           "1.0",
				"Toothbrushing (s)":                                "1.0",
				"Total Fat (g)":                                    "1.0",
				"UV Exposure (count)":                              "1.0",
				"Underwater Depth (m)":                             "1.0",
				"Underwater Temperature (ºC)":                      "1.0",
				"VO2 Max (ml/(kg·min))":                            "1.0",
				"Vitamin A (mcg)":                                  "1.0",
				"Vitamin B12 (mcg)":                                "1.0",
				"Vitamin B6 (mg)":                                  "1.0",
				"Vitamin C (mg)":                                   "1.0",
				"Vitamin D (mcg)":                                  "1.0",
				"Vitamin E (mg)":                                   "1.0",
				"Vitamin K (mcg)":                                  "1.0",
				"Waist Circumference (cm)":                         "1.0",
				"Walking + Running Distance (km)":                  "1.0",
				"Walking Asymmetry Percentage (%)":                 "1.0",
				"Walking Double Support Percentage (%)":            "1.0",
				"Walking Heart Rate Average (bpm)":                 "1.0",
				"Walking Speed (km/hr)":                            "1.0",
				"Walking Step Length (cm)":                         "1.0",
				"Water (mL)":                                       "1.0",
				"Weight/Body Mass (kg)":                            "1.0",
				"Wheelchair Distance (km)":                         "1.0",
				"Zinc (mg)":                                        "1.0",
			},
		},
	}
}

func getSummaryMessagesFromQueue(t *testing.T, queueConfig queue.QueueConfig) []models.Summary {
	summaries, err := queue.ReadAppleHealthSummaryDataFromQueue(queueConfig)
	testkit.AssertNoError(t, err)
	return summaries
}
