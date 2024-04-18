package main

import (
	"fmt"
	"math"
)

func main() {
	investmentAmount := 1000.00
	expectedReturn := 5.5
	years := 10.0

	var futureValue = investmentAmount * math.Pow(1+expectedReturn/100, years)
	fmt.Printf("Future value is: $%.2f\n", futureValue)
}
