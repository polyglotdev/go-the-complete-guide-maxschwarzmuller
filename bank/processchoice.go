package main

import (
	"fmt"
	"os"
)

// ProcessChoice is a function that takes in a choice and a pointer to a balance.
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
func ProcessChoice(choice int, balance *float64) {
	switch choice {
	case 1:
		fmt.Printf("Your balance is: $%.2f\n", *balance)
	case 2:
		depositedAmount := GetAmount()
		*balance += depositedAmount
		fmt.Printf("Your balance is: $%.2f\n", *balance)
		err := WriteBalance(*balance)
		if err != nil {
			fmt.Println("Error writing balance:", err)
		}
	case 3:
		withdrawAmount := GetAmount()
		if withdrawAmount > *balance {
			fmt.Println("Insufficient balance")
		} else {
			*balance -= withdrawAmount
			fmt.Printf("Your balance is: $%.2f\n", *balance)
			err := WriteBalance(*balance)
			if err != nil {
				fmt.Println("Error writing balance:", err)
			}
		}
	case 4:
		fmt.Println("Goodbye! ⭐")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
	}
}
