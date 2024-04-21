package prices

import "fmt"

type PriceJobWithTax struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (p *PriceJobWithTax) Process() {
	result := make(map[string]float64)
	for _, price := range p.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + p.TaxRate)
	}
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
