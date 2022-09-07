package controller

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"git.com/gol/model"
	"git.com/gol/service"
	"github.com/gin-gonic/gin"
)

func Test_HappyPathAllFlags(t *testing.T) {
	// Given
	service := service.New()
	controller := New(service)
	record := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(record)
	// When
	controller.GetListOfRegions(context)
	// Then
	if record.Code != 200 {
		t.Errorf("Unexpected HTTP return code. Expected 200, but got %d", record.Code)
	}
	regions := []model.Region{}
	if err := json.Unmarshal(record.Body.Bytes(), &regions); err != nil {
		t.Errorf("String could not be unmarshalled? %s", err)
	}
	if len(regions) == 0 {
		t.Errorf("Region object is empty!")
	}
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
	australianCountry := "AU"
	service := service.New()
	controller := New(service)
	// When
	flag := controller.GetEmojiForCountry(australianCountry)
	// Then
	if flag != expectedFlag {
		t.Errorf("Flag did not match! Was expecting %s, but got %s", expectedFlag, flag)
	}
}

func Test_GetUnknownFlag(t *testing.T) {
	// Given
	expectedFlag := ""
	unknownCountry := "Not Known"
	service := service.New()
	controller := New(service)
	// When
	flag := controller.GetEmojiForCountry(unknownCountry)
	// Then
	if flag != expectedFlag {
		t.Errorf("Flag did not match! Was expecting %s, but got %s", expectedFlag, flag)
	}
}

func Test_regionController_GetListOfRegions(t *testing.T) {
	type fields struct {
		regionService service.RegionService
	}
	type args struct {
		context *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := regionController{
				regionService: tt.fields.regionService,
			}
			controller.GetListOfRegions(tt.args.context)
		})
	}
}
