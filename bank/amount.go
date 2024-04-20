package main

import "fmt"

// GetAmount prompts the user to enter an amount and returns it as a float64.
// It does not perform any error checking on the input.
func GetAmount() float64 {
	var amount float64
	fmt.Print("Enter amount: ")
	fmt.Scanln(&amount)
	return amount
}
