package history

type History struct {
	Country   string
	Date      string `json:"date"`
	NewCases  int    `json:"new_cases"`
	NewDeaths int    `json:"new_deaths"`
}
