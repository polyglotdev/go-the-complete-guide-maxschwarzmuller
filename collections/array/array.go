package array

// Total takes in a slice of integers and returns the sum of all the integers in the slice.
func Total(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// TotalFloat64 takes in a slice of float64s and returns the sum of all the float64s in the slice.
func TotalFloat64(prices []float64) float64 {
	total := 0.0
	for _, price := range prices {
		total += price
	}
	return total
}
