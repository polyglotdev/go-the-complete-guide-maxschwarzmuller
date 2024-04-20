package main

import "fmt"

func main() {
	// Declare a pointer to an int.
	var ptr *int

	// Declare an int variable.
	num := 3

	// Assign the address of num to ptr.
	ptr = &num

	// Print the value of num.
	fmt.Println("Value of num:", num)

	// Print the address of num.
	fmt.Println("Address of num:", &num)

	// Print the value of ptr.
	fmt.Println("Value of ptr:", ptr)

	// Print the address of ptr.
	fmt.Println("Address of ptr:", &ptr)

	// Print the value of num using ptr.
	fmt.Println("Value of num using ptr:", *ptr)

	age := 30
	fmt.Println("Age before:", age)
	alterAge(&age)
	fmt.Println("Age after:", age)
}

// alterAge is a function that takes a pointer to an integer as an argument.
// It changes the value of the integer pointed to by the pointer to 25.
// It then returns the new value.
func alterAge(age *int) int {
	*age = 25
	return *age
}
