package user

import (
	"fmt"
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

type Admin struct {
	email    string
	password string
	User
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

// CreateUser is a function that prompts the user for their first name, last name, and birthdate.
// It then creates a new User struct with these values and the current time as the CreatedAt value.
// It does not take any arguments and returns a User struct.
func New() (*User, error) {
	firstName := GetUserData("Please enter your first name: ", false)
	// Prompt the user for their last name. The second argument is false because a last name is not a date.
	lastName := GetUserData("Please enter your last name: ", false)
	// Prompt the user for their birthdate. The second argument is true because a birthdate is a date.
	birthdate := GetUserData("Please enter your birthdate (MM/DD/YYYY): ", true)

	// Return a new User struct with the provided first name, last name, birthdate, and the current time.
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now(),
	}, nil
}

// NewAdmin is a constructor function for creating a new Admin object.
// It takes an email and a password as parameters and returns a pointer to the newly created Admin object.
// The email and password are used to set the respective fields of the Admin object.
func NewAdmin(email, password, firstName, lastName, birthdate string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			FirstName: firstName,
			LastName:  lastName,
			Birthdate: birthdate,
			CreatedAt: time.Now(),
		},
	}
}

// GetUserData is a function that prompts the user for input and validates it.
// It takes two parameters: a string for the prompt text and a boolean to indicate whether the input should be a date.
// If the input is empty, it prompts the user again.
// If the input should be a date, it checks whether the input matches the date format (MM/DD/YYYY).
// If the input should not be a date, it checks whether the input is a number and prompts the user again if it is.
// The function returns the input in title case.
func GetUserData(promptText string, isDate bool) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	if value == "" {
		fmt.Println("Invalid input. Please try again.")
		return GetUserData(promptText, isDate)
	}

	if isDate {
		match, _ := regexp.MatchString(`\d{2}/\d{2}/\d{4}`, value)
		if !match {
			fmt.Println("Invalid date format. Please try again.")
			return GetUserData(promptText, isDate)
		}
	} else {
		if _, err := strconv.Atoi(value); err == nil {
			fmt.Println("Invalid input. Name should not be a number. Please try again.")
			return GetUserData(promptText, isDate)
		}
	}

	titleCase := cases.Title(language.English)
	return titleCase.String(value)
}
