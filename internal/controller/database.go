package controller

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type Database struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func InsertAppleHealthSummaryDataIntoDatabase(database Database, summaries []Summary) error {

	err := createSummayTableIfNeeded(database)
	if err != nil {
		return err
	}

	dbConn, err := connectToTheDatabase(database)
	if err != nil {
		return err
	}
	defer dbConn.Close()

	query := fmt.Sprintf(`
        INSERT INTO summary %s
        SELECT * FROM UNNEST($1::summary[])
		ON CONFLICT (date) DO NOTHING
    `, summaries[0].Columns())
	_, err = dbConn.Exec(query, pq.Array(summaries))
	return err
}

func GetSummaryDataFromDatabase(database Database) ([]Summary, error) {

	dbConn, err := connectToTheDatabase(database)
	if err != nil {
		return nil, err
	}
	defer dbConn.Close()

	rows, err := dbConn.Query(`SELECT * FROM summary`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	summaries := []Summary{}
	for rows.Next() {

		summary := Summary{}
		err := rows.Scan(
			&summary.Date,
			&summary.ActiveEnergy,
			&summary.AlcoholConsumption,
			&summary.AppleExerciseTime,
			&summary.AppleMoveTime,
			&summary.AppleSleepingWristTemperature,
			&summary.AppleStandHour,
			&summary.AppleStandTime,
			&summary.AtrialFibrillationBurden,
			&summary.BasalBodyTemperature,
			&summary.BloodAlcoholContent,
			&summary.BloodGlucose,
			&summary.BloodOxygenSaturation,
			&summary.BloodPressureSystolic,
			&summary.BloodPressureDiastolic,
			&summary.BodyFatPercentage,
			&summary.BodyMassIndex,
			&summary.BodyTemperature,
			&summary.Caffeine,
			&summary.Calcium,
			&summary.Carbohydrates,
			&summary.CardioRecovery,
			&summary.Chloride,
			&summary.Chromium,
			&summary.Copper,
			&summary.CyclingCadence,
			&summary.CyclingDistance,
			&summary.CyclingFunctionalThresholdPower,
			&summary.CyclingPower,
			&summary.CyclingSpeed,
			&summary.DietaryBiotin,
			&summary.DietaryCholesterol,
			&summary.DietaryEnergy,
			&summary.DietarySugar,
			&summary.DistanceDownhillSnowSports,
			&summary.ElectrodermalActivity,
			&summary.EnvironmentalAudioExposure,
			&summary.Fiber,
			&summary.FlightsClimbed,
			&summary.Folate,
			&summary.ForcedExpiratoryVolume,
			&summary.ForcedVitalCapacity,
			&summary.Handwashing,
			&summary.HeadphoneAudioExposure,
			&summary.HeartRateMin,
			&summary.HeartRateMax,
			&summary.HeartRateAvg,
			&summary.HeartRateVariability,
			&summary.Height,
			&summary.HighHeartRateNotificationsMin,
			&summary.HighHeartRateNotificationsMax,
			&summary.HighHeartRateNotificationsAvg,
			&summary.InhalerUsage,
			&summary.InsulinDelivery,
			&summary.Iodine,
			&summary.Iron,
			&summary.IrregularHeartRateNotificationsMin,
			&summary.IrregularHeartRateNotificationsMax,
			&summary.IrregularHeartRateNotificationsAvg,
			&summary.LeanBodyMass,
			&summary.LowHeartRateNotificationsMin,
			&summary.LowHeartRateNotificationsMax,
			&summary.LowHeartRateNotificationsAvg,
			&summary.Magnesium,
			&summary.Manganese,
			&summary.MindfulMinutes,
			&summary.Molybdenum,
			&summary.MonounsaturatedFat,
			&summary.Niacin,
			&summary.NumberOfTimesFallen,
			&summary.PantothenicAcid,
			&summary.PeakExpiratoryFlowRate,
			&summary.PeripheralPerfusionIndex,
			&summary.PhysicalEffort,
			&summary.PolyunsaturatedFat,
			&summary.Potassium,
			&summary.Protein,
			&summary.PushCount,
			&summary.RespiratoryRate,
			&summary.RestingEnergy,
			&summary.RestingHeartRate,
			&summary.Riboflavin,
			&summary.RunningGroundContactTime,
			&summary.RunningPower,
			&summary.RunningSpeed,
			&summary.RunningStride,
			&summary.RunningVerticalOscillation,
			&summary.SaturatedFat,
			&summary.Selenium,
			&summary.SexualActivityUnspecified,
			&summary.SexualActivityProtectionUsed,
			&summary.SexualActivityProtectionNotUsed,
			&summary.SixMinuteWalkingTestDistance,
			&summary.SleepAnalysisAsleep,
			&summary.SleepAnalysisInBed,
			&summary.SleepAnalysisCore,
			&summary.SleepAnalysisDeep,
			&summary.SleepAnalysisREM,
			&summary.SleepAnalysisAwake,
			&summary.Sodium,
			&summary.StairSpeedDown,
			&summary.StairSpeedUp,
			&summary.StepCount,
			&summary.SwimmingDistance,
			&summary.SwimmingStrokeCount,
			&summary.Thiamin,
			&summary.TimeInDaylight,
			&summary.Toothbrushing,
			&summary.TotalFat,
			&summary.UVExposure,
			&summary.UnderwaterDepth,
			&summary.UnderwaterTemperature,
			&summary.VO2Max,
			&summary.VitaminA,
			&summary.VitaminB12,
			&summary.VitaminB6,
			&summary.VitaminC,
			&summary.VitaminD,
			&summary.VitaminE,
			&summary.VitaminK,
			&summary.WaistCircumference,
			&summary.WalkingPlusRunningDistance,
			&summary.WalkingAsymmetryPercentage,
			&summary.WalkingDoubleSupportPercentage,
			&summary.WalkingHeartRateAverage,
			&summary.WalkingSpeed,
			&summary.WalkingStepLength,
			&summary.Water,
			&summary.WeightDividedByBodyMass,
			&summary.WheelchairDistance,
			&summary.Zinc,
		)
		if err != nil {
			return nil, err
		}

		summaries = append(summaries, summary)
	}

	return summaries, nil
}

func connectToTheDatabase(database Database) (*sql.DB, error) {

	dbInfo := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		database.User,
		database.Password,
		database.Host,
		database.Port,
		database.Database,
	)

	dbConn, err := sql.Open("postgres", dbInfo)
	return dbConn, err
}

func createSummayTableIfNeeded(database Database) error {

	dbConn, err := connectToTheDatabase(database)
	if err != nil {
		return err
	}
	defer dbConn.Close()

	_, err = dbConn.Exec(`
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
