package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	result1 := transformNumbers(numbers, double)
	result2 := transformNumbers(numbers, func(n int) int {
		return n * 2
	})
	result3 := transformNumbers(numbers, func(n int) int {
		return n * 3
	})
	fmt.Println("----------------------------------------")
	fmt.Printf("Double numbers: %v\n", result1)
	fmt.Printf("Double numbers as func: %v\n", result2)
	fmt.Printf("Triple numbers: %v\n", result3)
	fmt.Println("----------------------------------------")
	fmt.Println("Numbers: ", numbers)
}

func double(number int) int {
	return number * 2
}

// transformNumbers is a function that takes a slice of integers and a function as arguments.
// It applies the provided function to each element of the slice and returns a new slice with the results.
// The provided function should take an integer as an argument and return an integer.
//
// Parameters:
//
//	n: A slice of integers to be transformed.
//	f: A function that takes an integer as an argument and returns an integer. This function is applied to each element of 'n'.
//
// Returns:
//
//	A new slice of integers where each element is the result of applying 'f' to the corresponding element in 'n'.
func transformNumbers(n []int, f func(int) int) []int {
	result := make([]int, len(n))
	for i, v := range n {
		result[i] = f(v)
	}
	return result
}
