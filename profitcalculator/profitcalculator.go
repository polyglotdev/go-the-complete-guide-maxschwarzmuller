package main

import "fmt"

func main() {
	revenue, expenses, taxRate := getUserInput()

	ebt := calculateEBT(revenue, expenses)
	profit := calculateProfit(ebt, taxRate)
	tax := calculateTax(ebt, taxRate)
	ratio := calculateRatio(ebt, profit)

	fmt.Printf("Profit: $%.2f\n", profit)
	fmt.Printf("Tax: $%.2f\n", tax)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

// calculateEBT calculates the Earnings Before Tax (EBT) for a business.
// It subtracts the total expenses from the total revenue.
// Parameters:
// revenue: The total revenue of the business.
// expenses: The total expenses of the business.
// Returns:
// The Earnings Before Tax (EBT) as a float64.
func calculateEBT(revenue, expenses float64) float64 {
	return revenue - expenses
}

// calculateProfit calculates the net profit by subtracting the tax from the earnings before tax (EBT).
// It takes two parameters: ebt and taxRate.
// ebt represents the earnings before tax.
// taxRate is the tax rate applied to the earnings. It's a percentage value, so it's divided by 100 in the calculation.
// The function returns the net profit as a float64.
func calculateProfit(ebt, taxRate float64) float64 {
	return ebt * (1 - taxRate/100)
}

// calculateTax is a function that calculates the tax amount based on the given EBT (Earnings Before Tax) and tax rate.
// It takes two parameters: ebt and taxRate, both of type float64.
// ebt represents the Earnings Before Tax.
// taxRate represents the tax rate in percentage.
// The function returns the calculated tax amount as a float64.
func calculateTax(ebt, taxRate float64) float64 {
	return ebt * taxRate / 100
}

// calculateRatio calculates the ratio of earnings before tax (ebt) to profit.
// It takes two parameters: ebt and profit, both of type float64.
// It returns the calculated ratio as a float64.
// If profit is 0, it will return +Inf or -Inf depending on the sign of ebt.
func calculateRatio(ebt, profit float64) float64 {
	return ebt / profit
}

// getUserInput is a function that prompts the user to input revenue, expenses, and tax rate.
// It returns these values as float64 types. If there is an error during the scanning process,
// it prints an error message and continues to the next input.
func getUserInput() (revenue, expenses, taxRate float64) {
	// Prompt the user to enter the revenue
	fmt.Print("Revenue: ")
	_, err := fmt.Scan(&revenue)
	// If there is an error, print it and continue
	if err != nil {
		fmt.Println("Error reading revenue:", err)
	}

	// Prompt the user to enter the expenses
	fmt.Print("Expenses: ")
	_, err = fmt.Scan(&expenses)
	// If there is an error, print it and continue
	if err != nil {
		fmt.Println("Error reading expenses:", err)
	}

	// Prompt the user to enter the tax rate
	fmt.Print("Tax Rate: ")
	_, err = fmt.Scan(&taxRate)
	// If there is an error, print it and continue
	if err != nil {
		fmt.Println("Error reading tax rate:", err)
	}

	// Return the entered values
	return revenue, expenses, taxRate
}
