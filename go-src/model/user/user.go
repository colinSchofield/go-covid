package user

type User struct {
	Id      string   `dynamo:"id" json:"id"`
	Name    string   `dynamo:"name" json:"name" binding:"required,max=50"`
	Age     int      `dynamo:"age" json:"age" binding:"required,max=100"`
	Gender  string   `dynamo:"gender" json:"gender" binding:"required"`
	Regions []string `dynamo:"regions" json:"regions"`
	Email   string   `dynamo:"email" json:"email"`
	Sms     string   `dynamo:"sms" json:"sms"`
}
