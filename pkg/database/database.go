package database

import (
	"apple-health-data-workflow/pkg/models"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func InsertAppleHealthSummaryDataIntoDatabase(databaseConfig DatabaseConfig, summaries []models.Summary) error {

	dbConn, err := connectToTheDatabase(databaseConfig)
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

func GetSummaryDataFromDatabase(databaseConfig DatabaseConfig) ([]models.Summary, error) {

	dbConn, err := connectToTheDatabase(databaseConfig)
	if err != nil {
		return nil, err
	}
	defer dbConn.Close()

	rows, err := dbConn.Query(`SELECT * FROM summary`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	summaries := []models.Summary{}
	for rows.Next() {

		summary := models.Summary{}
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

func connectToTheDatabase(databaseConfig DatabaseConfig) (*sql.DB, error) {

	dbInfo := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		databaseConfig.User,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Database,
	)

	dbConn, err := sql.Open("postgres", dbInfo)
	return dbConn, err
}
