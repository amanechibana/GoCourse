package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u user) outputValue() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthdate string) *user {
	if firstName == "" || lastName == "" || birthdate == "" {

	}

	return &user{
		firstName,
		lastName,
		birthdate,
		time.Now(),
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirth := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user = newUser(userFirstName, userLastName, userBirth)

	appUser.outputValue()
	appUser.clearUserName()
	appUser.outputValue()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
