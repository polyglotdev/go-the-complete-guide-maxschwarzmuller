package main

import "fmt"

func main() {
	prices := []float64{100.45, 10046.89, 850000.20, 67643.83}
	prices = append(prices, 3458.98, 576.23)
	fmt.Println(prices)
}
