package user

type User struct {
	Id      string   `dynamo:"id" json:"id"`
	Name    string   `dynamo:"name" json:"name"`
	Age     int      `dynamo:"age" json:"age"`
	Gender  string   `dynamo:"gender" json:"gender"`
	Regions []string `dynamo:"regions" json:"regions"`
	Email   string   `dynamo:"email" json:"email"`
	Sms     string   `dynamo:"sms" json:"sms"`
}
