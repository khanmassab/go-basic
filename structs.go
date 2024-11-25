package main

import (
	"errors"
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

func newUser(firstName, lastName, birthDate string) (*user, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, Last name and Birthdate are required")
	}
	return &user{
		firstName: firstName,
		lastName: lastName,
		birthDate:birthDate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user
	appUser, err := newUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.outputUserData()
	appUser.clearUserName()
	appUser.outputUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
