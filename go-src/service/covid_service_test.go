package service

import (
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"testing"

	"git.com/colinSchofield/go-covid/config"
	"git.com/colinSchofield/go-covid/custom_error"
	"git.com/colinSchofield/go-covid/model/daily"
	"git.com/colinSchofield/go-covid/model/history"
	"git.com/colinSchofield/go-covid/service/client"
	"github.com/stretchr/testify/assert"
)

type summaryClientMock struct {
	file string
}

func (sm summaryClientMock) GetCovid19DailySummary() (daily.Daily, error) {
	jsonAsString, err := os.ReadFile(sm.file)
	if err != nil {
		return daily.Daily{}, err
	}

	var daily daily.Daily
	if err := json.Unmarshal([]byte(jsonAsString), &daily); err != nil {
		return daily, err
	}

	return daily, nil
}

type summaryClientErrorMock struct{}

func (sm summaryClientErrorMock) GetCovid19DailySummary() (daily.Daily, error) {
	return daily.Daily{}, errors.New("Error reading Daily Summary")
}

type historyClientMock struct{}

func (hc historyClientMock) GetCovid19History(iso string) ([]history.History, error) {
	return []history.History{
		{Country: "", Date: "1-1-20", NewCases: 100, NewDeaths: 2},
		{Country: "", Date: "1-1-19", NewCases: 100, NewDeaths: 4},
	}, nil
}

type historyClientErrorMock struct{}

func (hc historyClientErrorMock) GetCovid19History(iso string) ([]history.History, error) {
	return nil, errors.New("A non-Timeout error occurred")
}

type historyClientTimeoutErrorMock struct{}

func (hc historyClientTimeoutErrorMock) GetCovid19History(iso string) ([]history.History, error) {
	return []history.History{
		{Country: "", Date: "1-1-20", NewCases: 100, NewDeaths: 2},
		{Country: "", Date: "1-1-19", NewCases: 100, NewDeaths: 4},
	}, custom_error.ClientTimeout{Wrapped: errors.New("a timeout error")}
}

func setHistoryEnvironmentVariables(t *testing.T) {
	t.Setenv(config.SUMMARY_END_POINT, "https://mock.com")
	t.Setenv(config.SUMMARY_HOST, "mock")
	t.Setenv(config.SUMMARY_KEY, "mock")
	t.Setenv(config.HISTORY_END_POINT, "https://mock.com")
	t.Setenv(config.HISTORY_HOST, "mock")
	t.Setenv(config.HISTORY_KEY, "mock")
	t.Setenv(config.EXCLUDE_REGIONS, "mock")
}

func Test_GetCovid19DailySummary(t *testing.T) {
	// Given
	setHistoryEnvironmentVariables(t)
	multiTest := []struct {
		fileName       string
		excludeRegions string
		elementsInMap  int
	}{
		{fileName: "test/model1.json", excludeRegions: "China", elementsInMap: 1},
		{fileName: "test/model1.json", excludeRegions: "China|Niue", elementsInMap: 0},
		{fileName: "test/model2.json", excludeRegions: "None", elementsInMap: 3},
		{fileName: "test/model2.json", excludeRegions: "China|Niue", elementsInMap: 1},
		{fileName: "test/model2.json", excludeRegions: "Italy|China|Niue|", elementsInMap: 0},
	}
	// When
	for _, test := range multiTest {
		clientMock := summaryClientMock{file: test.fileName}
		t.Setenv(config.EXCLUDE_REGIONS, test.excludeRegions)
		covidService := NewCovidService(clientMock, client.NewHistoryClient())
		if daily, err := covidService.GetCovid19DailySummary(); err != nil {
			t.Errorf("GetCovid19DailySummary returned an error of %s", err)
		} else {
			// Then
			assert.Equal(t, len(daily.Response), test.elementsInMap)
		}
	}
}

func Test_GetCovid19DailySummaryWithError(t *testing.T) {
	// Given
	setHistoryEnvironmentVariables(t)
	summaryClientError := summaryClientErrorMock{}
	covidService := NewCovidService(summaryClientError, client.NewHistoryClient())
	// When
	_, err := covidService.GetCovid19DailySummary()
	// Then
	assert.NotNil(t, err)
}

func Test_GetCovid19History(t *testing.T) {
	// Given
	multiTest := []struct {
		scenario          string
		country           string
		historyClientMock client.HistoryClient
		errorState        bool
	}{
		{scenario: "Happy Path", country: "Australia", historyClientMock: historyClientMock{}, errorState: false},
		{scenario: "Unknown Country", country: "Unknown", historyClientMock: historyClientMock{}, errorState: true},
		{scenario: "A non-timeout error", country: "China", historyClientMock: historyClientErrorMock{}, errorState: true},
		{scenario: "A timeout error", country: "Japan", historyClientMock: historyClientTimeoutErrorMock{}, errorState: false},
	}
	setHistoryEnvironmentVariables(t)
	// When
	for _, test := range multiTest {
		covidService := NewCovidService(client.NewSummaryClient(), test.historyClientMock)
		results, err := covidService.GetCovid19History(test.country)
		// Then
		if test.errorState {
			assert.Zero(t, len(results.Labels), test.scenario)
			assert.NotNil(t, err, test.scenario)
		} else {
			assert.NotZero(t, len(results.Labels), test.scenario)
			assert.Nil(t, err, test.scenario)
		}
	}
}

func Test_GenericReverseFunction(t *testing.T) {
	// Given
	input := []string{"1", "2", "3"}
	expected := []string{"3", "2", "1"}
	// When
	reverse(input)
	// Then
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("input (%v) and expected (%v) arrays are not equal", input, expected)
	}

	// Given
	inputInt := []int{1, 2, 3}
	expectedInt := []int{3, 2, 1}
	// When
	reverse(inputInt)
	// Then
	if !reflect.DeepEqual(inputInt, expectedInt) {
		t.Errorf("input (%v) and expected (%v) arrays are not equal", input, expected)
	}
}

func Test_GetDayInMonth(t *testing.T) {

	multiTest := []struct {
		input    string
		expected string
	}{
		{input: "2022-08-20", expected: "20"},
		{input: "2022-08-", expected: ""},
		{input: "2022/08/04", expected: "2022/08/04"},
		{input: "", expected: ""},
	}

	for _, test := range multiTest {
		assert.Equal(t, getDayInMonth(test.input), test.expected)
	}
}
