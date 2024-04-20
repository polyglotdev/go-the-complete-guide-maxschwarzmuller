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
	balance := 1000.00
	var depositedAmount float64

	switch choice {
	case 1:
		fmt.Printf("Your balance is: $%.2f\n", balance)
	case 2:
		fmt.Print("Enter amount to deposit: ")
		fmt.Scanln(&depositedAmount)
		balance += depositedAmount
		fmt.Printf("Your balance is: $%.2f\n", balance)
	case 3:
		fmt.Println("Withdraw money")
		fmt.Print("Enter amount to withdraw: ")
		fmt.Scanln(&depositedAmount)
		if depositedAmount > balance {
			fmt.Println("Insufficient balance")
		} else {
			balance -= depositedAmount
			fmt.Printf("Your balance is: $%.2f\n", balance)
		}
	case 4:
		fmt.Println("Exit")
	default:
		fmt.Println("Invalid choice")
	}
}
