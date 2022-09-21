package client

import (
	"testing"

	"git.com/colinSchofield/go-covid/config"
	"github.com/jarcoal/httpmock"
	"gopkg.in/resty.v1"
)

func setSummaryEnvironmentVariables(t *testing.T) {
	t.Setenv(config.SUMMARY_END_POINT, "https://covid-193.p.rapidapi.com/statistics")
	t.Setenv(config.SUMMARY_HOST, "covid-193.p.rapidapi.com")
	t.Setenv(config.SUMMARY_KEY, "cb1f09fd7dmsh35f7dd8afd27dfdp191e0cjsnca765ccf022a")
}

func Test_RestApiRequest(t *testing.T) {
	// Given
	setSummaryEnvironmentVariables(t)
	client := NewSummaryClient()
	// When
	response, err := client.GetCovid19DailySummary()
	// Then
	if err != nil {
		t.Errorf("Error encountered: %s", err)
	}
	if len(response.Response) == 0 || len(response.Response[0].Country) == 0 {
		t.Error("Rest request -- the returned value is empty!?!")
	}
}

func Test_RestApiViaMock(t *testing.T) {
	// Given
	setSummaryEnvironmentVariables(t)
	client := NewSummaryClient()
	httpmock.ActivateNonDefault(resty.DefaultClient.GetClient())
	httpmock.RegisterResponder("GET", "https://covid-193.p.rapidapi.com/statistics",
		httpmock.NewStringResponder(200, "cool dude!?!")) // TODO fix this..
	// When
	response, err := client.GetCovid19DailySummary()
	// Then
	if err != nil {
		t.Errorf("Error encountered: %s", err)
	}
	if len(response.Response) == 0 {
		t.Error("Rest request -- the returned value is empty!?!")
	}
}
