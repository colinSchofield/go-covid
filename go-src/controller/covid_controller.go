package controller

import (
	"errors"
	"fmt"
	"net/http"

	"git.com/colinSchofield/go-covid/config"
	"git.com/colinSchofield/go-covid/custom_error"
	"git.com/colinSchofield/go-covid/service"
	"github.com/gin-gonic/gin"
)

type CovidController interface {
	GetCovid19DailySummary(context *gin.Context)
	GetCovid19History(context *gin.Context)
}

type covidController struct {
	covidService service.CovidService
}

func NewCovidController(covidService service.CovidService) CovidController {
	return covidController{
		covidService: covidService,
	}
}

func (cc covidController) GetCovid19DailySummary(context *gin.Context) {

	if response, err := cc.covidService.GetCovid19DailySummary(); err != nil {
		wrappedError := fmt.Errorf("error in response to daily summary request! Returned error was: %w", err)
		config.Logger().Error(wrappedError)
		context.JSON(http.StatusInternalServerError, gin.H{"error": wrappedError.Error()})
	} else {
		context.JSON(http.StatusOK, response)
	}
}

func (cc covidController) GetCovid19History(context *gin.Context) {

	country := context.Param("country")
	if response, err := cc.covidService.GetCovid19History(country); err != nil {
		if errors.As(err, &custom_error.ClientTimeout{}) {
			config.Logger().Warnf("Canned values were used, as the timeout was exceeded waiting for a response %v", err)
			context.JSON(http.StatusOK, response)
		} else if errors.As(err, &custom_error.NotFound{}) {
			wrappedError := fmt.Errorf("unknown country -- cannot find a matching iso. Error was: %w", err)
			config.Logger().Error(wrappedError)
			context.JSON(http.StatusBadRequest, gin.H{"error": wrappedError.Error()})
		} else {
			wrappedError := fmt.Errorf("error in response to historical statistics request! Returned error was: %w", err)
			config.Logger().Error(wrappedError)
			context.JSON(http.StatusBadGateway, gin.H{"error": wrappedError.Error()})
		}
	} else {
		context.JSON(http.StatusOK, response)
	}
}
