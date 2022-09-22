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
		errorString := fmt.Sprintf("Error in response to daily summary request! Returned error was: %s", err)
		config.Logger().Errorf(errorString)
		context.JSON(
			http.StatusBadGateway,
			errorString,
		)
	} else {
		context.JSON(
			http.StatusOK,
			response,
		)
	}
}

func (cc covidController) GetCovid19History(context *gin.Context) {
	country := context.Param("country")
	if response, err := cc.covidService.GetCovid19History(country); err != nil {
		if errors.As(err, &custom_error.ClientTimeout{}) {
			config.Logger().Warnf("Canned values were used, as the timeout was exceeded waiting for a response %s", err)
			context.JSON(
				http.StatusOK,
				response,
			)

		}
		errorString := fmt.Sprintf("Error in response to historical statistics request! Returned error was: %s", err)
		config.Logger().Errorf(errorString)
		context.JSON(
			http.StatusBadGateway,
			errorString,
		)
	} else {
		context.JSON(
			http.StatusOK,
			response,
		)
	}
}
