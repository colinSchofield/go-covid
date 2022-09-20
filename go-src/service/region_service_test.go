package service

import (
	"encoding/json"
	"os"
	"testing"
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
