package ingester_test

import (
	"apple-health-data-workflow/internal/controller"
	"apple-health-data-workflow/internal/ingester"
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

type fixtures struct {
	sourceStorage       controller.Storage
	destinationDatabase controller.Database
}

func TestEveryAppleHealthSummaryColumnMustBeIngestedIntoTheDatabase(t *testing.T) {

	fixtures := setUpTest(
		t,
		[]csvFile{
			appleHealthSummaryFileWithAllColumns(),
		},
	)

	ingester.IngestAppleHealthSummaryData(
		fixtures.sourceStorage,
		fixtures.destinationDatabase,
	)

	ingestedData := getSummaryDataFromDatabaseByDates(t, []string{"2024-01-01T00:00:00Z"})
	testkit.AssertEqual(t, ingestedData, []controller.Summary{appleHealthSummaryDataWithAllColumns()})
}

func TestMultipleAppleHealthSummaryFilesMustBeIngestedIntoTheDatabase(t *testing.T) {

	fixtures := setUpTest(
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

	ingester.IngestAppleHealthSummaryData(
		fixtures.sourceStorage,
		fixtures.destinationDatabase,
	)

	ingestedData := getSummaryDataFromDatabaseByDates(
		t,
		[]string{
			"2024-01-02T00:00:00Z",
			"2024-01-03T00:00:00Z",
			"2024-01-04T00:00:00Z",
		},
	)
	testkit.AssertEqual(
		t,
		ingestedData,
		[]controller.Summary{
			{Date: "2024-01-02T00:00:00Z"},
			{Date: "2024-01-03T00:00:00Z"},
			{Date: "2024-01-04T00:00:00Z"},
		},
	)
}

func TestOnlyTheFirstAppleHealthSummaryWithTheSameDateMustBePersistedInTheDatabase(t *testing.T) {

	fixtures := setUpTest(
		t,
		[]csvFile{
			{
				name: "file1.csv",
				data: []map[string]string{
					{"Date": "1/6/2024 0:00:00", "Active Energy (kJ)": "1.0"},
					{"Date": "1/6/2024 0:00:00", "Active Energy (kJ)": "2.0"},
				},
			},
			{
				name: "file2.csv",
				data: []map[string]string{
					{"Date": "1/6/2024 0:00:00", "Active Energy (kJ)": "3.0"},
				},
			},
		},
	)

	ingester.IngestAppleHealthSummaryData(
		fixtures.sourceStorage,
		fixtures.destinationDatabase,
	)

	ingestedData := getSummaryDataFromDatabaseByDates(t, []string{"2024-01-06T00:00:00Z"})
	testkit.AssertEqual(
		t,
		ingestedData,
		[]controller.Summary{
			{Date: "2024-01-06T00:00:00Z", ActiveEnergy: 1.0},
		},
	)
}

func setUpTest(t *testing.T, csvFiles []csvFile) fixtures {

	testDir := createTestDirectory(t)
	createCSVFilesInStorage(t, testDir, csvFiles)

	sourceStorage := controller.Storage{
		Directory: testDir,
	}
	destinationDatabase := controller.Database{
		User:     "username",
		Password: "password",
		Host:     "warehouse",
		Port:     "5432",
		Database: "appleHealth",
	}

	return fixtures{
		sourceStorage,
		destinationDatabase,
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

func appleHealthSummaryDataWithAllColumns() controller.Summary {
	return controller.Summary{
		Date:                               "2024-01-01T00:00:00Z",
		ActiveEnergy:                       1.0,
		AlcoholConsumption:                 1.0,
		AppleExerciseTime:                  1.0,
		AppleMoveTime:                      1.0,
		AppleSleepingWristTemperature:      1.0,
		AppleStandHour:                     1.0,
		AppleStandTime:                     1.0,
		AtrialFibrillationBurden:           1.0,
		BasalBodyTemperature:               1.0,
		BloodAlcoholContent:                1.0,
		BloodGlucose:                       1.0,
		BloodOxygenSaturation:              1.0,
		BloodPressureSystolic:              1.0,
		BloodPressureDiastolic:             1.0,
		BodyFatPercentage:                  1.0,
		BodyMassIndex:                      1.0,
		BodyTemperature:                    1.0,
		Caffeine:                           1.0,
		Calcium:                            1.0,
		Carbohydrates:                      1.0,
		CardioRecovery:                     1.0,
		Chloride:                           1.0,
		Chromium:                           1.0,
		Copper:                             1.0,
		CyclingCadence:                     1.0,
		CyclingDistance:                    1.0,
		CyclingFunctionalThresholdPower:    1.0,
		CyclingPower:                       1.0,
		CyclingSpeed:                       1.0,
		DietaryBiotin:                      1.0,
		DietaryCholesterol:                 1.0,
		DietaryEnergy:                      1.0,
		DietarySugar:                       1.0,
		DistanceDownhillSnowSports:         1.0,
		ElectrodermalActivity:              1.0,
		EnvironmentalAudioExposure:         1.0,
		Fiber:                              1.0,
		FlightsClimbed:                     1.0,
		Folate:                             1.0,
		ForcedExpiratoryVolume:             1.0,
		ForcedVitalCapacity:                1.0,
		Handwashing:                        1.0,
		HeadphoneAudioExposure:             1.0,
		HeartRateMin:                       1.0,
		HeartRateMax:                       1.0,
		HeartRateAvg:                       1.0,
		HeartRateVariability:               1.0,
		Height:                             1.0,
		HighHeartRateNotificationsMin:      1.0,
		HighHeartRateNotificationsMax:      1.0,
		HighHeartRateNotificationsAvg:      1.0,
		InhalerUsage:                       1.0,
		InsulinDelivery:                    1.0,
		Iodine:                             1.0,
		Iron:                               1.0,
		IrregularHeartRateNotificationsMin: 1.0,
		IrregularHeartRateNotificationsMax: 1.0,
		IrregularHeartRateNotificationsAvg: 1.0,
		LeanBodyMass:                       1.0,
		LowHeartRateNotificationsMin:       1.0,
		LowHeartRateNotificationsMax:       1.0,
		LowHeartRateNotificationsAvg:       1.0,
		Magnesium:                          1.0,
		Manganese:                          1.0,
		MindfulMinutes:                     1.0,
		Molybdenum:                         1.0,
		MonounsaturatedFat:                 1.0,
		Niacin:                             1.0,
		NumberOfTimesFallen:                1.0,
		PantothenicAcid:                    1.0,
		PeakExpiratoryFlowRate:             1.0,
		PeripheralPerfusionIndex:           1.0,
		PhysicalEffort:                     1.0,
		PolyunsaturatedFat:                 1.0,
		Potassium:                          1.0,
		Protein:                            1.0,
		PushCount:                          1.0,
		RespiratoryRate:                    1.0,
		RestingEnergy:                      1.0,
		RestingHeartRate:                   1.0,
		Riboflavin:                         1.0,
		RunningGroundContactTime:           1.0,
		RunningPower:                       1.0,
		RunningSpeed:                       1.0,
		RunningStride:                      1.0,
		RunningVerticalOscillation:         1.0,
		SaturatedFat:                       1.0,
		Selenium:                           1.0,
		SexualActivityUnspecified:          1.0,
		SexualActivityProtectionUsed:       1.0,
		SexualActivityProtectionNotUsed:    1.0,
		SixMinuteWalkingTestDistance:       1.0,
		SleepAnalysisAsleep:                1.0,
		SleepAnalysisInBed:                 1.0,
		SleepAnalysisCore:                  1.0,
		SleepAnalysisDeep:                  1.0,
		SleepAnalysisREM:                   1.0,
		SleepAnalysisAwake:                 1.0,
		Sodium:                             1.0,
		StairSpeedDown:                     1.0,
		StairSpeedUp:                       1.0,
		StepCount:                          1.0,
		SwimmingDistance:                   1.0,
		SwimmingStrokeCount:                1.0,
		Thiamin:                            1.0,
		TimeInDaylight:                     1.0,
		Toothbrushing:                      1.0,
		TotalFat:                           1.0,
		UVExposure:                         1.0,
		UnderwaterDepth:                    1.0,
		UnderwaterTemperature:              1.0,
		VO2Max:                             1.0,
		VitaminA:                           1.0,
		VitaminB12:                         1.0,
		VitaminB6:                          1.0,
		VitaminC:                           1.0,
		VitaminD:                           1.0,
		VitaminE:                           1.0,
		VitaminK:                           1.0,
		WaistCircumference:                 1.0,
		WalkingPlusRunningDistance:         1.0,
		WalkingAsymmetryPercentage:         1.0,
		WalkingDoubleSupportPercentage:     1.0,
		WalkingHeartRateAverage:            1.0,
		WalkingSpeed:                       1.0,
		WalkingStepLength:                  1.0,
		Water:                              1.0,
		WeightDividedByBodyMass:            1.0,
		WheelchairDistance:                 1.0,
		Zinc:                               1.0,
	}
}

func getSummaryDataFromDatabaseByDates(t *testing.T, desiredDates []string) []controller.Summary {

	dbConfig := controller.Database{
		User:     "username",
		Password: "password",
		Host:     "warehouse",
		Port:     "5432",
		Database: "appleHealth",
	}

	summaries, err := controller.GetSummaryDataFromDatabase(dbConfig)
	testkit.AssertNoError(t, err)

	filteredSummaries := []controller.Summary{}
	for _, summary := range summaries {
		for _, date := range desiredDates {
			if summary.Date == date {
				filteredSummaries = append(filteredSummaries, summary)
				break
			}
		}
	}

	return filteredSummaries
}
