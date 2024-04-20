package main

import "fmt"

func main() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	// user input
	var choice int
	fmt.Print("Your choice: ")
	fmt.Scanln(&choice)
	fmt.Println("You selected:", choice)
}
