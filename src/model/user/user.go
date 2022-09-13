package user

type User struct {
	Id      string   `dynamo:"id"`
	Name    string   `dynamo:"name"`
	Age     int      `dynamo:"age"`
	Gender  string   `dynamo:"gender"`
	Regions []string `dynamo:"regions"`
	Email   string   `dynamo:"email"`
	Sms     string   `dynamo:"sms"`
}
