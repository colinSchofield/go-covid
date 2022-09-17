package model

type Region struct {
	Key         string `json:"key"`
	Location    string `json:"location"`
	CountryCode string `json:"countryCode"`
	Flag        string `json:"flag"`
}
