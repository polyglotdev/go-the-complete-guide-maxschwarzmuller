package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	result := doubleNumbers(numbers)
	fmt.Println(result)
}

func doubleNumbers(n []int) []int {
	result := make([]int, len(n))
	for i, v := range n {
		result[i] = v * 2
	}
	return result
}
