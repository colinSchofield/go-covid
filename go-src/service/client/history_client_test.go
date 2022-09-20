package client

import (
	"testing"
)

func Test_RestApiHistoryRequest(t *testing.T) {
	// Given
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
