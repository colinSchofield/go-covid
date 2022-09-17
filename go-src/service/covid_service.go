package service

import (
	"os"
	"strings"

	"git.com/colinSchofield/go-covid/config"
	"git.com/colinSchofield/go-covid/model/daily"
	"git.com/colinSchofield/go-covid/service/client"
)

type CovidService interface {
	GetCovid19DailySummary() (daily.Daily, error)
}

type covidService struct {
	summaryClient client.SummaryClient
	regionService RegionService
	excludeRegion string
}

func NewCovidService(summaryClient client.SummaryClient) CovidService {
	excludeRegions := os.Getenv("EXCLUDE_REGIONS")
	if len(excludeRegions) == 0 {
		excludeRegions = "All|Asia|Oceania|Europe|North-America|Africa|South-America|Diamond-Princess-|ccedil|eacute|MS-Zaandam|Diamond-Princess"
	}
	return covidService{
		summaryClient: summaryClient,
		regionService: NewRegionService(),
		excludeRegion: excludeRegions,
	}
}

func (cs covidService) GetCovid19DailySummary() (daily.Daily, error) {
	if summary, err := cs.summaryClient.GetCovid19DailySummary(); err != nil {
		config.Logger().Warnf("Unexpected error occurred fetching the daily summary information: %s", err)
		return summary, err
	} else {
		ix := 0
		for _, location := range summary.Response {
			if !strings.Contains(cs.excludeRegion, location.Country) {
				location.DecoratedCountry = cs.regionService.GetEmojiForCountry(location.Country) + " " + location.Country
				summary.Response[ix] = location
				ix++
			}
		}
		summary.Response = summary.Response[:ix]
		return summary, nil
	}
}
