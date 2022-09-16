package user

type DecoratedUser struct {
	User
	RegionList string `json:"regionList"`
	Contact    string `json:"contact"`
}
