package client

import (
	"errors"
	"testing"

	"git.com/colinSchofield/go-covid/config"
	"git.com/colinSchofield/go-covid/custom_error"
)

func setHistoryEnvironmentVariables(t *testing.T) {
	t.Setenv(config.HISTORY_END_POINT, "https://vaccovid-coronavirus-vaccine-and-treatment-tracker.p.rapidapi.com/api/covid-ovid-data/sixmonth/%s")
	t.Setenv(config.HISTORY_HOST, "vaccovid-coronavirus-vaccine-and-treatment-tracker.p.rapidapi.com")
	t.Setenv(config.HISTORY_KEY, "cb1f09fd7dmsh35f7dd8afd27dfdp191e0cjsnca765ccf022a")
}

func Test_RestApiHistoryRequest(t *testing.T) {
	// Given
	setHistoryEnvironmentVariables(t)
	client := NewHistoryClient()
	// When
	response, err := client.GetCovid19History("aus")
	// Then
	if err != nil && !errors.As(err, &custom_error.ClientTimeout{}) {
		t.Error("ClientTimeout was not found in the error chain")
	}
	if len(response) == 0 {
		t.Error("Rest request -- the returned value is empty!?!")
	}
}
