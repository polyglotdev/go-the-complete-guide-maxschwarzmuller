package main

import (
	"fmt"

	"example.com/collections/array"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	total := array.Total(numbers)
	fmt.Printf("The total is: %d\n", total)

	prices := []float64{10.99, 5.99, 3.99, 7.99}
	totalPrice := array.TotalFloat64(prices[:])
	fmt.Printf("The total price is: $%.2f\n", totalPrice)

	average := array.Average(numbers)
	fmt.Printf("The average is: %.2f\n", average)

	averagePrice := array.AverageFloat64(prices[:])
	fmt.Printf("The average price is: $%.2f\n", averagePrice)
}
