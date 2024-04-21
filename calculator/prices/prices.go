package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// PriceJobWithTax represents a job that processes prices with tax.
type PriceJobWithTax struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

// Process calculates the price after tax for each price in InputPrices.
// It creates a map where the keys are the original prices (as strings) and the values are the prices after tax.
// The results are not returned, but can be accessed through the TaxRate field of the PriceJobWithTax struct.
func (p *PriceJobWithTax) Process() {
	p.LoadData()

	result := make(map[string]string)
	for _, price := range p.InputPrices {
		taxIncludedPrices := price * (1 + p.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrices)
	}
	fmt.Println(result)
}

func (j *PriceJobWithTax) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))

	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)

		if err != nil {
			fmt.Println("Error parsing line", lineIndex, ":", err)
			file.Close()
			return
		}

		prices[lineIndex] = floatPrice
	}

	j.InputPrices = prices
}

// NewPriceJobWithTax creates a new PriceJobWithTax object.
// It takes a tax rate as an argument and initializes the InputPrices with default values.
// It also initializes an empty map for TaxIncludedPrices.
// The function returns a pointer to the newly created PriceJobWithTax object.
func NewPriceJobWithTax(taxRate float64) *PriceJobWithTax {
	return &PriceJobWithTax{
		InputPrices:       []float64{10, 20, 30},
		TaxRate:           taxRate,
		TaxIncludedPrices: make(map[string]float64),
	}
}
