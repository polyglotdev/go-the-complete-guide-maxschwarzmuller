package prices

import "fmt"

type PriceJobWithTax struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (p *PriceJobWithTax) Calculate() {
	for _, price := range p.InputPrices {
		p.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = price * (1 + p.TaxRate)
	}
}

func NewPriceJobWithTax(taxRate float64) *PriceJobWithTax {
	return &PriceJobWithTax{
		InputPrices:       []float64{10, 20, 30},
		TaxRate:           taxRate,
		TaxIncludedPrices: make(map[string]float64),
	}
}
