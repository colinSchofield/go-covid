package main

import (
	"git.com/colinSchofield/go-covid/config"
	"git.com/colinSchofield/go-covid/controller"
	"git.com/colinSchofield/go-covid/service"
	"git.com/colinSchofield/go-covid/service/client"
	"github.com/gin-gonic/gin"
)

var (
	userService   service.UserService   = service.NewUserService()
	regionService service.RegionService = service.NewRegionService()
	summaryClient client.SummaryClient  = client.NewSummaryClient()

	covidService     service.CovidService        = service.NewCovidService(summaryClient)
	covidController  controller.CovidController  = controller.NewCovidController(covidService)
	regionController controller.RegionController = controller.NewRegionController(regionService)
	userController   controller.UserController   = controller.NewUserController(userService)
)

const (
	apiVersion = "/api/1.0"
)

func main() {
	config.Logger().Info("Starting Rest API Service..")
	router := gin.Default()
	router.GET(apiVersion+"/list/daily", covidController.GetCovid19DailySummary)
	router.GET(apiVersion+"/list/regions", regionController.GetListOfRegions)
	router.POST(apiVersion+"/user", userController.CreateUser)
	router.PUT(apiVersion+"/user", userController.UpdateUser)
	router.GET(apiVersion+"/user/:id", userController.GetUser)
	router.GET(apiVersion+"/user/list", userController.GetListOfAllUsers)
	router.DELETE(apiVersion+"/user/:id", userController.DeleteUser)
	router.Run()
}
