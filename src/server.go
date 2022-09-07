package main

import (
	"git.com/gol/controller"
	"git.com/gol/service"
	"github.com/gin-gonic/gin"
)

var (
	regionService    service.RegionService       = service.New()
	regionController controller.RegionController = controller.New(regionService)
)

const (
	apiVersion = "/api/1.0"
)

func main() {
	router := gin.Default()

	router.GET(apiVersion+"/list/regions", regionController.GetListOfRegions)

	router.Run()
}
