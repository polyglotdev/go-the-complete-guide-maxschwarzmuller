package prices

type PriceJobWithTax struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewPriceJobWithTax(taxRate float64, inputPrices []float64) *PriceJobWithTax {
	return &PriceJobWithTax{
		InputPrices:       inputPrices,
		TaxRate:           taxRate,
		TaxIncludedPrices: make(map[string]float64),
	}
}
