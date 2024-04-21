package main

import "fmt"

func main() {
	prices := []float64{100.45, 10046.89, 850000.20, 67643.83}
	fmt.Printf("Prices Header: %T\nCapacity: %d\nLength: %d\nPrices: %v\n", prices, cap(prices), len(prices), prices)
	fmt.Println("----------------------------------------------------------------------")
	prices = append(prices, 3458.98, 576.23)
	fmt.Printf("Prices Header: %T\nCapacity: %d\nLength: %d\nPrices: %v\n", prices, cap(prices), len(prices), prices)
	fmt.Println("----------------------------------------------------------------------")

	prices = prices[1:]
	fmt.Println("Prices after re-slice:", prices)
	fmt.Printf("Prices Header: %T\nCapacity: %d\nLength: %d\nPrices: %v\n", prices, cap(prices), len(prices), prices)
	fmt.Println("----------------------------------------------------------------------")

	discountPrices := []float64{101.66, 1056.29, 850.50, 6.83}
	prices = append(prices, discountPrices...)
	fmt.Printf("Prices Header: %T\nCapacity: %d\nLength: %d\nPrices: %v\n", prices, cap(prices), len(prices), prices)
	fmt.Println("----------------------------------------------------------------------")
	// Start with a slice of length 0 and capacity 0.
	s := make([]int, 0)
	fmt.Printf("Capacity: %d, Length: %d\n", cap(s), len(s))

	// Append an element to the slice.
	s = append(s, 1)
	fmt.Printf("Capacity: %d, Length: %d\n", cap(s), len(s))

	// Keep appending elements until the capacity reaches 1024.
	for len(s) < 1024 {
		s = append(s, 1)
	}
	fmt.Printf("Capacity: %d, Length: %d\n", cap(s), len(s))

	// Append one more element.
	s = append(s, 1)
	fmt.Printf("Capacity: %d, Length: %d\n", cap(s), len(s))
}
