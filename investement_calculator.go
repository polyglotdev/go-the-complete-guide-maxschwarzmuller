package main

import (
	"fmt"
	"math"
)

const (
	inflationRate = 2.5
)

func main() {
	var investmentAmount float64
	var expectedReturn float64
	var years float64

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturn)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturn/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future value is: $%.2f\n", futureValue)
	fmt.Printf("Future real value is: $%.2f\n", futureRealValue)
}
