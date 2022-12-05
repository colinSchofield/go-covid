package controller

/*
	controller -- this layer has direct access to the web/http layer. Its purpose is to mediate access with the service layer

	This service provides access to CRUD (i.e. Create, Read, Update & Delete) of values stored in a User AWS DynamoDB.
	All appropriate Web HTTP responses and modes of access are used
*/

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/colinSchofield/go-covid/config"
	"github.com/colinSchofield/go-covid/custom_error"
	"github.com/colinSchofield/go-covid/model/user"
	"github.com/colinSchofield/go-covid/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(context *gin.Context)
	UpdateUser(context *gin.Context)
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

	switch updateUser, err := unmarshalUser(context); {
	case err != nil:
		config.Logger().Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case len(updateUser.Id) > 0:
		wrappedError := fmt.Errorf("error validating user on Create -- user contains an Id? request User: %v", updateUser)
		config.Logger().Error(wrappedError)
		context.JSON(http.StatusBadRequest, gin.H{"error": wrappedError.Error()})
	default:

		if resultUser, err := uc.userService.CreateOrUpdateUser(updateUser); err != nil {
			config.Logger().Error(err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			config.Logger().Infof("User was created successfully (%v)", resultUser)
			context.JSON(http.StatusOK, resultUser)
		}
	}
}

func (uc userController) UpdateUser(context *gin.Context) {

	if updateUser, err := unmarshalUser(context); err != nil {
		config.Logger().Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else if resultUser, err := uc.userService.CreateOrUpdateUser(updateUser); err != nil {
		config.Logger().Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		config.Logger().Infof("User was updated successfully (%v)", resultUser)
		context.JSON(http.StatusOK, updateUser)
	}
}

func unmarshalUser(context *gin.Context) (user.User, error) {

	var user user.User
	if err := context.BindJSON(&user); err != nil && strings.Contains(err.Error(), "validation") {
		wrappedError := fmt.Errorf("client side validation error: %w", err)
		return user, custom_error.Validation{Wrapped: wrappedError}
	} else {
		return user, err
	}
}

func (uc userController) GetUser(context *gin.Context) {

	id := context.Param("id")
	if user, err := uc.userService.GetUser(id); err != nil {
		wrappedError := fmt.Errorf("user with id of (%s) was not found: %w", id, err)
		config.Logger().Error(wrappedError)
		context.JSON(http.StatusNotFound, gin.H{"error": wrappedError.Error()})
	} else {
		context.JSON(http.StatusOK, user)
	}
}

func (uc userController) GetListOfAllUsers(context *gin.Context) {

	if users, err := uc.userService.GetListOfAllUsers(); err != nil {
		wrappedError := fmt.Errorf("error reading all users: %w", err)
		config.Logger().Error(wrappedError)
		context.JSON(http.StatusInternalServerError, gin.H{"error": wrappedError.Error()})
	} else {
		context.JSON(http.StatusOK, users)
	}
}

func (uc userController) DeleteUser(context *gin.Context) {

	id := context.Param("id")
	if err := uc.userService.DeleteUser(id); err != nil {
		wrappedError := fmt.Errorf("error occurred deleting node: %w", err)
		config.Logger().Error(wrappedError)
		context.JSON(http.StatusInternalServerError, gin.H{"error": wrappedError.Error()})
	} else {
		context.Writer.WriteHeader(http.StatusOK)
	}
}
