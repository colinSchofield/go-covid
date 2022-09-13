package service

import (
	"fmt"

	"git.com/colinSchofield/go-covid/model/user"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/google/uuid"
	"github.com/guregu/dynamo"
)

type UserService interface {
	CreateUser(user user.User) (user.User, error)
	UpdateUser(user user.User) (user.User, error)
	GetListOfAllUsers() ([]user.DecoratedUser, error)
	GetUser(id string) (user.User, error)
	DeleteUser(id string) error
	GetDecoratedUser(user user.User) (user.DecoratedUser, error)
}

type userService struct {
	regionService RegionService
	table         dynamo.Table
}

func NewUserService() UserService {
	regionService := NewRegionService()
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String("ap-southeast-2")}) // TODO add environment variable
	table := db.Table("User")                                                 // TODO add environment variable

	return userService{
		regionService: regionService,
		table:         table,
	}
}

func (userService userService) CreateUser(person user.User) (user.User, error) {

	person.Id = uuid.New().String()
	if err := userService.table.Put(person).Run(); err != nil {
		return user.User{}, fmt.Errorf("error creating user: %w", err)
	}
	return person, nil
}

func (userService userService) UpdateUser(person user.User) (user.User, error) {

	var result user.User
	if err := userService.table.Update("id", person.Id).
		Set("name", person.Name).
		Set("age", person.Age).
		Set("gender", person.Gender).
		Set("regions", person.Regions).
		Value(&result); err != nil {
		return user.User{}, fmt.Errorf("error updating user: %w", err)
	}
	return result, nil
}

func (userService userService) GetListOfAllUsers() ([]user.DecoratedUser, error) {

	var users []user.User
	if err := userService.table.Scan().All(&users); err != nil {
		return nil, fmt.Errorf("error reading all users: %w", err)
	}
	var decoratedList []user.DecoratedUser
	for _, user := range users {
		if decorated, err := userService.GetDecoratedUser(user); err != nil {
			return nil, fmt.Errorf("error calling GetDecoratedUser: %w", err)
		} else {
			decoratedList = append(decoratedList, decorated)
		}
	}
	return decoratedList, nil
}

func (userService userService) GetUser(id string) (user.User, error) {

	var result user.User
	if err := userService.table.Get("id", id).One(&result); err != nil {
		return user.User{}, fmt.Errorf("error reading user: %w", err)
	}
	return result, nil
}

func (userService userService) DeleteUser(id string) error {

	if err := userService.table.Delete("id", id).Run(); err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}

func (userService userService) GetDecoratedUser(person user.User) (user.DecoratedUser, error) {
	var regionList string
	var contact string

	for _, country := range person.Regions {
		regionList += userService.regionService.GetEmojiForCountry(country) + " "
	}

	if len(person.Email) > 0 {
		contact += "💌 "
	}
	if len(person.Sms) > 0 {
		contact += "💬"
	}

	decorated := user.DecoratedUser{
		User:       person,
		RegionList: regionList,
		Contact:    contact,
	}

	return decorated, nil
}
