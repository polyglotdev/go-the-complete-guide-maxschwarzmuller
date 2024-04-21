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

func transformNumbers(n []int, f func(int) int) []int {
	result := make([]int, len(n))
	for i, v := range n {
		result[i] = f(v)
	}
	return result
}
