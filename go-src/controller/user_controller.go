package controller

import (
	"fmt"
	"net/http"

	"git.com/colinSchofield/go-covid/config"
	"git.com/colinSchofield/go-covid/model/user"
	"git.com/colinSchofield/go-covid/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(context *gin.Context)
	UpdateUser(context *gin.Context)
	createOrUpdateUser(context *gin.Context)
	GetListOfAllUsers(context *gin.Context)
	GetUser(context *gin.Context)
	DeleteUser(context *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return userController{
		userService: userService,
	}
}

func (uc userController) CreateUser(context *gin.Context) {
	uc.createOrUpdateUser(context)
}

func (uc userController) UpdateUser(context *gin.Context) {
	uc.createOrUpdateUser(context)
}

func (uc userController) createOrUpdateUser(context *gin.Context) {
	if user, err := unmarshalUser(context); err == nil { // 'err == nil' is intentional as unmarshalUser handles error case
		if person, err := uc.userService.CreateOrUpdateUser(user); err != nil {
			errorString := fmt.Sprintf("Error creating (or updating) user %s! Returned error was: %s", user.Name, err)
			config.Logger().Errorf(errorString)
			context.JSON(
				http.StatusBadRequest,
				errorString,
			)
		} else {
			context.JSON(
				http.StatusCreated,
				person,
			)
		}
	}
}

func unmarshalUser(context *gin.Context) (user.User, error) {
	var user user.User
	if err := context.BindJSON(&user); err != nil {
		errorString := fmt.Sprintf("Error unmarshalling user (%s)! Returned error was: %s", user.Name, err)
		config.Logger().Errorf(errorString)
		context.JSON(
			http.StatusBadRequest,
			errorString,
		)
		return user, fmt.Errorf(errorString)
	} else {
		return user, nil
	}
}

func (uc userController) GetUser(context *gin.Context) {
	id := context.Param("id")
	if user, err := uc.userService.GetUser(id); err != nil {
		context.Writer.WriteHeader(http.StatusNotFound)
	} else {
		context.JSON(
			http.StatusOK,
			user,
		)
	}
}

func (uc userController) GetListOfAllUsers(context *gin.Context) {
	if users, err := uc.userService.GetListOfAllUsers(); err != nil {
		config.Logger().Warnf("error reading all users: %s", err)
		context.Writer.WriteHeader(http.StatusInternalServerError)
	} else {
		context.JSON(
			http.StatusOK,
			users,
		)
	}
}

func (uc userController) DeleteUser(context *gin.Context) {
	id := context.Param("id")
	if err := uc.userService.DeleteUser(id); err != nil {
		config.Logger().Warnf("error occurred deleting node: %s", err)
		context.Writer.WriteHeader(http.StatusNotFound)

	} else {
		context.Writer.WriteHeader(http.StatusOK)
	}
}
