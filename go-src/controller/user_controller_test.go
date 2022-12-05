package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/colinSchofield/go-covid/model/user"
	"github.com/colinSchofield/go-covid/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type userServiceMock struct{}

func (us userServiceMock) CreateOrUpdateUser(person user.User) (user.User, error) {
	return user.User{}, nil
}
func (us userServiceMock) UpdateUser(person user.User) (user.User, error) {
	return user.User{}, nil
}
func (us userServiceMock) GetListOfAllUsers() ([]user.DecoratedUser, error) {
	return []user.DecoratedUser{}, nil
}
func (us userServiceMock) GetUser(id string) (user.User, error) {
	return user.User{}, nil
}
func (us userServiceMock) DeleteUser(id string) error {
	return nil
}
func (us userServiceMock) GetDecoratedUser(person user.User) (user.DecoratedUser, error) {
	return user.DecoratedUser{}, nil
}

type userServiceErrorMock struct{}

func (us userServiceErrorMock) CreateOrUpdateUser(person user.User) (user.User, error) {
	return user.User{}, errors.New("AWS Error")
}
func (us userServiceErrorMock) UpdateUser(person user.User) (user.User, error) {
	return user.User{}, errors.New("AWS Error")
}
func (us userServiceErrorMock) GetListOfAllUsers() ([]user.DecoratedUser, error) {
	return []user.DecoratedUser{}, errors.New("AWS Error")
}
func (us userServiceErrorMock) GetUser(id string) (user.User, error) {
	return user.User{}, errors.New("AWS Error")
}
func (us userServiceErrorMock) DeleteUser(id string) error {
	return errors.New("AWS Error")
}
func (us userServiceErrorMock) GetDecoratedUser(person user.User) (user.DecoratedUser, error) {
	return user.DecoratedUser{}, errors.New("AWS Error")
}

var (
	region     = []string{"Australia", "S-Korea"}
	createBody = map[string]interface{}{
		"id":      "",
		"name":    "Owen",
		"age":     41,
		"gender":  "Male",
		"regions": region,
		"email":   "",
		"sms":     "",
	}
	updateBody = map[string]interface{}{
		"id":      "79f043fe-feec-4582-b856-cdb6f04eec29",
		"name":    "Owen",
		"age":     41,
		"gender":  "Male",
		"regions": region,
		"email":   "",
		"sms":     "",
	}
)

func Test_AllUserControllerViaMultiTest(t *testing.T) {

	multiTest := []struct {
		scenario        string
		requestType     string
		requestBody     any
		userServiceMock service.UserService
		methodCall      string
		expectedStatus  int
	}{
		{scenario: "Create a User, with the Happy Path", requestType: "POST", requestBody: createBody, userServiceMock: userServiceMock{}, methodCall: "CreateUser", expectedStatus: http.StatusOK},
		{scenario: "Create a User, but the user already exists", requestType: "POST", requestBody: updateBody, userServiceMock: userServiceMock{}, methodCall: "CreateUser", expectedStatus: http.StatusBadRequest},
		{scenario: "Create a User, but there was an AWS error", requestType: "POST", requestBody: createBody, userServiceMock: userServiceErrorMock{}, methodCall: "CreateUser", expectedStatus: http.StatusInternalServerError},
		{scenario: "Update a User, with the Happy Path", requestType: "PUT", requestBody: updateBody, userServiceMock: userServiceMock{}, methodCall: "UpdateUser", expectedStatus: http.StatusOK},
		{scenario: "Update a User, but there was an AWS error", requestType: "PUT", requestBody: updateBody, userServiceMock: userServiceErrorMock{}, methodCall: "UpdateUser", expectedStatus: http.StatusInternalServerError},
		{scenario: "Get a User, with the Happy Path", requestType: "GET", requestBody: nil, userServiceMock: userServiceMock{}, methodCall: "GetUser", expectedStatus: http.StatusOK},
		{scenario: "Get a User, but no user was found", requestType: "GET", requestBody: nil, userServiceMock: userServiceErrorMock{}, methodCall: "GetUser", expectedStatus: http.StatusNotFound},
		{scenario: "Get a list of all User, with the Happy Path", requestType: "GET", requestBody: nil, userServiceMock: userServiceMock{}, methodCall: "GetListOfAllUsers", expectedStatus: http.StatusOK},
		{scenario: "Get a list of all Users, but there was an AWS error", requestType: "GET", requestBody: nil, userServiceMock: userServiceErrorMock{}, methodCall: "GetListOfAllUsers", expectedStatus: http.StatusInternalServerError},
		{scenario: "Delete a User, with the Happy Path", requestType: "DELETE", requestBody: nil, userServiceMock: userServiceMock{}, methodCall: "DeleteUser", expectedStatus: http.StatusOK},
		{scenario: "Delete a User, but there was an AWS error", requestType: "DELETE", requestBody: nil, userServiceMock: userServiceErrorMock{}, methodCall: "DeleteUser", expectedStatus: http.StatusInternalServerError},
	}

	for _, test := range multiTest {
		// Given
		context, record := mockHttpRequest(test.requestBody, test.requestType)
		userController := NewUserController(test.userServiceMock)
		// When
		method := reflect.ValueOf(userController).MethodByName(test.methodCall)
		param := make([]reflect.Value, 1)
		param[0] = reflect.ValueOf(context)
		method.Call(param)
		// Then
		assert.Equal(t, record.Code, test.expectedStatus, test.scenario)
	}
}

// Provides the tests with an HTTP Body
func mockHttpRequest(body any, httpMethod string) (*gin.Context, *httptest.ResponseRecorder) {

	gin.SetMode(gin.TestMode)
	record := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(record)

	context.Request = &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}

	context.Request.Method = httpMethod
	context.Request.Header.Set("Content-Type", "application/json")

	if jsonBytes, err := json.Marshal(body); err != nil {
		panic(err)
	} else {
		context.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
		return context, record
	}
}
