package main

import "fmt"

// GetChoice prompts the user to enter a choice and returns the entered integer.
// It does not perform any validation on the input.
func GetChoice() int {
	var choice int
	fmt.Print("Your choice: ")
	fmt.Scanln(&choice)
	return choice
}
