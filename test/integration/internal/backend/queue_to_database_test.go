package backend_test

import (
	"apple-health-data-workflow/internal/backend"
	"apple-health-data-workflow/pkg/database"
	"apple-health-data-workflow/pkg/models"
	"apple-health-data-workflow/pkg/queue"
	"apple-health-data-workflow/pkg/testkit"
	"testing"

	_ "github.com/lib/pq"
)

type queueToDatabaseFixtures struct {
	queueConfig    queue.QueueConfig
	databaseConfig database.DatabaseConfig
}

func TestEveryAppleHealthSummaryColumnMustBeIngestedIntoTheDatabase(t *testing.T) {

	fixtures := queueToDatabaseSetUp(
		t,
		[]models.Summary{
			appleHealthSummaryDataWithAllColumns(),
		},
	)

	backend.ReadAppleHealthSummaryDataFromQueueAndIngestIntoDatabase(
		fixtures.queueConfig,
		fixtures.databaseConfig,
	)

	ingestedData := getSummaryDataFromDatabaseByDates(t, []string{"2024-01-01T00:00:00Z"})
	testkit.AssertEqual(t, ingestedData, []models.Summary{appleHealthSummaryDataWithAllColumns()})
}

func TestMultipleAppleHealthSummaryMessagesMustBeIngestedIntoTheDatabase(t *testing.T) {

	fixtures := queueToDatabaseSetUp(
		t,
		[]models.Summary{
			{Date: "2024-01-02T00:00:00Z"},
			{Date: "2024-01-03T00:00:00Z"},
			{Date: "2024-01-04T00:00:00Z"},
		},
	)

	backend.ReadAppleHealthSummaryDataFromQueueAndIngestIntoDatabase(
		fixtures.queueConfig,
		fixtures.databaseConfig,
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
		[]models.Summary{
			{Date: "2024-01-02T00:00:00Z"},
			{Date: "2024-01-03T00:00:00Z"},
			{Date: "2024-01-04T00:00:00Z"},
		},
	)
}

func TestOnlyTheFirstAppleHealthSummaryWithTheSameDateMustBePersistedInTheDatabase(t *testing.T) {

	fixtures := queueToDatabaseSetUp(
		t,
		[]models.Summary{
			{Date: "2024-01-05T00:00:00Z", ActiveEnergy: 1.0},
			{Date: "2024-01-05T00:00:00Z", ActiveEnergy: 2.0},
			{Date: "2024-01-05T00:00:00Z", ActiveEnergy: 3.0},
		},
	)

	backend.ReadAppleHealthSummaryDataFromQueueAndIngestIntoDatabase(
		fixtures.queueConfig,
		fixtures.databaseConfig,
	)

	ingestedData := getSummaryDataFromDatabaseByDates(t, []string{"2024-01-05T00:00:00Z"})
	testkit.AssertEqual(
		t,
		ingestedData,
		[]models.Summary{
			{Date: "2024-01-05T00:00:00Z", ActiveEnergy: 1.0},
		},
	)
}

func queueToDatabaseSetUp(t *testing.T, summaries []models.Summary) queueToDatabaseFixtures {

	queueConfig := queue.QueueConfig{
		Server: "datahub:9092",
		Topic:  "ingester-test.summary",
	}
	databaseConfig := database.DatabaseConfig{
		User:     "username",
		Password: "password",
		Host:     "warehouse",
		Port:     "5432",
		Database: "apple_health",
	}

	err := queue.SendAppleHealthSummaryDataToQueue(queueConfig, summaries)
	testkit.AssertNoError(t, err)

	return queueToDatabaseFixtures{
		queueConfig,
		databaseConfig,
	}
}

func getSummaryDataFromDatabaseByDates(t *testing.T, desiredDates []string) []models.Summary {

	databaseConfig := database.DatabaseConfig{
		User:     "username",
		Password: "password",
		Host:     "warehouse",
		Port:     "5432",
		Database: "apple_health",
	}

	summaries, err := database.GetSummaryDataFromDatabase(databaseConfig)
	testkit.AssertNoError(t, err)

	filteredSummaries := []models.Summary{}
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
