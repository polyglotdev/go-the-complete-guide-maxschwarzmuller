package main

import (
	"fmt"
	"os"
)

func main() {
	balance := 1000.00
	var choice int
	var depositedAmount float64

	for {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		fmt.Print("Your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is: $%.2f\n", balance)
		case 2:
			fmt.Print("Enter amount to deposit: ")
			fmt.Scanln(&depositedAmount)
			balance += depositedAmount
			fmt.Printf("Your balance is: $%.2f\n", balance)
		case 3:
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
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	}
}
