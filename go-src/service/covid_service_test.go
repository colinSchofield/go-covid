package service

import (
	"encoding/json"
	"os"
	"testing"

	"git.com/colinSchofield/go-covid/model/daily"
	"git.com/colinSchofield/go-covid/service/client"
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

func Test_GetCovid19DailySummary(t *testing.T) {

	multiTest := []struct {
		fileName       string
		excludeRegions string
		elementsInMap  int
	}{
		{
			fileName:       "test/model1.json",
			excludeRegions: "China",
			elementsInMap:  1,
		},
		{
			fileName:       "test/model1.json",
			excludeRegions: "China|Niue",
			elementsInMap:  0,
		},
		{
			fileName:       "test/model2.json",
			excludeRegions: "None",
			elementsInMap:  3,
		},
		{
			fileName:       "test/model2.json",
			excludeRegions: "China|Niue",
			elementsInMap:  1,
		},
		{
			fileName:       "test/model2.json",
			excludeRegions: "Italy|China|Niue|",
			elementsInMap:  0,
		},
	}

	for _, test := range multiTest {

		clientMock := summaryClientMock{file: test.fileName}
		t.Setenv("EXCLUDE_REGIONS", test.excludeRegions)
		covidService := NewCovidService(clientMock, client.NewHistoryClient())
		if daily, err := covidService.GetCovid19DailySummary(); err != nil {
			t.Errorf("GetCovid19DailySummary returned an error of %s", err)
		} else {
			if len(daily.Response) != test.elementsInMap {
				t.Errorf("GetCovid19DailySummary returned %d countries, but %d were expected.", len(daily.Response), test.elementsInMap)
			}
		}
	}
}
