package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/colinSchofield/go-covid/model"
	"github.com/colinSchofield/go-covid/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_HappyPathAllFlags(t *testing.T) {
	// Given
	service := service.NewRegionService()
	controller := NewRegionController(service)
	record := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(record)
	regions := []model.Region{}
	// When
	controller.GetListOfRegions(context)
	// Then
	assert.Equal(t, record.Code, http.StatusOK)
	if err := json.Unmarshal(record.Body.Bytes(), &regions); err != nil {
		t.Errorf("String could not be unmarshalled: %s", err)
	}
	assert.Greater(t, len(regions), 0)
	for _, region := range regions {
		if region.Key == "Australia" {
			return
		}
	}
	t.Errorf("Could not find Australia!?!")
}

func Test_HappyPathGetFlag(t *testing.T) {
	// Given
	expectedFlag := "🇦🇺"
	australianCountry := "Australia"
	service := service.NewRegionService()
	controller := NewRegionController(service)
	// When
	flag := controller.GetEmojiForCountry(australianCountry)
	// Then
	assert.Equal(t, flag, expectedFlag)
}

func Test_GetUnknownFlag(t *testing.T) {
	// Given
	expectedFlag := ""
	unknownCountry := "Not Known"
	service := service.NewRegionService()
	controller := NewRegionController(service)
	// When
	flag := controller.GetEmojiForCountry(unknownCountry)
	// Then
	assert.Equal(t, flag, expectedFlag)
}
