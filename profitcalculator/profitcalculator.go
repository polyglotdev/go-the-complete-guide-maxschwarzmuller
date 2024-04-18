package main

import "fmt"

const (
	taxRate = 30
)

func main() {
	var revenue float64
	var expenses float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	profit := revenue - expenses
	tax := profit * taxRate / 100
	netProfit := profit - tax

	fmt.Printf("Profit: $%.2f\n", profit)
	fmt.Printf("Tax: $%.2f\n", tax)
	fmt.Printf("Net Profit: $%.2f\n", netProfit)
}
