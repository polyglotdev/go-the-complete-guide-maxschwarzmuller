package main

import (
	"fmt"

	"example.com/calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		priceJob := prices.NewPriceJobWithTax(taxRate)
		priceJob.Process()
	}
	fmt.Printf("Result: %v\n", result)
}
