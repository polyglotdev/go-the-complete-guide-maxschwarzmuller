package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Birthdate string
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	fmt.Println(firstName, lastName, birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	if value == "" {
		fmt.Println("Invalid input. Please try again.")
		return getUserData(promptText)
	}
	return value
}
