package controller

/*
	controller -- this layer has direct access to the web/http layer. Its purpose is to mediate access to the service layer

	This provides the regions of the world, as static values, that will not change as shown below:

	GetListOfRegions -- List of countries in the world, along with their emoji flag and iso values
	GetEmojiForCountry -- Provides, for a given country, its emoji flag
*/

import (
	"net/http"

	"git.com/colinSchofield/go-covid/service"
	"github.com/gin-gonic/gin"
)

type RegionController interface {
	GetListOfRegions(context *gin.Context)
	GetEmojiForCountry(country string) string
}

type regionController struct {
	regionService service.RegionService
}

func NewRegionController(regionService service.RegionService) RegionController {
	return regionController{
		regionService: regionService,
	}
}

func (rc regionController) GetListOfRegions(context *gin.Context) {
	context.JSON(http.StatusOK, rc.regionService.GetListOfRegions())
}

func (rc regionController) GetEmojiForCountry(country string) string {
	return rc.regionService.GetEmojiForCountry(country)
}
