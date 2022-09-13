package controller

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

func (controller regionController) GetListOfRegions(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		controller.regionService.GetListOfRegions(),
	)
}

func (controller regionController) GetEmojiForCountry(country string) string {
	return controller.regionService.GetEmojiForCountry(country)
}
