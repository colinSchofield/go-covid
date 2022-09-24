package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.com/colinSchofield/go-covid/custom_error"
	"git.com/colinSchofield/go-covid/model/daily"
	"git.com/colinSchofield/go-covid/model/history"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type covidServiceMock struct{}

func (cs covidServiceMock) GetCovid19DailySummary() (daily.Daily, error) {
	return daily.Daily{}, nil
}
func (cs covidServiceMock) GetCovid19History(country string) (history.TableDetails, error) {
	return history.TableDetails{}, nil
}

type covidServiceErrorMock struct{}

func (cs covidServiceErrorMock) GetCovid19DailySummary() (daily.Daily, error) {
	return daily.Daily{}, errors.New("AWS Error")
}
func (cs covidServiceErrorMock) GetCovid19History(country string) (history.TableDetails, error) {
	return history.TableDetails{}, errors.New("AWS Error")
}

type covidServiceTimeoutMock struct{}

func (cs covidServiceTimeoutMock) GetCovid19DailySummary() (daily.Daily, error) {
	return daily.Daily{}, custom_error.ClientTimeout{Wrapped: errors.New("time out")}
}
func (cs covidServiceTimeoutMock) GetCovid19History(country string) (history.TableDetails, error) {
	return history.TableDetails{}, custom_error.ClientTimeout{Wrapped: errors.New("time out")}
}

type covidServiceNotFoundMock struct{}

func (cs covidServiceNotFoundMock) GetCovid19DailySummary() (daily.Daily, error) {
	return daily.Daily{}, custom_error.NotFound{Wrapped: errors.New("not found")}
}
func (cs covidServiceNotFoundMock) GetCovid19History(country string) (history.TableDetails, error) {
	return history.TableDetails{}, custom_error.NotFound{Wrapped: errors.New("not found")}
}

func Test_GetCovid19DailySummaryHappyPath(t *testing.T) {
	// Given
	covidController := NewCovidController(covidServiceMock{})
	record := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(record)
	// When
	covidController.GetCovid19DailySummary(context)
	// Then
	assert.Equal(t, record.Code, http.StatusOK)
}

func Test_GetCovid19DailySummaryAWSError(t *testing.T) {
	// Given
	covidController := NewCovidController(covidServiceErrorMock{})
	record := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(record)
	// When
	covidController.GetCovid19DailySummary(context)
	// Then
	assert.Equal(t, record.Code, http.StatusInternalServerError)
}

func Test_GetCovid19HistoryHappyPath(t *testing.T) {
	// Given
	covidController := NewCovidController(covidServiceMock{})
	record := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(record)
	// When
	covidController.GetCovid19History(context)
	// Then
	assert.Equal(t, record.Code, http.StatusOK)
}

func Test_GetCovid19HistoryAWSError(t *testing.T) {
	// Given
	covidController := NewCovidController(covidServiceErrorMock{})
	record := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(record)
	// When
	covidController.GetCovid19History(context)
	// Then
	assert.Equal(t, record.Code, http.StatusInternalServerError)
}

func Test_GetCovid19HistoryTimeout(t *testing.T) {
	// Given
	covidController := NewCovidController(covidServiceTimeoutMock{})
	record := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(record)
	// When
	covidController.GetCovid19History(context)
	// Then
	assert.Equal(t, record.Code, http.StatusOK)
}

func Test_GetCovid19HistoryNotFound(t *testing.T) {
	// Given
	covidController := NewCovidController(covidServiceNotFoundMock{})
	record := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(record)
	// When
	covidController.GetCovid19History(context)
	// Then
	assert.Equal(t, record.Code, http.StatusBadRequest)
}
