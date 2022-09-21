package client

import (
	"fmt"

	"git.com/colinSchofield/go-covid/config"
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
	response, err := hc.client.R().
		SetHeader("Accept", "application/json").
		SetHeader(hostHeader, hc.apiHost).
		SetHeader(apiKey, hc.apiKey).
		SetResult(&pastWeek).
		Get(endPoint)

	if err != nil {
		config.Logger().Errorf("Error acquiring Restful Web Service API.. The error was: %s", err)
		return pastWeek, err
	}

	if response.StatusCode() != 200 {
		config.Logger().Error("HTTP Status Code indicated error: ", response.StatusCode())
		return pastWeek, fmt.Errorf("HTTP Status Code indicated error: %d", response.StatusCode())
	}

	config.Logger().Infof("We have received the a payload of size %d", len(response.Body()))
	config.Logger().Debugf("Contents of payload were: %s", response.Body())
	return pastWeek, nil
}
