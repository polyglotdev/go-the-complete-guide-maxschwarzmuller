package main

import (
	"os"
	"strconv"
)

// ReadBalance reads the balance from a file named "balance.txt".
// If the file does not exist or its content is not a valid float,
// it returns a default balance of 1000.00.
func ReadBalance() float64 {
	// Read the file "balance.txt".
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		// If the file doesn't exist, return the default balance.
		return 1000.00
	}
	// Parse the file content to a float.
	balance, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		// If the file content is not a valid float, return the default balance.
		return 1000.00
	}
	// Return the balance.
	return balance
}
