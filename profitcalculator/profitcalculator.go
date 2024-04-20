package main

import "fmt"

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: $%.2f\n", ebt)
	fmt.Printf("Profit: $%.2f\n", profit)
	fmt.Printf("Ratio: %.3f%%\n", ratio)
}

// calculateFinancials calculates the earnings before tax (EBT), profit after tax, and the ratio of EBT to profit.
// It takes three parameters: revenue, expenses, and taxRate.
// Revenue and expenses are the total income and costs respectively, and taxRate is the percentage of tax.
// It returns three float64 values: EBT, profit, and the ratio of EBT to profit.
func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses         // calculate earnings before tax
	profit := ebt * (1 - taxRate/100) // calculate profit after tax
	ratio := ebt / profit             // calculate the ratio of EBT to profit
	return ebt, profit, ratio
}

// getUserInput is a function that prompts the user for input.
// It takes a string argument, infoText, which is displayed as a prompt to the user.
// The function then scans the user's input and returns it as a float64.
func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}
