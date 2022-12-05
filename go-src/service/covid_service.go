package service

/*
	service -- the service layer provides a boundary to the backend, exposed through a set of interfaces

	All requests relating to obtaining Covid information are directed through this interface, as shown below:

	GetCovid19DailySummary -- Provides via a Rest API call to services on RapidAPI, the daily Covid statistics
	GetCovid19History -- Provides via a Rest API call to services on RapidAPI, the highlevel statistics for the past month
*/

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/colinSchofield/go-covid/config"
	"github.com/colinSchofield/go-covid/custom_error"
	"github.com/colinSchofield/go-covid/model/daily"
	"github.com/colinSchofield/go-covid/model/history"
	"github.com/colinSchofield/go-covid/service/client"
	"github.com/patrickmn/go-cache"
)

type CovidService interface {
	GetCovid19DailySummary() (*daily.Daily, error)
	GetCovid19History(country string) (*history.TableDetails, error)
}

type covidService struct {
	clientCache   *cache.Cache
	summaryClient client.SummaryClient
	historyClient client.HistoryClient
	regionService RegionService
	excludeRegion string
}

func NewCovidService(summaryClient client.SummaryClient, historyClient client.HistoryClient) CovidService {
	ttl := time.Duration(config.GetCacheTimeToLive()) * time.Minute
	return covidService{
		clientCache:   cache.New(ttl, 4*ttl),
		summaryClient: summaryClient,
		historyClient: historyClient,
		regionService: NewRegionService(),
		excludeRegion: config.GetExcludeRegion(),
	}
}

func (cs covidService) GetCovid19DailySummary() (*daily.Daily, error) {

	if dailyCache, found := cs.clientCache.Get("daily"); found {
		config.Logger().Debug("Cache of Daily Summary was used")
		return dailyCache.(*daily.Daily), nil
	}

	if summary, err := cs.summaryClient.GetCovid19DailySummary(); err != nil {
		wrappedError := fmt.Errorf("unexpected error occurred fetching the daily summary information: %w", err)
		config.Logger().Error(wrappedError)
		return summary, wrappedError
	} else {
		ix := 0
		for _, location := range summary.Response {
			if !strings.Contains(cs.excludeRegion, location.Country) &&
				len(cs.regionService.GetIsoForCountry(location.Country)) > 0 {

				location.DecoratedCountry = cs.regionService.GetEmojiForCountry(location.Country) + " " + location.Country
				summary.Response[ix] = location
				ix++
			}
		}
		summary.Response = summary.Response[:ix]
		cs.clientCache.Add("daily", summary, cache.DefaultExpiration)
		config.Logger().Debug("Stored Cache of Daily Summary")
		return summary, nil
	}
}

func reverse[S any](input []S) {

	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}

func getDayInMonth(date string) string {

	split := strings.Split(date, "-")
	if len(split) != 3 {
		return date
	}
	return split[2]
}

func (cs covidService) GetCovid19History(country string) (*history.TableDetails, error) {

	if dailyCache, found := cs.clientCache.Get(country); found {
		config.Logger().Debug("Cache of Covid history for location %d was used", country)
		return dailyCache.(*history.TableDetails), nil
	}

	config.Logger().Debugf("Finding historical details for country %s", country)
	iso := cs.regionService.GetIsoForCountry(country)
	if iso == "" {
		return &history.TableDetails{}, custom_error.NotFound{Wrapped: fmt.Errorf("no iso for country of %s", country)}
	}
	config.Logger().Debugf("Country %s, equates to iso of %s", country, iso)
	if historyStats, err := cs.historyClient.GetCovid19History(iso); err != nil && !errors.As(err, &custom_error.ClientTimeout{}) {

		wrappedError := fmt.Errorf("unexpected error occurred fetching the historical information: %w", err)
		config.Logger().Error(wrappedError)
		return &history.TableDetails{}, wrappedError
	} else {

		config.Logger().Debugf("%d results were returned for country (%s)", len(historyStats), country)
		labels := make([]string, len(historyStats))
		newCases := make([]int, len(historyStats))
		newDeaths := make([]int, len(historyStats))
		for ix, dayStats := range historyStats {
			labels[ix] = getDayInMonth(dayStats.Date)
			newCases[ix] = dayStats.NewCases
			newDeaths[ix] = dayStats.NewDeaths
		}

		reverse(labels)
		reverse(newCases)
		reverse(newDeaths)

		historyStats := &history.TableDetails{
			Flag:     cs.regionService.GetEmojiForCountry(country),
			Labels:   labels,
			NewCases: newCases,
			Deaths:   newDeaths,
		}

		cs.clientCache.Add(country, historyStats, cache.DefaultExpiration)
		config.Logger().Debug("Stored Cache of Covid history for location %d", country)
		return historyStats, nil
	}
}
