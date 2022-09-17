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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	config.Logger().Info("Starting Rest API Service..")
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET(apiVersion+"/list/daily", covidController.GetCovid19DailySummary)
	router.GET(apiVersion+"/list/regions", regionController.GetListOfRegions)
	router.POST(apiVersion+"/user", userController.CreateUser)
	router.PUT(apiVersion+"/user/:id", userController.UpdateUser)
	router.GET(apiVersion+"/user/:id", userController.GetUser)
	router.GET(apiVersion+"/user/list", userController.GetListOfAllUsers)
	router.DELETE(apiVersion+"/user/:id", userController.DeleteUser)
	router.Run()
}
