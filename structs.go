package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

// (u user) is a Receiver (argument)
func (u user) outputUserData(){
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

//Mutation requires a pointer in Receiver Argument
func (u *user) clearUserName (){
	u.firstName = ""
	u.lastName = ""
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user = user{
		firstName: userFirstName,
		lastName: userLastName,
		birthDate:userBirthdate,
		createdAt: time.Now(),
	}

	appUser.outputUserData()
	appUser.clearUserName()
	appUser.outputUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
