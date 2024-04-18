package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount = 1000
	var expectedReturn = 5.5
	var years = 10

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturn/100, float64(years))
	fmt.Println("Future value is: $", futureValue)
}
