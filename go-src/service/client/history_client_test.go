package client

import (
	"testing"

	"git.com/colinSchofield/go-covid/config"
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
	if err != nil {
		t.Errorf("Error encountered: %s", err)
	}
	if len(response) == 0 {
		t.Error("Rest request -- the returned value is empty!?!")
	}
}
