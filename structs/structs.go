package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// User is a struct that represents a user in the system.
type User struct {
	FirstName string
	LastName  string
	Birthdate string
	CreatedAt time.Time
}

// PrintUser is a method on the User struct that prints the user's details.
// It prints the user's first name, last name, birthdate, and the date the user was created.
func (u User) PrintUser() {
	fmt.Printf("User First Name: %s\n", u.FirstName)
	fmt.Printf("User Last Name: %s\n", u.LastName)
	fmt.Printf("User Birthdate: %s\n", u.Birthdate)
	fmt.Printf("User Created At: %s\n", u.CreatedAt)
	fmt.Println("---------------------------------------------------------------------")
}

func main() {
	elijah, err := createUser()
	if err != nil {
		log.Fatal(err)
	}
	elijah.PrintUser()

	ezra, err := createUser()
	if err != nil {
		log.Fatal(err)
	}
	ezra.PrintUser()
}

// createUser is a function that prompts the user for their first name, last name, and birthdate.
// It then creates a new User struct with these values and the current time as the CreatedAt value.
// It does not take any arguments and returns a User struct.
func createUser() (*User, error) {
	firstName := getUserData("Please enter your first name: ", false)
	// Prompt the user for their last name. The second argument is false because a last name is not a date.
	lastName := getUserData("Please enter your last name: ", false)
	// Prompt the user for their birthdate. The second argument is true because a birthdate is a date.
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ", true)

	// Return a new User struct with the provided first name, last name, birthdate, and the current time.
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now(),
	}, nil
}

// getUserData is a function that prompts the user for input and validates it.
// It takes two parameters: a string for the prompt text and a boolean to indicate whether the input should be a date.
// If the input is empty, it prompts the user again.
// If the input should be a date, it checks whether the input matches the date format (MM/DD/YYYY).
// If the input should not be a date, it checks whether the input is a number and prompts the user again if it is.
// The function returns the input in title case.
func getUserData(promptText string, isDate bool) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	if value == "" {
		fmt.Println("Invalid input. Please try again.")
		return getUserData(promptText, isDate)
	}

	if isDate {
		match, _ := regexp.MatchString(`\d{2}/\d{2}/\d{4}`, value)
		if !match {
			fmt.Println("Invalid date format. Please try again.")
			return getUserData(promptText, isDate)
		}
	} else {
		if _, err := strconv.Atoi(value); err == nil {
			fmt.Println("Invalid input. Name should not be a number. Please try again.")
			return getUserData(promptText, isDate)
		}
	}

	titleCase := cases.Title(language.English)
	return titleCase.String(value)
}
