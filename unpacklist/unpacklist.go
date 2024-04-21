package main

import "fmt"

func main() {
	prices := []float64{100.45, 10046.89, 850000.20, 67643.83}
	fmt.Printf("Prices Header: %T\nCapacity: %d\nLength: %d\nPrices: %v\n", prices, cap(prices), len(prices), prices)
	fmt.Println("----------------------------------------------------------------------")
	prices = append(prices, 3458.98, 576.23)
	fmt.Println("Prices:", prices)

	prices = prices[1:]
	fmt.Println("Prices after re-slice:", prices)
}
