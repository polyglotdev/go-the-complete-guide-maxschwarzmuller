package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	result := calculate(prices, taxRates)
	fmt.Printf("result: %v\n", result)
}

// calculate is a function that takes two slices of float64 as input: prices and taxRates.
// It calculates the total tax rate by summing up all the tax rates in the taxRates slice.
// Then, it applies this total tax rate to each price in the prices slice by adding the product of the price and the total tax rate to the original price.
// The function returns the modified prices slice.
//
// Parameters:
// prices: A slice of float64 representing the prices of items.
// taxRates: A slice of float64 representing the tax rates for each item.
//
// Returns:
// A slice of float64 representing the prices of items after applying the total tax rate.
func calculate(prices []float64, taxRates []float64) []float64 {
	// Calculate total tax rate
	totalTaxRate := 0.0
	for _, rate := range taxRates {
		totalTaxRate += rate
	}

	// Apply total tax rate to each price
	for i, price := range prices {
		prices[i] = price + (price * totalTaxRate)
	}

	return prices
}
