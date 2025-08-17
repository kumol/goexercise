package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func newUser(firstName, lastName, birthDate string) *user {
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
}

func main() {

	firstName := getInput("Please enter your first name: ")
	lastName := getInput("Please enter your last name: ")
	birthDate := getInput("Please enter your birth date (YYYY-MM-DD): ")
	fmt.Printf("Hello %s %s, born on %s!\n", firstName, lastName, birthDate)

	var user1 *user
	user1 = newUser(firstName, lastName, birthDate)
	user1.displayUser()
	user1.updateUser()
	user1.displayUser()
}

func (user *user) displayUser() {
	fmt.Printf("User Details:\nFirst Name: %s\nLast Name: %s\nBirth Date: %s\nCreated At: %s\n",
		user.firstName, user.lastName, user.birthDate, user.createdAt.Format(time.RFC3339))
}

func (user *user) updateUser() {
	user.firstName = getInput("Please enter your new first name: ")
	user.lastName = getInput("Please enter your new last name: ")
}

func getInput(outPut string) string {
	var input string
	fmt.Println(outPut)
	fmt.Scan(&input)
	return input
}
