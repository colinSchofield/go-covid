package client

import (
	"fmt"

	"git.com/colinSchofield/go-covid/config"
	"git.com/colinSchofield/go-covid/model/daily"
	"gopkg.in/resty.v1"
)

type SummaryClient interface {
	GetCovid19DailySummary() (daily.Daily, error)
}

type summaryClient struct {
	client      *resty.Client
	apiEndPoint string
	apiHost     string
	apiKey      string
}

const (
	hostHeader = "x-rapidapi-host"
	apiKey     = "x-rapidapi-key"
)

func NewSummaryClient() SummaryClient {
	return summaryClient{
		client:      resty.New(),
		apiEndPoint: "https://covid-193.p.rapidapi.com/statistics",
		apiHost:     "covid-193.p.rapidapi.com",
		apiKey:      "cb1f09fd7dmsh35f7dd8afd27dfdp191e0cjsnca765ccf022a",
	}
}

func (sc summaryClient) GetCovid19DailySummary() (daily.Daily, error) {

	var summary daily.Daily
	response, err := sc.client.R().
		SetHeader("Accept", "application/json").
		SetHeader(hostHeader, sc.apiHost).
		SetHeader(apiKey, sc.apiKey).
		SetResult(&summary).
		Get(sc.apiEndPoint)

	if err != nil {
		config.Logger().Errorf("Error acquiring Restful Web Service API.. The error was: %s", err)
		return summary, err
	}

	if response.StatusCode() != 200 {
		config.Logger().Error("HTTP Status Code indicated error: ", response.StatusCode())
		return summary, fmt.Errorf("HTTP Status Code indicated error: %d", response.StatusCode())
	}

	config.Logger().Infof("We have received the a payload of size %d", len(response.Body()))
	config.Logger().Debugf("Contents of payload were: %s", response.Body())
	return summary, nil
}
