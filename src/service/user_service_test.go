// This is (mostly) an integration test.. 🙈
package service

import (
	"testing"

	"git.com/colinSchofield/go-covid/model/user"
)

func Test_CreatUserAndDelete(t *testing.T) {
	// Given
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
	createdPerson, err := userService.CreateUser(person)
	// Then
	if err != nil {
		t.Errorf("Something went wrong: %v", err)
	}
	if createdPerson.Id == "" {
		t.Errorf("hash was not set!")
	}
	if err := userService.DeleteUser(createdPerson.Id); err != nil {
		t.Errorf("an error occurred during delete: %v", err)
	}
}

func Test_GetListOfAllUsers(t *testing.T) {
	// Given
	userService := NewUserService()
	// When
	users, err := userService.GetListOfAllUsers()
	// Then
	if err != nil {
		t.Errorf("Something went wrong: %v", err)
	}
	if len(users) == 0 {
		t.Error("The returned list is empty?!")
	}
	if len(users[0].Id) == 0 || users[0].Age == 0 {
		t.Error("Values are empty?!")
	}
}

func Test_TestCrudOnUser(t *testing.T) {
	// Given
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
	createdPerson, err := userService.CreateUser(person)
	// Then
	if err != nil {
		t.Errorf("Something went wrong: %v", err)
	}
	if _, err := userService.GetUser(createdPerson.Id); err != nil {
		t.Errorf("Could not find the created user: %v", err)
	}
	if err := userService.DeleteUser(createdPerson.Id); err != nil {
		t.Errorf("User could not be deleted: %v", err)
	}
	if _, err := userService.GetUser(createdPerson.Id); err == nil {
		t.Errorf("This created user, should have been deleted!?")
	}
}

func Test_DecoratedUserRegionList(t *testing.T) {
	// Given
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
		t.Error("Region List information is not correct?")
	}
}

func Test_DecoratedUserContact(t *testing.T) {
	// Given
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
		t.Error("Contact information is not correct?")
	}
}

func Test_DecoratedUserEmptyContact(t *testing.T) {
	// Given
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
		t.Error("Contact information is not correct?")
	}
}
