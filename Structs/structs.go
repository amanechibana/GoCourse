package main

import (
	"fmt"
	"time"
)

type user struct{
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

func (u user) outputValue(){
	fmt.Println(u.firstName,u.lastName,u.birthdate,u.createdAt)
}

func (u *user) clearUserName(){
	u.firstName = ""
	u.lastName = ""
}

func main() {
	var appUser user = user{
		getUserData("Please enter your first name: "),
		getUserData("Please enter your last name: "),
		getUserData("Please enter your birthdate (MM/DD/YYYY): "),
		time.Now(),
	}

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
