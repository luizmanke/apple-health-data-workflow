package ingester

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strings"
)

type Summary struct {
	Date                               string  `mapstructure:"Date"`
	ActiveEnergy                       float32 `mapstructure:"Active Energy (kJ)"`
	AlcoholConsumption                 float32 `mapstructure:"Alcohol Consumption (count)"`
	AppleExerciseTime                  float32 `mapstructure:"Apple Exercise Time (min)"`
	AppleMoveTime                      float32 `mapstructure:"Apple Move Time (min)"`
	AppleSleepingWristTemperature      float32 `mapstructure:"Apple Sleeping Wrist Temperature (ºC)"`
	AppleStandHour                     float32 `mapstructure:"Apple Stand Hour (hours)"`
	AppleStandTime                     float32 `mapstructure:"Apple Stand Time (min)"`
	AtrialFibrillationBurden           float32 `mapstructure:"Atrial Fibrillation Burden (%)"`
	BasalBodyTemperature               float32 `mapstructure:"Basal Body Temperature (ºC)"`
	BloodAlcoholContent                float32 `mapstructure:"Blood Alcohol Content (%)"`
	BloodGlucose                       float32 `mapstructure:"Blood Glucose (mmol/L)"`
	BloodOxygenSaturation              float32 `mapstructure:"Blood Oxygen Saturation (%)"`
	BloodPressureSystolic              float32 `mapstructure:"Blood Pressure [Systolic] (mmHg)"`
	BloodPressureDiastolic             float32 `mapstructure:"Blood Pressure [Diastolic] (mmHg)"`
	BodyFatPercentage                  float32 `mapstructure:"Body Fat Percentage (%)"`
	BodyMassIndex                      float32 `mapstructure:"Body Mass Index (count)"`
	BodyTemperature                    float32 `mapstructure:"Body Temperature (ºC)"`
	Caffeine                           float32 `mapstructure:"Caffeine (mg)"`
	Calcium                            float32 `mapstructure:"Calcium (mg)"`
	Carbohydrates                      float32 `mapstructure:"Carbohydrates (g)"`
	CardioRecovery                     float32 `mapstructure:"Cardio Recovery (count/min)"`
	Chloride                           float32 `mapstructure:"Chloride (mg)"`
	Chromium                           float32 `mapstructure:"Chromium (mcg)"`
	Copper                             float32 `mapstructure:"Copper (mg)"`
	CyclingCadence                     float32 `mapstructure:"Cycling Cadence (count/min)"`
	CyclingDistance                    float32 `mapstructure:"Cycling Distance (km)"`
	CyclingFunctionalThresholdPower    float32 `mapstructure:"Cycling Functional Threshold Power (watts)"`
	CyclingPower                       float32 `mapstructure:"Cycling Power (watts)"`
	CyclingSpeed                       float32 `mapstructure:"Cycling Speed (km/hr)"`
	DietaryBiotin                      float32 `mapstructure:"Dietary Biotin (mcg)"`
	DietaryCholesterol                 float32 `mapstructure:"Dietary Cholesterol (mg)"`
	DietaryEnergy                      float32 `mapstructure:"Dietary Energy (kJ)"`
	DietarySugar                       float32 `mapstructure:"Dietary Sugar (g)"`
	DistanceDownhillSnowSports         float32 `mapstructure:"Distance Downhill Snow Sports (km)"`
	ElectrodermalActivity              float32 `mapstructure:"Electrodermal Activity (S)"`
	EnvironmentalAudioExposure         float32 `mapstructure:"Environmental Audio Exposure (dBASPL)"`
	Fiber                              float32 `mapstructure:"Fiber (g)"`
	FlightsClimbed                     float32 `mapstructure:"Flights Climbed (count)"`
	Folate                             float32 `mapstructure:"Folate (mcg)"`
	ForcedExpiratoryVolume             float32 `mapstructure:"Forced Expiratory Volume 1 (L)"`
	ForcedVitalCapacity                float32 `mapstructure:"Forced Vital Capacity (L)"`
	Handwashing                        float32 `mapstructure:"Handwashing (s)"`
	HeadphoneAudioExposure             float32 `mapstructure:"Headphone Audio Exposure (dBASPL)"`
	HeartRateMin                       float32 `mapstructure:"Heart Rate [Min] (bpm)"`
	HeartRateMax                       float32 `mapstructure:"Heart Rate [Max] (bpm)"`
	HeartRateAvg                       float32 `mapstructure:"Heart Rate [Avg] (bpm)"`
	HeartRateVariability               float32 `mapstructure:"Heart Rate Variability (ms)"`
	Height                             float32 `mapstructure:"Height (m)"`
	HighHeartRateNotificationsMin      float32 `mapstructure:"High Heart Rate Notifications [Min] (count)"`
	HighHeartRateNotificationsMax      float32 `mapstructure:"High Heart Rate Notifications [Max] (count)"`
	HighHeartRateNotificationsAvg      float32 `mapstructure:"High Heart Rate Notifications [Avg] (count)"`
	InhalerUsage                       float32 `mapstructure:"Inhaler Usage (count)"`
	InsulinDelivery                    float32 `mapstructure:"Insulin Delivery (IU)"`
	Iodine                             float32 `mapstructure:"Iodine (mcg)"`
	Iron                               float32 `mapstructure:"Iron (mg)"`
	IrregularHeartRateNotificationsMin float32 `mapstructure:"Irregular Heart Rate Notifications [Min] (count)"`
	IrregularHeartRateNotificationsMax float32 `mapstructure:"Irregular Heart Rate Notifications [Max] (count)"`
	IrregularHeartRateNotificationsAvg float32 `mapstructure:"Irregular Heart Rate Notifications [Avg] (count)"`
	LeanBodyMass                       float32 `mapstructure:"Lean Body Mass (kg)"`
	LowHeartRateNotificationsMin       float32 `mapstructure:"Low Heart Rate Notifications [Min] (count)"`
	LowHeartRateNotificationsMax       float32 `mapstructure:"Low Heart Rate Notifications [Max] (count)"`
	LowHeartRateNotificationsAvg       float32 `mapstructure:"Low Heart Rate Notifications [Avg] (count)"`
	Magnesium                          float32 `mapstructure:"Magnesium (mg)"`
	Manganese                          float32 `mapstructure:"Manganese (mg)"`
	MindfulMinutes                     float32 `mapstructure:"Mindful Minutes (min)"`
	Molybdenum                         float32 `mapstructure:"Molybdenum (mcg)"`
	MonounsaturatedFat                 float32 `mapstructure:"Monounsaturated Fat (g)"`
	Niacin                             float32 `mapstructure:"Niacin (mg)"`
	NumberOfTimesFallen                float32 `mapstructure:"Number of Times Fallen (falls)"`
	PantothenicAcid                    float32 `mapstructure:"Pantothenic Acid (mg)"`
	PeakExpiratoryFlowRate             float32 `mapstructure:"Peak Expiratory Flow Rate (L/min)"`
	PeripheralPerfusionIndex           float32 `mapstructure:"Peripheral Perfusion Index (%)"`
	PhysicalEffort                     float32 `mapstructure:"Physical Effort (MET)"`
	PolyunsaturatedFat                 float32 `mapstructure:"Polyunsaturated Fat (g)"`
	Potassium                          float32 `mapstructure:"Potassium (mg)"`
	Protein                            float32 `mapstructure:"Protein (g)"`
	PushCount                          float32 `mapstructure:"Push Count (count)"`
	RespiratoryRate                    float32 `mapstructure:"Respiratory Rate (count/min)"`
	RestingEnergy                      float32 `mapstructure:"Resting Energy (kJ)"`
	RestingHeartRate                   float32 `mapstructure:"Resting Heart Rate (bpm)"`
	Riboflavin                         float32 `mapstructure:"Riboflavin (mg)"`
	RunningGroundContactTime           float32 `mapstructure:"Running Ground Contact Time (ms)"`
	RunningPower                       float32 `mapstructure:"Running Power (watts)"`
	RunningSpeed                       float32 `mapstructure:"Running Speed (km/hr)"`
	RunningStride                      float32 `mapstructure:"Running Stride Length (m)"`
	RunningVerticalOscillation         float32 `mapstructure:"Running Vertical Oscillation (cm)"`
	SaturatedFat                       float32 `mapstructure:"Saturated Fat (g)"`
	Selenium                           float32 `mapstructure:"Selenium (mcg)"`
	SexualActivityUnspecified          float32 `mapstructure:"Sexual Activity [Unspecified] (times)"`
	SexualActivityProtectionUsed       float32 `mapstructure:"Sexual Activity [Protection Used] (times)"`
	SexualActivityProtectionNotUsed    float32 `mapstructure:"Sexual Activity [Protection Not Used] (times)"`
	SixMinuteWalkingTestDistance       float32 `mapstructure:"Six-Minute Walking Test Distance (m)"`
	SleepAnalysisAsleep                float32 `mapstructure:"Sleep Analysis [Asleep] (hr)"`
	SleepAnalysisInBed                 float32 `mapstructure:"Sleep Analysis [In Bed] (hr)"`
	SleepAnalysisCore                  float32 `mapstructure:"Sleep Analysis [Core] (hr)"`
	SleepAnalysisDeep                  float32 `mapstructure:"Sleep Analysis [Deep] (hr)"`
	SleepAnalysisREM                   float32 `mapstructure:"Sleep Analysis [REM] (hr)"`
	SleepAnalysisAwake                 float32 `mapstructure:"Sleep Analysis [Awake] (hr)"`
	Sodium                             float32 `mapstructure:"Sodium (mg)"`
	StairSpeedDown                     float32 `mapstructure:"Stair Speed: Down (m/s)"`
	StairSpeedUp                       float32 `mapstructure:"Stair Speed: Up (m/s)"`
	StepCount                          float32 `mapstructure:"Step Count (steps)"`
	SwimmingDistance                   float32 `mapstructure:"Swimming Distance (m)"`
	SwimmingStrokeCount                float32 `mapstructure:"Swimming Stroke Count (count)"`
	Thiamin                            float32 `mapstructure:"Thiamin (mg)"`
	TimeInDaylight                     float32 `mapstructure:"Time in Daylight (min)"`
	Toothbrushing                      float32 `mapstructure:"Toothbrushing (s)"`
	TotalFat                           float32 `mapstructure:"Total Fat (g)"`
	UVExposure                         float32 `mapstructure:"UV Exposure (count)"`
	UnderwaterDepth                    float32 `mapstructure:"Underwater Depth (m)"`
	UnderwaterTemperature              float32 `mapstructure:"Underwater Temperature (ºC)"`
	VO2Max                             float32 `mapstructure:"VO2 Max (ml/(kg·min))"`
	VitaminA                           float32 `mapstructure:"Vitamin A (mcg)"`
	VitaminB12                         float32 `mapstructure:"Vitamin B12 (mcg)"`
	VitaminB6                          float32 `mapstructure:"Vitamin B6 (mg)"`
	VitaminC                           float32 `mapstructure:"Vitamin C (mg)"`
	VitaminD                           float32 `mapstructure:"Vitamin D (mcg)"`
	VitaminE                           float32 `mapstructure:"Vitamin E (mg)"`
	VitaminK                           float32 `mapstructure:"Vitamin K (mcg)"`
	WaistCircumference                 float32 `mapstructure:"Waist Circumference (cm)"`
	WalkingPlusRunningDistance         float32 `mapstructure:"Walking + Running Distance (km)"`
	WalkingAsymmetryPercentage         float32 `mapstructure:"Walking Asymmetry Percentage (%)"`
	WalkingDoubleSupportPercentage     float32 `mapstructure:"Walking Double Support Percentage (%)"`
	WalkingHeartRateAverage            float32 `mapstructure:"Walking Heart Rate Average (bpm)"`
	WalkingSpeed                       float32 `mapstructure:"Walking Speed (km/hr)"`
	WalkingStepLength                  float32 `mapstructure:"Walking Step Length (cm)"`
	Water                              float32 `mapstructure:"Water (mL)"`
	WeightDividedByBodyMass            float32 `mapstructure:"Weight/Body Mass (kg)"`
	WheelchairDistance                 float32 `mapstructure:"Wheelchair Distance (km)"`
	Zinc                               float32 `mapstructure:"Zinc (mg)"`
}

func (s Summary) Columns() string {

	object := reflect.ValueOf(s)
	columns := []string{}

	for i := 0; i < object.NumField(); i++ {
		field := object.Type().Field(i)
		columns = append(columns, field.Name)
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
