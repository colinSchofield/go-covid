package service

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type isoCountryConversion struct {
	Country           string
	ThreeLetterSymbol string
}

func readConversionFile() ([]isoCountryConversion, error) {
	jsonAsString, err := os.ReadFile("test/country-and-iso.json")
	if err != nil {
		return nil, err
	}

	var conversion []isoCountryConversion
	if err := json.Unmarshal([]byte(jsonAsString), &conversion); err != nil {
		return nil, err
	}

	return conversion, nil
}

func Test_RegionISOFromCountryCode(t *testing.T) {

	regionService := NewRegionService()
	regions := regionService.GetListOfRegions()

	if conversion, err := readConversionFile(); err != nil {
		t.Errorf("Error reading the conversion file.. Error was %s", err)
	} else {

		if len(conversion) == 0 {
			t.Error("The conversion file is empty!?")
		}

		for _, country := range conversion {
			isMatched := false
			for _, region := range regions {

				if region.Iso == country.ThreeLetterSymbol {
					isMatched = true
					break
				}
			}
			if !isMatched {
				t.Errorf("Error did not match iso of %s", country.ThreeLetterSymbol)
			}
		}
	}
}

func Test_GetIsoForCountry(t *testing.T) {

	regionService := NewRegionService()

	multiTest := []struct {
		scenario string
		country  string
		iso      string
	}{
		{scenario: "Australia", country: "Australia", iso: "aus"},
		{scenario: "USA", country: "USA", iso: "usa"},
		{scenario: "Canada", country: "Canada", iso: "can"},
		{scenario: "Unknown Location", country: "Unknown", iso: ""},
	}

	for _, test := range multiTest {

		iso := regionService.GetIsoForCountry(test.country)
		assert.Equal(t, iso, test.iso, test.iso, test.scenario)
	}
}
