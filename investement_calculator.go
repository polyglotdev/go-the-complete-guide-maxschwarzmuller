package main

import (
	"fmt"
	"math"
)

const (
	inflationRate = 2.5
)

func main() {
	investmentAmount, expectedReturn, years := getUserInput()

	futureValue := calculateFutureValue(investmentAmount, expectedReturn, years)
	futureRealValue := calculateFutureRealValue(futureValue, years)

	fmt.Printf("Future value is: $%.2f\n", futureValue)
	fmt.Printf("Future real value is: $%.2f\n", futureRealValue)
}

// getUserInput is a function that prompts the user to input their investment details.
// It returns three float64 values: the investment amount, the expected return rate, and the number of years.
func getUserInput() (investmentAmount, expectedReturn, years float64) {
	// Prompt the user for the investment amount
	fmt.Print("Investment amount: ")
	_, err := fmt.Scan(&investmentAmount)
	// If there's an error reading the input, print an error message
	if err != nil {
		fmt.Println("Error reading investment amount:", err)
	}

	// Prompt the user for the expected return rate
	fmt.Print("Expected Return Rate: ")
	_, err = fmt.Scan(&expectedReturn)
	// If there's an error reading the input, print an error message
	if err != nil {
		fmt.Println("Error reading expected return rate:", err)
	}

	// Prompt the user for the number of years
	fmt.Print("Years: ")
	_, err = fmt.Scan(&years)
	// If there's an error reading the input, print an error message
	if err != nil {
		fmt.Println("Error reading years:", err)
	}

	// Return the investment details
	return
}

func calculateFutureValue(investmentAmount, expectedReturn, years float64) float64 {
	return investmentAmount * math.Pow(1+expectedReturn/100, years)
}

func calculateFutureRealValue(futureValue, years float64) float64 {
	return futureValue / math.Pow(1+inflationRate/100, years)
}
