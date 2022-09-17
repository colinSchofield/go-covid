package controller

import (
	"fmt"
	"net/http"

	"git.com/colinSchofield/go-covid/config"
	"git.com/colinSchofield/go-covid/service"
	"github.com/gin-gonic/gin"
)

type CovidController interface {
	GetCovid19DailySummary(context *gin.Context)
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
