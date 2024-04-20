package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to Go Bank!")
	balance := 1000.00

	for {
		printMenu()
		choice := getChoice()
		processChoice(choice, &balance)
	}
}

// printMenu is a function that displays the main menu of the Go Bank application.
// It provides the user with the following options:
//
// 1. Check Balance: This option allows the user to view their current account balance.
//
// 2. Deposit Money: This option allows the user to add funds to their account.
//
// 3. Withdraw Money: This option allows the user to remove funds from their account.
//
// 4. Exit: This option allows the user to close the application.
func printMenu() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}

// getChoice prompts the user to enter a choice and returns the entered integer.
// It does not perform any validation on the input.
func getChoice() int {
	var choice int
	fmt.Print("Your choice: ")
	fmt.Scanln(&choice)
	return choice
}

// getAmount prompts the user to enter an amount and returns it as a float64.
// It does not perform any error checking on the input.
func getAmount() float64 {
	var amount float64
	fmt.Print("Enter amount: ")
	fmt.Scanln(&amount)
	return amount
}

// processChoice is a function that takes in a choice and a pointer to a balance.
// It performs different operations based on the choice provided.
// Choice 1: Prints the current balance.
//
// Choice 2: Asks for an amount to deposit, adds it to the balance and prints the updated balance.
//
// Choice 3: Asks for an amount to withdraw, checks if the balance is sufficient,
//
//	if so, deducts it from the balance and prints the updated balance.
//
// Choice 4: Exits the program.
// Any other choice: Prints an error message.
func processChoice(choice int, balance *float64) {
	switch choice {
	case 1:
		fmt.Printf("Your balance is: $%.2f\n", *balance)
	case 2:
		depositedAmount := getAmount()
		*balance += depositedAmount
		fmt.Printf("Your balance is: $%.2f\n", *balance)
	case 3:
		withdrawAmount := getAmount()
		if withdrawAmount > *balance {
			fmt.Println("Insufficient balance")
		} else {
			*balance -= withdrawAmount
			fmt.Printf("Your balance is: $%.2f\n", *balance)
		}
	case 4:
		fmt.Println("Goodbye! ⭐")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
	}
}
