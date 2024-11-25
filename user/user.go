package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

type Admin struct{
	email string
	password string
	User
}

func NewAdmin(email, password string) Admin{
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "Massab",
			lastName: "Rasheed",
			birthDate: "null",
			createdAt: time.Now(),
		},
	}
}

// (u user) is a Receiver (argument)
func (u User) OutputUserData(){
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

//Mutation requires a pointer in Receiver Argument
func (u *User) ClearUserName (){
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}
	return &User{
		firstName: firstName,
		lastName: lastName,
		birthDate:birthDate,
		createdAt: time.Now(),
	}, nil
}