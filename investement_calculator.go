package main

import (
	"fmt"
	"math"
)

const (
	inflationRate = 2.5
)

func main() {
	investmentAmount, expectedReturn, years := 1000.00, 5.5, 10.0

	var futureValue = investmentAmount * math.Pow(1+expectedReturn/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Printf("Future value is: $%.2f\n", futureValue)
	fmt.Printf("Future real value is: $%.2f\n", futureRealValue)
}
