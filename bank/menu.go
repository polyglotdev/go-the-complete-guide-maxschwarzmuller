package main

import "fmt"

// PrintMenu is a function that displays the main menu of the Go Bank application.
// It provides the user with the following options:
//
// 1. Check Balance: This option allows the user to view their current account balance.
//
// 2. Deposit Money: This option allows the user to add funds to their account.
//
// 3. Withdraw Money: This option allows the user to remove funds from their account.
//
// 4. Exit: This option allows the user to close the application.
func PrintMenu() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}
