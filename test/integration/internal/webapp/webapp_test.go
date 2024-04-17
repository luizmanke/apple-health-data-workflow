package webapp_test

import (
	"apple-health-data-workflow/internal/controller"
	"apple-health-data-workflow/internal/webapp"
	"apple-health-data-workflow/pkg/testkit"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSummaryHandlerMustDisplayDropdownAndLineChart(t *testing.T) {

	responseRecorder := httptest.NewRecorder()
	dbConfig := controller.Database{
		User:     "username",
		Password: "password",
		Host:     "warehouse",
		Port:     "5432",
		Database: "apple_health",
	}

	webapp.DisplaySummaryChart(responseRecorder, dbConfig)

	response := responseRecorder.Result()
	bodyBytes, err := io.ReadAll(response.Body)
	htmlContent := string(bodyBytes)

	testkit.AssertNoError(t, err)
	testkit.AssertEqual(t, response.StatusCode, http.StatusOK)
	testkit.AssertContians(t, htmlContent, "<select id=\"choiceDropdown\" class=\"form-select\">")
	testkit.AssertContians(t, htmlContent, "<canvas id=\"lineChart\"></canvas>")
}
