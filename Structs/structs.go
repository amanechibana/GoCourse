package main

import (
	"fmt"
	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirth := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirth)

	if err != nil {
		fmt.Println(err)
		return
	}

	Admin := user.NewAdmin("test@gmail.com","thelog","pog","log","24/24/2424")

	Admin.OutputValue()

	appUser.OutputValue()
	appUser.ClearUserName()
	appUser.OutputValue()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
