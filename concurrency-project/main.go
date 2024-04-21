package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	c := make([]chan bool, len(taxRates))

	for i, taxRate := range taxRates {
		c[i] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go func() {
			err := priceJob.Process()
			if err != nil {
				fmt.Println("Error processing job:", err)
			}
		}()

		// err := priceJob.IOManager.WriteResult(priceJob)

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}
}
