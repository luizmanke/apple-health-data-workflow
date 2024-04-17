package controller

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strings"
)

type Summary struct {
	Date                               string  `db:"date" mapstructure:"Date"`
	ActiveEnergy                       float32 `db:"active_energy" mapstructure:"Active Energy (kJ)"`
	AlcoholConsumption                 float32 `db:"alcohol_consumption" mapstructure:"Alcohol Consumption (count)"`
	AppleExerciseTime                  float32 `db:"apple_exercise_time" mapstructure:"Apple Exercise Time (min)"`
	AppleMoveTime                      float32 `db:"apple_move_time" mapstructure:"Apple Move Time (min)"`
	AppleSleepingWristTemperature      float32 `db:"apple_sleeping_wrist_temperature" mapstructure:"Apple Sleeping Wrist Temperature (ºC)"`
	AppleStandHour                     float32 `db:"apple_stand_hour" mapstructure:"Apple Stand Hour (hours)"`
	AppleStandTime                     float32 `db:"apple_stand_time" mapstructure:"Apple Stand Time (min)"`
	AtrialFibrillationBurden           float32 `db:"atrial_fibrillation_burden" mapstructure:"Atrial Fibrillation Burden (%)"`
	BasalBodyTemperature               float32 `db:"basal_body_temperature" mapstructure:"Basal Body Temperature (ºC)"`
	BloodAlcoholContent                float32 `db:"blood_alcohol_content" mapstructure:"Blood Alcohol Content (%)"`
	BloodGlucose                       float32 `db:"blood_glucose" mapstructure:"Blood Glucose (mmol/L)"`
	BloodOxygenSaturation              float32 `db:"blood_oxygen_saturation" mapstructure:"Blood Oxygen Saturation (%)"`
	BloodPressureSystolic              float32 `db:"blood_pressure_systolic" mapstructure:"Blood Pressure [Systolic] (mmHg)"`
	BloodPressureDiastolic             float32 `db:"blood_pressure_diastolic" mapstructure:"Blood Pressure [Diastolic] (mmHg)"`
	BodyFatPercentage                  float32 `db:"body_fat_percentage" mapstructure:"Body Fat Percentage (%)"`
	BodyMassIndex                      float32 `db:"body_mass_index" mapstructure:"Body Mass Index (count)"`
	BodyTemperature                    float32 `db:"body_temperature" mapstructure:"Body Temperature (ºC)"`
	Caffeine                           float32 `db:"caffeine" mapstructure:"Caffeine (mg)"`
	Calcium                            float32 `db:"calcium" mapstructure:"Calcium (mg)"`
	Carbohydrates                      float32 `db:"carbohydrates" mapstructure:"Carbohydrates (g)"`
	CardioRecovery                     float32 `db:"cardio_recovery" mapstructure:"Cardio Recovery (count/min)"`
	Chloride                           float32 `db:"chloride" mapstructure:"Chloride (mg)"`
	Chromium                           float32 `db:"chromium" mapstructure:"Chromium (mcg)"`
	Copper                             float32 `db:"copper" mapstructure:"Copper (mg)"`
	CyclingCadence                     float32 `db:"cycling_cadence" mapstructure:"Cycling Cadence (count/min)"`
	CyclingDistance                    float32 `db:"cycling_distance" mapstructure:"Cycling Distance (km)"`
	CyclingFunctionalThresholdPower    float32 `db:"cycling_functional_threshold_power" mapstructure:"Cycling Functional Threshold Power (watts)"`
	CyclingPower                       float32 `db:"cycling_power" mapstructure:"Cycling Power (watts)"`
	CyclingSpeed                       float32 `db:"cycling_speed" mapstructure:"Cycling Speed (km/hr)"`
	DietaryBiotin                      float32 `db:"dietary_biotin" mapstructure:"Dietary Biotin (mcg)"`
	DietaryCholesterol                 float32 `db:"dietary_cholesterol" mapstructure:"Dietary Cholesterol (mg)"`
	DietaryEnergy                      float32 `db:"dietary_energy" mapstructure:"Dietary Energy (kJ)"`
	DietarySugar                       float32 `db:"dietary_sugar" mapstructure:"Dietary Sugar (g)"`
	DistanceDownhillSnowSports         float32 `db:"distance_downhill_snow_sports" mapstructure:"Distance Downhill Snow Sports (km)"`
	ElectrodermalActivity              float32 `db:"electrodermal_activity" mapstructure:"Electrodermal Activity (S)"`
	EnvironmentalAudioExposure         float32 `db:"environmental_audio_exposure" mapstructure:"Environmental Audio Exposure (dBASPL)"`
	Fiber                              float32 `db:"fiber" mapstructure:"Fiber (g)"`
	FlightsClimbed                     float32 `db:"flights_climbed" mapstructure:"Flights Climbed (count)"`
	Folate                             float32 `db:"folate" mapstructure:"Folate (mcg)"`
	ForcedExpiratoryVolume             float32 `db:"forced_expiratory_volume" mapstructure:"Forced Expiratory Volume 1 (L)"`
	ForcedVitalCapacity                float32 `db:"forced_vital_capacity" mapstructure:"Forced Vital Capacity (L)"`
	Handwashing                        float32 `db:"handwashing" mapstructure:"Handwashing (s)"`
	HeadphoneAudioExposure             float32 `db:"headphone_audio_exposure" mapstructure:"Headphone Audio Exposure (dBASPL)"`
	HeartRateMin                       float32 `db:"heart_rate_min" mapstructure:"Heart Rate [Min] (bpm)"`
	HeartRateMax                       float32 `db:"heart_rate_max" mapstructure:"Heart Rate [Max] (bpm)"`
	HeartRateAvg                       float32 `db:"heart_rate_avg" mapstructure:"Heart Rate [Avg] (bpm)"`
	HeartRateVariability               float32 `db:"heart_rate_variability" mapstructure:"Heart Rate Variability (ms)"`
	Height                             float32 `db:"height" mapstructure:"Height (m)"`
	HighHeartRateNotificationsMin      float32 `db:"high_heart_rate_notifications_min" mapstructure:"High Heart Rate Notifications [Min] (count)"`
	HighHeartRateNotificationsMax      float32 `db:"high_heart_rate_notifications_max" mapstructure:"High Heart Rate Notifications [Max] (count)"`
	HighHeartRateNotificationsAvg      float32 `db:"high_heart_rate_notifications_avg" mapstructure:"High Heart Rate Notifications [Avg] (count)"`
	InhalerUsage                       float32 `db:"inhaler_usage" mapstructure:"Inhaler Usage (count)"`
	InsulinDelivery                    float32 `db:"insulin_delivery" mapstructure:"Insulin Delivery (IU)"`
	Iodine                             float32 `db:"iodine" mapstructure:"Iodine (mcg)"`
	Iron                               float32 `db:"iron" mapstructure:"Iron (mg)"`
	IrregularHeartRateNotificationsMin float32 `db:"irregular_heart_rate_notifications_min" mapstructure:"Irregular Heart Rate Notifications [Min] (count)"`
	IrregularHeartRateNotificationsMax float32 `db:"irregular_heart_rate_notifications_max" mapstructure:"Irregular Heart Rate Notifications [Max] (count)"`
	IrregularHeartRateNotificationsAvg float32 `db:"irregular_heart_rate_notifications_avg" mapstructure:"Irregular Heart Rate Notifications [Avg] (count)"`
	LeanBodyMass                       float32 `db:"lean_body_mass" mapstructure:"Lean Body Mass (kg)"`
	LowHeartRateNotificationsMin       float32 `db:"low_heart_rate_notifications_min" mapstructure:"Low Heart Rate Notifications [Min] (count)"`
	LowHeartRateNotificationsMax       float32 `db:"low_heart_rate_notifications_max" mapstructure:"Low Heart Rate Notifications [Max] (count)"`
	LowHeartRateNotificationsAvg       float32 `db:"low_heart_rate_notifications_avg" mapstructure:"Low Heart Rate Notifications [Avg] (count)"`
	Magnesium                          float32 `db:"magnesium" mapstructure:"Magnesium (mg)"`
	Manganese                          float32 `db:"manganese" mapstructure:"Manganese (mg)"`
	MindfulMinutes                     float32 `db:"mindful_minutes" mapstructure:"Mindful Minutes (min)"`
	Molybdenum                         float32 `db:"molybdenum" mapstructure:"Molybdenum (mcg)"`
	MonounsaturatedFat                 float32 `db:"monounsaturated_fat" mapstructure:"Monounsaturated Fat (g)"`
	Niacin                             float32 `db:"niacin" mapstructure:"Niacin (mg)"`
	NumberOfTimesFallen                float32 `db:"number_of_times_fallen" mapstructure:"Number of Times Fallen (falls)"`
	PantothenicAcid                    float32 `db:"pantothenic_acid" mapstructure:"Pantothenic Acid (mg)"`
	PeakExpiratoryFlowRate             float32 `db:"peak_expiratory_flow_rate" mapstructure:"Peak Expiratory Flow Rate (L/min)"`
	PeripheralPerfusionIndex           float32 `db:"peripheral_perfusion_index" mapstructure:"Peripheral Perfusion Index (%)"`
	PhysicalEffort                     float32 `db:"physical_effort" mapstructure:"Physical Effort (MET)"`
	PolyunsaturatedFat                 float32 `db:"polyunsaturated_fat" mapstructure:"Polyunsaturated Fat (g)"`
	Potassium                          float32 `db:"potassium" mapstructure:"Potassium (mg)"`
	Protein                            float32 `db:"protein" mapstructure:"Protein (g)"`
	PushCount                          float32 `db:"push_count" mapstructure:"Push Count (count)"`
	RespiratoryRate                    float32 `db:"respiratory_rate" mapstructure:"Respiratory Rate (count/min)"`
	RestingEnergy                      float32 `db:"resting_energy" mapstructure:"Resting Energy (kJ)"`
	RestingHeartRate                   float32 `db:"resting_heart_rate" mapstructure:"Resting Heart Rate (bpm)"`
	Riboflavin                         float32 `db:"riboflavin" mapstructure:"Riboflavin (mg)"`
	RunningGroundContactTime           float32 `db:"running_ground_contact_time" mapstructure:"Running Ground Contact Time (ms)"`
	RunningPower                       float32 `db:"running_power" mapstructure:"Running Power (watts)"`
	RunningSpeed                       float32 `db:"running_speed" mapstructure:"Running Speed (km/hr)"`
	RunningStride                      float32 `db:"running_stride" mapstructure:"Running Stride Length (m)"`
	RunningVerticalOscillation         float32 `db:"running_vertical_oscillation" mapstructure:"Running Vertical Oscillation (cm)"`
	SaturatedFat                       float32 `db:"saturated_fat" mapstructure:"Saturated Fat (g)"`
	Selenium                           float32 `db:"selenium" mapstructure:"Selenium (mcg)"`
	SexualActivityUnspecified          float32 `db:"sexual_activity_unspecified" mapstructure:"Sexual Activity [Unspecified] (times)"`
	SexualActivityProtectionUsed       float32 `db:"sexual_activity_protection_used" mapstructure:"Sexual Activity [Protection Used] (times)"`
	SexualActivityProtectionNotUsed    float32 `db:"sexual_activity_protection_not_used" mapstructure:"Sexual Activity [Protection Not Used] (times)"`
	SixMinuteWalkingTestDistance       float32 `db:"six_minute_walking_test_distance" mapstructure:"Six-Minute Walking Test Distance (m)"`
	SleepAnalysisAsleep                float32 `db:"sleep_analysis_asleep" mapstructure:"Sleep Analysis [Asleep] (hr)"`
	SleepAnalysisInBed                 float32 `db:"sleep_analysis_in_bed" mapstructure:"Sleep Analysis [In Bed] (hr)"`
	SleepAnalysisCore                  float32 `db:"sleep_analysis_core" mapstructure:"Sleep Analysis [Core] (hr)"`
	SleepAnalysisDeep                  float32 `db:"sleep_analysis_deep" mapstructure:"Sleep Analysis [Deep] (hr)"`
	SleepAnalysisREM                   float32 `db:"sleep_analysis_rem" mapstructure:"Sleep Analysis [REM] (hr)"`
	SleepAnalysisAwake                 float32 `db:"sleep_analysis_awake" mapstructure:"Sleep Analysis [Awake] (hr)"`
	Sodium                             float32 `db:"sodium" mapstructure:"Sodium (mg)"`
	StairSpeedDown                     float32 `db:"stair_speed_down" mapstructure:"Stair Speed: Down (m/s)"`
	StairSpeedUp                       float32 `db:"stair_speed_up" mapstructure:"Stair Speed: Up (m/s)"`
	StepCount                          float32 `db:"step_count" mapstructure:"Step Count (steps)"`
	SwimmingDistance                   float32 `db:"swimming_distance" mapstructure:"Swimming Distance (m)"`
	SwimmingStrokeCount                float32 `db:"swimming_stroke_count" mapstructure:"Swimming Stroke Count (count)"`
	Thiamin                            float32 `db:"thiamin" mapstructure:"Thiamin (mg)"`
	TimeInDaylight                     float32 `db:"time_in_daylight" mapstructure:"Time in Daylight (min)"`
	Toothbrushing                      float32 `db:"toothbrushing" mapstructure:"Toothbrushing (s)"`
	TotalFat                           float32 `db:"total_fat" mapstructure:"Total Fat (g)"`
	UVExposure                         float32 `db:"uv_exposure" mapstructure:"UV Exposure (count)"`
	UnderwaterDepth                    float32 `db:"underwater_depth" mapstructure:"Underwater Depth (m)"`
	UnderwaterTemperature              float32 `db:"underwater_temperature" mapstructure:"Underwater Temperature (ºC)"`
	VO2Max                             float32 `db:"vo2_ax" mapstructure:"VO2 Max (ml/(kg·min))"`
	VitaminA                           float32 `db:"vitamin_a" mapstructure:"Vitamin A (mcg)"`
	VitaminB12                         float32 `db:"vitamin_b12" mapstructure:"Vitamin B12 (mcg)"`
	VitaminB6                          float32 `db:"vitamin_b6" mapstructure:"Vitamin B6 (mg)"`
	VitaminC                           float32 `db:"vitamin_c" mapstructure:"Vitamin C (mg)"`
	VitaminD                           float32 `db:"vitamin_d" mapstructure:"Vitamin D (mcg)"`
	VitaminE                           float32 `db:"vitamin_e" mapstructure:"Vitamin E (mg)"`
	VitaminK                           float32 `db:"vitamin_k" mapstructure:"Vitamin K (mcg)"`
	WaistCircumference                 float32 `db:"waist_circumference" mapstructure:"Waist Circumference (cm)"`
	WalkingPlusRunningDistance         float32 `db:"walking_plus_running_distance" mapstructure:"Walking + Running Distance (km)"`
	WalkingAsymmetryPercentage         float32 `db:"walking_asymmetry_percentage" mapstructure:"Walking Asymmetry Percentage (%)"`
	WalkingDoubleSupportPercentage     float32 `db:"walking_double_support_percentage" mapstructure:"Walking Double Support Percentage (%)"`
	WalkingHeartRateAverage            float32 `db:"walking_heart_rate_average" mapstructure:"Walking Heart Rate Average (bpm)"`
	WalkingSpeed                       float32 `db:"walking_speed" mapstructure:"Walking Speed (km/hr)"`
	WalkingStepLength                  float32 `db:"walking_step_length" mapstructure:"Walking Step Length (cm)"`
	Water                              float32 `db:"water" mapstructure:"Water (mL)"`
	WeightDividedByBodyMass            float32 `db:"weight_divided_by_body_mass" mapstructure:"Weight/Body Mass (kg)"`
	WheelchairDistance                 float32 `db:"wheelchair_distance" mapstructure:"Wheelchair Distance (km)"`
	Zinc                               float32 `db:"zinc" mapstructure:"Zinc (mg)"`
}

func (s Summary) Columns() string {

	object := reflect.ValueOf(s)
	columns := []string{}

	for i := 0; i < object.NumField(); i++ {
		field := object.Type().Field(i)
		columns = append(columns, field.Tag.Get("db"))
	}

	return fmt.Sprintf("(%s)", strings.Join(columns, ","))
}

// Reference: https://stackoverflow.com/a/47622594
func (s Summary) Value() (driver.Value, error) {

	object := reflect.ValueOf(s)
	values := []string{}

	for i := 0; i < object.NumField(); i++ {
		field := object.Field(i)
		values = append(values, fmt.Sprintf("%v", field.Interface()))
	}

	return fmt.Sprintf("(%s)", strings.Join(values, ",")), nil
}
