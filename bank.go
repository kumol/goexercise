package main

import (
	"fmt"
)

var totalBalance int = 1000

func main() {
	fmt.Println("Welcome to the go Bank Application!")
	for {
		process := takeInput("What you want to do? \n1. Check Balance \n2. Deposit Money \n3. Withdraw Money \n4. Exit \nYour choice: ")
		if process == 1 {
			fmt.Println("Checking balance...")
			fmt.Println("Your current balance is: ", totalBalance)
		} else if process == 2 {
			amount := takeInput("Enter the amount to deposit: ")
			fmt.Println("Depositing money...")
			totalBalance = totalBalance + amount
			fmt.Println("New balance is: ", totalBalance)
		} else if process == 3 {
			amount := takeInput("Enter the amount to withdraw: ")
			fmt.Println("Withdrawing money...")
			if amount > totalBalance {
				fmt.Println("Insufficient balance, you can withdraw only ", totalBalance)
				continue
			} else {
				totalBalance = totalBalance - amount
				fmt.Println("New balance is: ", totalBalance)
			}
		} else if process == 4 {
			fmt.Println("Exiting...")
			break
		} else {
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func takeInput(disPlayText string) (selectedValue int) {
	fmt.Print(disPlayText)
	fmt.Scan(&selectedValue)
	return selectedValue
}
