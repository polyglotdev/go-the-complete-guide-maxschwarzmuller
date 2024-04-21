package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	c := make([]chan bool, len(taxRates))
	ec := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		c[i] = make(chan bool)
		ec[i] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(c[i], ec[i])
	}

	for i := range taxRates {
		select {
		case <-c[i]:
			fmt.Println("Job completed successfully")
		case err := <-ec[i]:
			fmt.Println("Job failed with error:", err)
		}
	}
}
