package prices

import (
	"encoding/json"
	"fmt"

	"example.com/calculator/conversion"
	"example.com/calculator/filemanager"
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

	jsonResult, err := json.Marshal(result)
	if err != nil {
		fmt.Println("error marshaling JSON: " + err.Error())
		return
	}

	err = filemanager.WriteJSON(fmt.Sprintf("result_%0f.json", p.TaxRate*100), jsonResult)
	if err != nil {
		fmt.Println("error writing JSON file: " + err.Error())
		return
	}
}

func (j *PriceJobWithTax) LoadData() {
	lines, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println("error reading lines: " + err.Error())
		return
	}
	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println("error converting strings to floats: " + err.Error())
		return
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
