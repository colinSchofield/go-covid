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
	userService      service.UserService         = service.NewUserService()
	userController   controller.UserController   = controller.NewUserController(userService)
)

const (
	apiVersion = "/api/1.0"
)

func main() {
	config.Logger().Info("Starting Rest API Service..")
	router := gin.Default()
	router.GET(apiVersion+"/list/regions", regionController.GetListOfRegions)
	router.POST(apiVersion+"/user", userController.CreateUser)
	router.PUT(apiVersion+"/user", userController.UpdateUser)
	router.GET(apiVersion+"/user/:id", userController.GetUser)
	router.GET(apiVersion+"/user/list", userController.GetListOfAllUsers)
	router.DELETE(apiVersion+"/user/:id", userController.DeleteUser)
	router.Run()
}
