package client

import (
	"testing"
	"time"

	"github.com/colinSchofield/go-covid/config"
	"github.com/colinSchofield/go-covid/model/history"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func setHistoryEnvironmentVariables(t *testing.T) {
	t.Setenv(config.HISTORY_END_POINT, "https://mock.com/%s")
	t.Setenv(config.HISTORY_HOST, "mocked")
	t.Setenv(config.HISTORY_KEY, "mocked")
}

func Test_RestApiHistoryRequest(t *testing.T) {
	// Given
	httpmock.Activate() // Mocking via jarcoal/httpmock
	mockResponse := []history.History{{Country: "", Date: "1-1-20", NewCases: 100, NewDeaths: 2}}
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://mock.com/aus", httpmock.NewJsonResponderOrPanic(200, mockResponse))
	setHistoryEnvironmentVariables(t)
	client := NewHistoryClient()
	// When
	response, err := client.GetCovid19History("aus")
	// Then
	assert.Equal(t, err, nil)
	assert.Greater(t, len(response), 0)
}

func Test_RestApiHistoryRequestErrorCode(t *testing.T) {
	// Given
	httpmock.Activate() // Mocking via jarcoal/httpmock
	mockResponse := []history.History{{Country: "", Date: "1-1-20", NewCases: 100, NewDeaths: 2}}
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://mock.com/aus", httpmock.NewJsonResponderOrPanic(500, mockResponse))
	setHistoryEnvironmentVariables(t)
	client := NewHistoryClient()
	// When
	response, err := client.GetCovid19History("aus")
	// Then
	assert.NotNil(t, err)
	assert.Equal(t, len(response), 0)
}

func Test_RestApiHistoryRequestConnectionError(t *testing.T) {
	// Given
	httpmock.Activate() // Mocking via jarcoal/httpmock
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://mock.com/aus", nil)
	setHistoryEnvironmentVariables(t)
	client := NewHistoryClient()
	// When
	response, err := client.GetCovid19History("aus")
	// Then
	assert.NotNil(t, err)
	assert.Equal(t, len(response), 0)
}

func Test_RestApiHistoryRequestTimeout(t *testing.T) {
	// Given
	httpmock.Activate() // Mocking via jarcoal/httpmock
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://mock.com/aus", httpmock.NewStringResponder(200, "").Delay(6*time.Second))
	setHistoryEnvironmentVariables(t)
	client := NewHistoryClient()
	// When
	response, err := client.GetCovid19History("aus")
	// Then
	assert.NotNil(t, err)
	assert.Greater(t, len(response), 0)
}
