package client

/*
	client -- controlled by the service layer. It provides access to 3rd party API's and message brokers

	RapidAPI is a market place for B2B API's and, in this particular application, it provides Covid statistics, as shown below:

	GetCovid19History -- Provides via a Rest API call to services on RapidAPI, the highlevel statistics for the past month
	NOTE: This service has a timeout value, that if exceeded, will cause some canned results to be returned.. This is because
	RapidAPI is NOT a professional service (for free information at least!) and consequently it is at times unreliable.
*/

import (
	"fmt"
	"strings"
	"time"

	"git.com/colinSchofield/go-covid/config"
	"git.com/colinSchofield/go-covid/custom_error"
	"git.com/colinSchofield/go-covid/model/history"
	"gopkg.in/resty.v1"
)

type HistoryClient interface {
	GetCovid19History(iso string) ([]history.History, error)
}

type historyClient struct {
	client      *resty.Client
	apiEndPoint string
	apiHost     string
	apiKey      string
}

func NewHistoryClient() HistoryClient {
	return historyClient{
		client:      resty.New(),
		apiEndPoint: config.GetHistoryEndPoint(),
		apiHost:     config.GetHistoryHost(),
		apiKey:      config.GetHistoryKey(),
	}
}

func (hc historyClient) GetCovid19History(iso string) ([]history.History, error) {

	var pastWeek []history.History
	endPoint := fmt.Sprintf(hc.apiEndPoint, iso)
	response, err := hc.client.
		SetTimeout(time.Duration(5*time.Second)).
		R().
		SetHeader("Accept", "application/json").
		SetHeader(hostHeader, hc.apiHost).
		SetHeader(apiKey, hc.apiKey).
		SetResult(&pastWeek).
		Get(endPoint)

	if err != nil && (strings.Contains(err.Error(), "Client.Timeout") || strings.Contains(err.Error(), "deadline exceeded")) {
		return fakeHistoricalData(), custom_error.ClientTimeout{Wrapped: fmt.Errorf("the client request resulted in a Timeout.. Error was: %w", err)}
	}

	if err != nil {
		wrappedError := fmt.Errorf("error acquiring Restful Web Service API.. The error was: %w", err)
		config.Logger().Error(wrappedError)
		return pastWeek, wrappedError
	}

	if response.StatusCode() != 200 {
		config.Logger().Error("HTTP Status Code indicated error: ", response.StatusCode())
		return pastWeek, fmt.Errorf("HTTP Status Code indicated error: %d", response.StatusCode())
	}

	config.Logger().Infof("We have received the a payload of size %d", len(response.Body()))
	config.Logger().Debugf("Contents of payload were: %s", response.Body())
	return pastWeek, nil
}

// RapidAPI is NOT a professional service and consequently it is at times unreliable -- here are a few 'canned' results
func fakeHistoricalData() []history.History {

	return []history.History{
		{Country: "", Date: "1-1-20", NewCases: 100, NewDeaths: 2},
		{Country: "", Date: "1-1-19", NewCases: 100, NewDeaths: 4},
		{Country: "", Date: "1-1-18", NewCases: 80, NewDeaths: 1},
		{Country: "", Date: "1-1-17", NewCases: 60, NewDeaths: 0},
		{Country: "", Date: "1-1-16", NewCases: 55, NewDeaths: 0},
		{Country: "", Date: "1-1-15", NewCases: 82, NewDeaths: 2},
		{Country: "", Date: "1-1-14", NewCases: 94, NewDeaths: 4},
		{Country: "", Date: "1-1-13", NewCases: 110, NewDeaths: 6},
		{Country: "", Date: "1-1-12", NewCases: 154, NewDeaths: 8},
		{Country: "", Date: "1-1-11", NewCases: 141, NewDeaths: 4},
		{Country: "", Date: "1-1-10", NewCases: 147, NewDeaths: 9},
		{Country: "", Date: "1-1-9", NewCases: 191, NewDeaths: 16},
		{Country: "", Date: "1-1-8", NewCases: 202, NewDeaths: 22},
		{Country: "", Date: "1-1-7", NewCases: 202, NewDeaths: 21},
		{Country: "", Date: "1-1-6", NewCases: 206, NewDeaths: 18},
		{Country: "", Date: "1-1-5", NewCases: 255, NewDeaths: 31},
		{Country: "", Date: "1-1-4", NewCases: 251, NewDeaths: 28},
		{Country: "", Date: "1-1-3", NewCases: 230, NewDeaths: 24},
		{Country: "", Date: "1-1-2", NewCases: 220, NewDeaths: 18},
		{Country: "", Date: "1-1-1", NewCases: 225, NewDeaths: 20},
	}
}
