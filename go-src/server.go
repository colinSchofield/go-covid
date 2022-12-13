package main

/*
	The main package routes requests to the controller.

	This microservice provides support for a React based responsive web application. It is hosted on Kubernetes via AWS EKS.
	This microservice uses gin as its web container, which provides an elegant routing of services with high performance.
	The components are all instantiated in this file via constructor injection, allowing mocking and improved testability.

	The design pattern of 'separation of concerns' is employed here. The layers are split up using the Go package mechanism,
	as shown below:

	main -- routes the request to the controller
	controller -- this layer has direct access to the web/http layer. Its purpose is to mediate access to the service layer
	service -- the service layer provides a boundary to the backend, exposed through a set of interfaces
	client -- controlled by the service layer. It provides access to 3rd party API's, message and event brokers
	custom_error -- the set of application specific custom errors, that are typically used for communication between layers
	model -- Used as the VO (i.e Value Objects) for communicating between layers and ultimately as JSON output to the web
*/

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/colinSchofield/go-covid/config"
	"github.com/colinSchofield/go-covid/controller"
	"github.com/colinSchofield/go-covid/service"
	"github.com/colinSchofield/go-covid/service/client"
	"github.com/gin-gonic/gin"
)

// The components are all instantiated in this file via constructor injection, allowing mocking and improved testability
var (
	userService   service.UserService   = service.NewUserService()
	regionService service.RegionService = service.NewRegionService()
	summaryClient client.SummaryClient  = client.NewSummaryClient()
	historyClient client.HistoryClient  = client.NewHistoryClient()

	covidService     service.CovidService        = service.NewCovidService(summaryClient, historyClient)
	covidController  controller.CovidController  = controller.NewCovidController(covidService)
	regionController controller.RegionController = controller.NewRegionController(regionService)
	userController   controller.UserController   = controller.NewUserController(userService)
)

// Middleware allowing CORS (Cross-Origin Resource Sharing) Access
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

var ginLambda *ginadapter.GinLambda

func init() {
	config.Logger().Info("Starting Rest API Service..")
	apiVersion := config.GetApiVersion()
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET(apiVersion+"/list/daily", covidController.GetCovid19DailySummary)
	router.GET(apiVersion+"/list/monthly/:country", covidController.GetCovid19History)
	router.GET(apiVersion+"/list/regions", regionController.GetListOfRegions)

	router.POST(apiVersion+"/user", userController.CreateUser)
	router.PUT(apiVersion+"/user/:id", userController.UpdateUser)
	router.GET(apiVersion+"/user/:id", userController.GetUser)
	router.GET(apiVersion+"/user/list", userController.GetListOfAllUsers)
	router.DELETE(apiVersion+"/user/:id", userController.DeleteUser)

	ginLambda = ginadapter.New(router)
}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(HandleRequest)
}
