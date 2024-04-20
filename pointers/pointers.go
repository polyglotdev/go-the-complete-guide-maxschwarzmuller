package main

func main() {
	// Declare a pointer to an int.
	var ptr *int

	// Declare an int variable.
	num := 3

	// Assign the address of num to ptr.
	ptr = &num

	// Print the value of num.
	println("Value of num:", num)

	// Print the address of num.
	println("Address of num:", &num)

	// Print the value of ptr.
	println("Value of ptr:", ptr)

	// Print the address of ptr.
	println("Address of ptr:", &ptr)

	// Print the value of num using ptr.
	println("Value of num using ptr:", *ptr)
}
