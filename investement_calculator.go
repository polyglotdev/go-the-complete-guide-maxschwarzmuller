package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000.00
	var expectedReturn = 5.5
	var years = 10

	var futureValue = investmentAmount * math.Pow(1+expectedReturn/100, float64(years))
	fmt.Println("Future value is: $", futureValue)
}
