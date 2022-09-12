package client

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"gopkg.in/resty.v1"
)

func Test_RestApiRequest(t *testing.T) {
	// Given
	client := NewSummaryClient()
	// When
	response, err := client.GetCovid19DailySummary()
	// Then
	if err != nil {
		t.Errorf("Error encountered: %s", err)
	}
	if len(response) == 0 {
		t.Error("Rest request -- the returned value is empty!?!")
	}
}

func Test_RestApiViaMock(t *testing.T) {
	// Given
	client := NewSummaryClient()
	httpmock.ActivateNonDefault(resty.DefaultClient.GetClient())
	httpmock.RegisterResponder("GET", "https://covid-193.p.rapidapi.com/statistics",
		httpmock.NewStringResponder(200, "cool dude!?!"))
	// When
	response, err := client.GetCovid19DailySummary()
	// Then
	if err != nil {
		t.Errorf("Error encountered: %s", err)
	}
	if len(response) == 0 {
		t.Error("Rest request -- the returned value is empty!?!")
	}
}
