package main

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Birthdate string
	CreatedAt time.Time
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!
	user := User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now(),
	}

	fmt.Printf("User First Name: %s\n", user.FirstName)
	fmt.Printf("User Last Name: %s\n", user.LastName)
	fmt.Printf("User Birthdate: %s\n", user.Birthdate)
	fmt.Printf("User Created At: %s\n", user.CreatedAt)
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
