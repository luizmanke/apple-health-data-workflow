package ingester

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lib/pq"
	"github.com/mitchellh/mapstructure"
)

type SourceStorage struct {
	Directory string
}

type DestinationDatabase struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func IngestAppleHealthSummaryData(sourceStorage SourceStorage, destinationDatabase DestinationDatabase) {

	fileNames, err := listCSVFiles(sourceStorage)
	if err != nil {
		panic(err)
	}

	err = createSummayTable(destinationDatabase)
	if err != nil {
		panic(err)
	}

	for _, fileName := range fileNames {

		summaries, err := readAppleHealthSummaryDataFromCSVFile(sourceStorage, fileName)
		if err != nil {
			panic(err)
		}

		err = insertAppleHealthSummaryDataIntoDatabase(destinationDatabase, summaries)
		if err != nil {
			panic(err)
		}
	}
}

func listCSVFiles(sourceStorage SourceStorage) ([]string, error) {
	fileNames := []string{}

	directory := sourceStorage.Directory
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

func createSummayTable(destinationDatabase DestinationDatabase) error {

	db, err := connectToTheDatabase(destinationDatabase)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS summary (
			Date TIMESTAMP PRIMARY KEY,
			ActiveEnergy FLOAT,
			AlcoholConsumption FLOAT,
			AppleExerciseTime FLOAT,
			AppleMoveTime FLOAT,
			AppleSleepingWristTemperature FLOAT,
			AppleStandHour FLOAT,
			AppleStandTime FLOAT,
			AtrialFibrillationBurden FLOAT,
			BasalBodyTemperature FLOAT,
			BloodAlcoholContent FLOAT,
			BloodGlucose FLOAT,
			BloodOxygenSaturation FLOAT,
			BloodPressureSystolic FLOAT,
			BloodPressureDiastolic FLOAT,
			BodyFatPercentage FLOAT,
			BodyMassIndex FLOAT,
			BodyTemperature FLOAT,
			Caffeine FLOAT,
			Calcium FLOAT,
			Carbohydrates FLOAT,
			CardioRecovery FLOAT,
			Chloride FLOAT,
			Chromium FLOAT,
			Copper FLOAT,
			CyclingCadence FLOAT,
			CyclingDistance FLOAT,
			CyclingFunctionalThresholdPower FLOAT,
			CyclingPower FLOAT,
			CyclingSpeed FLOAT,
			DietaryBiotin FLOAT,
			DietaryCholesterol FLOAT,
			DietaryEnergy FLOAT,
			DietarySugar FLOAT,
			DistanceDownhillSnowSports FLOAT,
			ElectrodermalActivity FLOAT,
			EnvironmentalAudioExposure FLOAT,
			Fiber FLOAT,
			FlightsClimbed FLOAT,
			Folate FLOAT,
			ForcedExpiratoryVolume FLOAT,
			ForcedVitalCapacity FLOAT,
			Handwashing FLOAT,
			HeadphoneAudioExposure FLOAT,
			HeartRateMin FLOAT,
			HeartRateMax FLOAT,
			HeartRateAvg FLOAT,
			HeartRateVariability FLOAT,
			Height FLOAT,
			HighHeartRateNotificationsMin FLOAT,
			HighHeartRateNotificationsMax FLOAT,
			HighHeartRateNotificationsAvg FLOAT,
			InhalerUsage FLOAT,
			InsulinDelivery FLOAT,
			Iodine FLOAT,
			Iron FLOAT,
			IrregularHeartRateNotificationsMin FLOAT,
			IrregularHeartRateNotificationsMax FLOAT,
			IrregularHeartRateNotificationsAvg FLOAT,
			LeanBodyMass FLOAT,
			LowHeartRateNotificationsMin FLOAT,
			LowHeartRateNotificationsMax FLOAT,
			LowHeartRateNotificationsAvg FLOAT,
			Magnesium FLOAT,
			Manganese FLOAT,
			MindfulMinutes FLOAT,
			Molybdenum FLOAT,
			MonounsaturatedFat FLOAT,
			Niacin FLOAT,
			NumberOfTimesFallen FLOAT,
			PantothenicAcid FLOAT,
			PeakExpiratoryFlowRate FLOAT,
			PeripheralPerfusionIndex FLOAT,
			PhysicalEffort FLOAT,
			PolyunsaturatedFat FLOAT,
			Potassium FLOAT,
			Protein FLOAT,
			PushCount FLOAT,
			RespiratoryRate FLOAT,
			RestingEnergy FLOAT,
			RestingHeartRate FLOAT,
			Riboflavin FLOAT,
			RunningGroundContactTime FLOAT,
			RunningPower FLOAT,
			RunningSpeed FLOAT,
			RunningStride FLOAT,
			RunningVerticalOscillation FLOAT,
			SaturatedFat FLOAT,
			Selenium FLOAT,
			SexualActivityUnspecified FLOAT,
			SexualActivityProtectionUsed FLOAT,
			SexualActivityProtectionNotUsed FLOAT,
			SixMinuteWalkingTestDistance FLOAT,
			SleepAnalysisAsleep FLOAT,
			SleepAnalysisInBed FLOAT,
			SleepAnalysisCore FLOAT,
			SleepAnalysisDeep FLOAT,
			SleepAnalysisREM FLOAT,
			SleepAnalysisAwake FLOAT,
			Sodium FLOAT,
			StairSpeedDown FLOAT,
			StairSpeedUp FLOAT,
			StepCount FLOAT,
			SwimmingDistance FLOAT,
			SwimmingStrokeCount FLOAT,
			Thiamin FLOAT,
			TimeInDaylight FLOAT,
			Toothbrushing FLOAT,
			TotalFat FLOAT,
			UVExposure FLOAT,
			UnderwaterDepth FLOAT,
			UnderwaterTemperature FLOAT,
			VO2Max FLOAT,
			VitaminA FLOAT,
			VitaminB12 FLOAT,
			VitaminB6 FLOAT,
			VitaminC FLOAT,
			VitaminD FLOAT,
			VitaminE FLOAT,
			VitaminK FLOAT,
			WaistCircumference FLOAT,
			WalkingPlusRunningDistance FLOAT,
			WalkingAsymmetryPercentage FLOAT,
			WalkingDoubleSupportPercentage FLOAT,
			WalkingHeartRateAverage FLOAT,
			WalkingSpeed FLOAT,
			WalkingStepLength FLOAT,
			Water FLOAT,
			WeightDividedByBodyMass FLOAT,
			WheelchairDistance FLOAT,
			Zinc FLOAT
		)
	`)

	return err
}

func readAppleHealthSummaryDataFromCSVFile(sourceStorage SourceStorage, fileName string) ([]Summary, error) {

	filePath := filepath.Join(sourceStorage.Directory, fileName)
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

func insertAppleHealthSummaryDataIntoDatabase(destinationDatabase DestinationDatabase, summaries []Summary) error {

	db, err := connectToTheDatabase(destinationDatabase)
	if err != nil {
		return err
	}
	defer db.Close()

	query := fmt.Sprintf(`
        INSERT INTO summary %s
        SELECT * FROM UNNEST($1::summary[])
		ON CONFLICT (date) DO NOTHING
    `, summaries[0].Columns())
	_, err = db.Exec(query, pq.Array(summaries))
	return err
}

func connectToTheDatabase(destinationDatabase DestinationDatabase) (*sql.DB, error) {

	dbInfo := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		destinationDatabase.User,
		destinationDatabase.Password,
		destinationDatabase.Host,
		destinationDatabase.Port,
		destinationDatabase.Database,
	)

	db, err := sql.Open("postgres", dbInfo)
	return db, err
}
