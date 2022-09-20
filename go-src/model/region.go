package model

type Region struct {
	Key         string `json:"key"`
	Location    string `json:"location"`
	CountryCode string `json:"countryCode"`
	Iso         string `json:"iso"`
	Flag        string `json:"flag"`
}
