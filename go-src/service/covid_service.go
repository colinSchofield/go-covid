package service

import (
	"fmt"
	"os"
	"strings"

	"git.com/colinSchofield/go-covid/config"
	"git.com/colinSchofield/go-covid/model/daily"
	"git.com/colinSchofield/go-covid/model/history"
	"git.com/colinSchofield/go-covid/service/client"
)

type CovidService interface {
	GetCovid19DailySummary() (daily.Daily, error)
	GetCovid19History(country string) (history.TableDetails, error)
}

type covidService struct {
	summaryClient client.SummaryClient
	historyClient client.HistoryClient
	regionService RegionService
	excludeRegion string
}

func NewCovidService(summaryClient client.SummaryClient, historyClient client.HistoryClient) CovidService {
	excludeRegions := os.Getenv("EXCLUDE_REGIONS") // TODO
	if len(excludeRegions) == 0 {
		excludeRegions = "All|Asia|Oceania|Europe|North-America|Africa|South-America|Diamond-Princess-|Cura&ccedil;ao|R&eacute;union|MS-Zaandam-|Diamond-Princess|guam|Cook Islands|Palau|Nauru|Kiribati|Niue|Tuvalu|Tonga|Micronesia|DPRK"
	}

	return covidService{
		summaryClient: summaryClient,
		historyClient: historyClient,
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

func reverse[S any](input []S) {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}

func (cs covidService) GetCovid19History(country string) (history.TableDetails, error) {
	config.Logger().Debugf("Finding historical details for country %s", country)
	iso := cs.regionService.GetIsoForCountry(country)
	if iso == "" {
		return history.TableDetails{}, fmt.Errorf("no iso for country of %s", country)
	}
	config.Logger().Debugf("Country %s, equates to iso of %s", country, iso)
	if historyStats, err := cs.historyClient.GetCovid19History(iso); err != nil {
		config.Logger().Warnf("Unexpected error occurred fetching the historical information: %s", err)
		return history.TableDetails{}, err
	} else {

		labels := make([]string, len(historyStats))
		newCases := make([]int, len(historyStats))
		newDeaths := make([]int, len(historyStats))
		for ix, dayStats := range historyStats {
			labels[ix] = dayStats.Date
			newCases[ix] = dayStats.NewCases
			newDeaths[ix] = dayStats.NewDeaths
		}

		reverse(labels)
		reverse(newCases)
		reverse(newDeaths)

		return history.TableDetails{
			Flag:     cs.regionService.GetEmojiForCountry(country),
			Labels:   labels,
			NewCases: newCases,
			Deaths:   newDeaths,
		}, nil
	}
}
