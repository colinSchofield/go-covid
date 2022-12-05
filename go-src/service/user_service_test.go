// This is (mostly) an integration test.. 🙈
package service

import (
	"testing"

	"github.com/colinSchofield/go-covid/config"
	"github.com/colinSchofield/go-covid/model/user"
)

func setEnvironmentVariables(t *testing.T) {
	t.Setenv(config.AWS_REGION, "ap-southeast-2")
	t.Setenv(config.DB_TABLE_NAME, "User")
}

func Test_CreatUserAndDelete(t *testing.T) {
	// Given
	setEnvironmentVariables(t)
	userService := NewUserService()
	person := user.User{
		Id:      "", // this value will be set as a hash
		Name:    "Test Person",
		Gender:  "Male",
		Regions: []string{"Here", "There"},
		Email:   "tp@pt.com",
		Sms:     "123-432-1233",
	}
	// When
	createdPerson, err := userService.CreateOrUpdateUser(person)
	// Then
	if err != nil {
		t.Errorf("something went wrong: %s", err)
	}
	if createdPerson.Id == "" {
		t.Errorf("hash was not set!")
	}
	if err := userService.DeleteUser(createdPerson.Id); err != nil {
		t.Errorf("an error occurred during delete: %s", err)
	}
}

func Test_GetListOfAllUsers(t *testing.T) {
	// Given
	setEnvironmentVariables(t)
	userService := NewUserService()
	// When
	users, err := userService.GetListOfAllUsers()
	// Then
	if err != nil {
		t.Errorf("something went wrong: %s", err)
	}
	if len(users) == 0 {
		t.Error("the returned list is empty?!")
	}
	if len(users[0].Id) == 0 || users[0].Age == 0 {
		t.Error("values are empty?!")
	}
}

func Test_TestCrudOnUser(t *testing.T) {
	// Given
	setEnvironmentVariables(t)
	userService := NewUserService()
	person := user.User{
		Id:      "", // this value will be set as a hash
		Name:    "Test Person",
		Gender:  "Male",
		Age:     0,
		Regions: []string{"Here", "There"},
		Email:   "tp@pt.com",
		Sms:     "123-432-1233",
	}
	// When
	createdPerson, err := userService.CreateOrUpdateUser(person)
	// Then
	if err != nil {
		t.Errorf("something went wrong: %s", err)
	}
	// Given
	createdPerson.Age = 66
	// When
	updatedPerson, err := userService.CreateOrUpdateUser(createdPerson)
	// Then
	if err != nil {
		t.Errorf("something went wrong: %s", err)
	}
	if updatedPerson.Age != 66 {
		t.Errorf("age was not updated: %d", createdPerson.Age)
	}
	if _, err := userService.GetUser(createdPerson.Id); err != nil {
		t.Errorf("could not find the created user: %s", err)
	}
	if err := userService.DeleteUser(createdPerson.Id); err != nil {
		t.Errorf("user could not be deleted: %s", err)
	}
	if _, err := userService.GetUser(createdPerson.Id); err == nil {
		t.Errorf("this created user, should have been deleted!?")
	}
}

func Test_DecoratedUserRegionList(t *testing.T) {
	// Given
	setEnvironmentVariables(t)
	userService := NewUserService()
	person := user.User{
		Id:      "",
		Name:    "Test Person",
		Age:     0,
		Gender:  "Male",
		Regions: []string{"Australia", "Canada", "UK"},
		Email:   "",
		Sms:     "",
	}
	// When
	decoratedUser, _ := userService.GetDecoratedUser(person)
	// Then
	if decoratedUser.RegionList != "🇦🇺 🇨🇦 🇬🇧 " {
		t.Error("Region List information is not correct?")
	}
}

func Test_DecoratedUserRegionListUnknown(t *testing.T) {
	// Given
	setEnvironmentVariables(t)
	userService := NewUserService()
	person := user.User{
		Id:      "",
		Name:    "Test Person",
		Age:     0,
		Gender:  "Male",
		Regions: []string{"Here", "There", "Where?"},
		Email:   "",
		Sms:     "",
	}
	// When
	decoratedUser, _ := userService.GetDecoratedUser(person)
	// Then
	if decoratedUser.RegionList != "   " {
		t.Error("region List information is not correct?")
	}
}

func Test_DecoratedUserContact(t *testing.T) {
	// Given
	setEnvironmentVariables(t)
	userService := NewUserService()
	person := user.User{
		Id:      "",
		Name:    "Test Person",
		Age:     0,
		Gender:  "Male",
		Regions: []string{"Here", "There"},
		Email:   "tp@pt.com",
		Sms:     "123-432-1233",
	}
	// When
	decoratedUser, _ := userService.GetDecoratedUser(person)
	// Then
	if decoratedUser.Contact != "💌 💬" {
		t.Error("contact information is not correct?")
	}
}

func Test_DecoratedUserEmptyContact(t *testing.T) {
	// Given
	setEnvironmentVariables(t)
	userService := NewUserService()
	person := user.User{
		Id:      "",
		Name:    "Test Person",
		Age:     0,
		Gender:  "Male",
		Regions: []string{"Here", "There"},
		Email:   "",
		Sms:     "",
	}
	// When
	decoratedUser, _ := userService.GetDecoratedUser(person)
	// Then
	if decoratedUser.Contact != "" {
		t.Error("contact information is not correct?")
	}
}
