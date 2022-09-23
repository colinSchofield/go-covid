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
		apiEndPoint: config.GetSummaryEndPoint(),
		apiHost:     config.GetSummaryHost(),
		apiKey:      config.GetSummaryKey(),
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
		wrappedError := fmt.Errorf("error acquiring Restful Web Service API.. The error was: %w", err)
		config.Logger().Error(wrappedError)
		return summary, wrappedError
	}

	if response.StatusCode() != 200 {
		config.Logger().Error("HTTP Status Code indicated error: ", response.StatusCode())
		return summary, fmt.Errorf("HTTP Status Code indicated error: %d", response.StatusCode())
	}

	config.Logger().Infof("We have received the a payload of size %d", len(response.Body()))
	config.Logger().Debugf("Contents of payload were: %s", response.Body())
	return summary, nil
}
