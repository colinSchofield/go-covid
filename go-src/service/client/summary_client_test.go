package client

import (
	"testing"

	"github.com/colinSchofield/go-covid/config"
	"github.com/colinSchofield/go-covid/model/daily"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func setSummaryEnvironmentVariables(t *testing.T) {
	t.Setenv(config.SUMMARY_END_POINT, "https://mock.com")
	t.Setenv(config.SUMMARY_HOST, "mocked")
	t.Setenv(config.SUMMARY_KEY, "mocked")
}

func Test_RestApiClientRequest(t *testing.T) {
	// Given
	httpmock.Activate() // Mocking via jarcoal/httpmock
	mockResponse := daily.Daily{Results: 0}
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://mock.com", httpmock.NewJsonResponderOrPanic(200, mockResponse))
	setSummaryEnvironmentVariables(t)
	client := NewSummaryClient()
	// When
	response, err := client.GetCovid19DailySummary()
	// Then
	assert.Equal(t, err, nil)
	assert.NotNil(t, response)
}

func Test_RestApiClientRequestErrorCode(t *testing.T) {
	// Given
	httpmock.Activate() // Mocking via jarcoal/httpmock
	mockResponse := daily.Daily{Results: 0}
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://mock.com", httpmock.NewJsonResponderOrPanic(500, mockResponse))
	setSummaryEnvironmentVariables(t)
	client := NewSummaryClient()
	// When
	response, err := client.GetCovid19DailySummary()
	// Then
	assert.NotNil(t, err)
	assert.NotNil(t, response, 0)
}

func Test_RestApiClientRequestConnectionError(t *testing.T) {
	// Given
	httpmock.Activate() // Mocking via jarcoal/httpmock
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://mock.com", nil)
	setSummaryEnvironmentVariables(t)
	client := NewSummaryClient()
	// When
	response, err := client.GetCovid19DailySummary()
	// Then
	assert.NotNil(t, err)
	assert.NotNil(t, response, 0)
}
