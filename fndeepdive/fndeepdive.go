package main

import "fmt"

type transformType func(int, int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	result1 := transformNumbers(numbers, 2, func(n int, multiplier int) int {
		return n * multiplier
	})

	result2 := transformNumbers(numbers, 3, func(n, multiplier int) int {
		return n * multiplier
	})

	result3 := transformNumbers(numbers, 4, func(n, multiplier int) int {
		return n * multiplier
	})

	fmt.Println("----------------------------------------")
	fmt.Printf("Double numbers: %v\n", result1)
	fmt.Printf("Triple numbers as func: %v\n", result2)
	fmt.Printf("Quadruple numbers: %v\n", result3)
	fmt.Println("----------------------------------------")
	fmt.Println("Numbers: ", numbers)
	fmt.Println("----------------------------------------")
	transformer := getTransformerFunction()
	result4 := transformNumbers(numbers, 5, transformer)
	fmt.Printf("5x numbers as func: %v\n", result4)

	fmt.Println("----------------------------------------")
	transformer2 := createTransformer(6)
	result5 := transformNumbers(numbers, 6, transformer2)
	fmt.Printf("6x numbers as func: %v\n", result5)
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
func transformNumbers(n []int, multiplier int, f transformType) []int {
	result := make([]int, len(n))
	for i, v := range n {
		result[i] = f(v, multiplier)
	}
	return result
}

func getTransformerFunction() transformType {
	return func(n int, multiplier int) int {
		return n * multiplier
	}
}

func createTransformer(multiplier int) transformType {
	return func(n int, multiplier int) int {
		return n * multiplier
	}
}
