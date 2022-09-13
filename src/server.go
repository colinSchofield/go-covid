package main

import (
	"git.com/colinSchofield/go-covid/config"
	"git.com/colinSchofield/go-covid/controller"
	"git.com/colinSchofield/go-covid/service"
	"github.com/gin-gonic/gin"
)

var (
	regionService    service.RegionService       = service.NewRegionService()
	regionController controller.RegionController = controller.NewRegionController(regionService)
)

const (
	apiVersion = "/api/1.0"
)

func main() {
	config.Logger().Info("Starting Rest API Service..")
	router := gin.Default()
	router.GET(apiVersion+"/list/regions", regionController.GetListOfRegions)
	router.Run()
}
