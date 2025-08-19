package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		//errors.New("Error: All fields are required.")
		return nil, errors.New("Error: All fields are required.")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func (user *User) DisplayUser() {
	fmt.Printf("User Details:\nFirst Name: %s\nLast Name: %s\nBirth Date: %s\nCreated At: %s\n",
		user.firstName, user.lastName, user.birthDate, user.createdAt.Format(time.RFC3339))
}

func (user *User) UpdateUser() {
	user.firstName = ""
	user.lastName = ""
}
