package main

import (
	"fmt"

	"eample.com/struct/user"
)

func main() {
	firstName := getInput("Please enter your first name: ")
	lastName := getInput("Please enter your last name: ")
	birthDate := getInput("Please enter your birth date (YYYY-MM-DD): ")
	fmt.Printf("Hello %s %s, born on %s!\n", firstName, lastName, birthDate)

	var user1 *user.User

	user1, er := user.New(firstName, lastName, birthDate)
	if er != nil {
		fmt.Println(er)
		return
	}
	user1.DisplayUser()
	user1.UpdateUser()
	user1.DisplayUser()
}

func getInput(outPut string) string {
	var input string
	fmt.Println(outPut)
	fmt.Scan(&input)
	return input
}
