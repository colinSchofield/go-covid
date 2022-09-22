package service

import (
	"fmt"

	"git.com/colinSchofield/go-covid/config"
	"git.com/colinSchofield/go-covid/custom_error"
	"git.com/colinSchofield/go-covid/model/user"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/google/uuid"
	"github.com/guregu/dynamo"
)

type UserService interface {
	CreateOrUpdateUser(user user.User) (user.User, error)
	updateUser(person user.User) (user.User, error)
	GetListOfAllUsers() ([]user.DecoratedUser, error)
	GetUser(id string) (user.User, error)
	DeleteUser(id string) error
	getDecoratedUser(user user.User) (user.DecoratedUser, error)
}

type userService struct {
	regionService RegionService
	table         dynamo.Table
}

func NewUserService() UserService {
	regionService := NewRegionService()
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String(config.GetAwsRegion())})
	table := db.Table(config.GetDbTableName())

	return userService{
		regionService: regionService,
		table:         table,
	}
}

func (us userService) CreateOrUpdateUser(person user.User) (user.User, error) {

	if len(person.Id) > 0 {
		config.Logger().Debugf("User already exists, with id of %s.. Updating the user", person.Id)
		return us.updateUser(person)
	}

	config.Logger().Debugf("Creating user, with name of %s", person.Name)
	person.Id = uuid.New().String()
	if err := us.table.Put(person).Run(); err != nil {
		return user.User{}, fmt.Errorf("error creating user: %w", err)
	}
	return person, nil
}

func (us userService) updateUser(person user.User) (user.User, error) {

	var result user.User
	if err := us.table.Update("id", person.Id).
		Set("name", person.Name).
		Set("age", person.Age).
		Set("gender", person.Gender).
		Set("regions", person.Regions).
		Value(&result); err != nil {
		return user.User{}, fmt.Errorf("error updating user: %w", err)
	}
	return result, nil
}

func (us userService) GetListOfAllUsers() ([]user.DecoratedUser, error) {

	var users []user.User
	if err := us.table.Scan().All(&users); err != nil {
		return nil, fmt.Errorf("error reading all users: %w", err)
	}
	var decoratedList []user.DecoratedUser
	for _, user := range users {
		if decorated, err := us.getDecoratedUser(user); err != nil {
			return nil, fmt.Errorf("error calling GetDecoratedUser: %w", err)
		} else {
			decoratedList = append(decoratedList, decorated)
		}
	}
	return decoratedList, nil
}

func (us userService) GetUser(id string) (user.User, error) {

	var result user.User
	if err := us.table.Get("id", id).One(&result); err != nil {
		return user.User{}, custom_error.NotFound{Wrapped: err}
	}
	return result, nil
}

func (us userService) DeleteUser(id string) error {

	if err := us.table.Delete("id", id).Run(); err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}

func (us userService) getDecoratedUser(person user.User) (user.DecoratedUser, error) {
	var regionList string
	var contact string

	for _, country := range person.Regions {
		regionList += us.regionService.GetEmojiForCountry(country) + " "
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
