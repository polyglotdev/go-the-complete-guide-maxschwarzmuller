package main

import (
	"fmt"
	"os"
)

// writeBalance writes the given balance to a file named "balance.txt".
// The balance is formatted to two decimal places.
// It returns an error if the write operation fails.
func WriteBalance(balance float64) error {
	// Attempt to write the balance to the file.
	err := os.WriteFile("balance.txt", []byte(fmt.Sprintf("%.2f", balance)), 0644)
	// If an error occurred, return it.
	if err != nil {
		return err
	}
	// If no error occurred, return nil.
	return nil
}
