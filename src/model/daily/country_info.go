package daily

type CountryInfo struct {
	Country          string `json:"country"`
	DecoratedCountry string `json:"decoratedCountry"`
	Cases            Cases  `json:"cases"`
	Deaths           Deaths `json:"deaths"`
	Tests            Tests  `json:"tests"`
	Day              string `json:"day"`
	Time             string `json:"time"`
}
