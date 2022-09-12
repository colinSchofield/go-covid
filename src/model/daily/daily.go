package daily

type Daily struct {
	Parameters []string      `json:"parameters"`
	Errors     []string      `json:"errors"`
	Results    int           `json:"results"`
	Response   []CountryInfo `json:"response"`
}
