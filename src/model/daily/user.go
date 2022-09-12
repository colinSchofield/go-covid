package daily

type User struct {
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	Age         int         `json:"age"`
	Email       string      `json:"email"`
	Phone       string      `json:"phone"`
	CountryInfo CountryInfo `json:"countryInfo"`
}
