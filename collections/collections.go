package main

import (
	"fmt"

	"example.com/collections/array"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	total := array.Total(numbers)
	fmt.Printf("The total is: %d\n", total)

	prices := []float64{10.99, 5.99, 3.99, 7.99}
	totalPrice := array.TotalFloat64(prices[:])
	fmt.Printf("The total price is: $%.2f\n", totalPrice)

	average := array.Average(numbers)
	fmt.Printf("The average is: %.2f\n", average)

	averagePrice := array.AverageFloat64(prices[:])
	fmt.Printf("The average price is: $%.2f\n", averagePrice)

	cpAveragePrice := averagePrice
	fmt.Println("------------------------------------")
	fmt.Printf("The average price is: $%.2f\nThis is from cpAveragePrice\n", cpAveragePrice)

	fmt.Println("------------------------------------")
	firstTwo := array.FirstTwo(numbers)
	fmt.Printf("The first two numbers in numbers slice are: %v\n", firstTwo)

	myArray := [5]int{1, 2, 3, 4, 5}
	threeFourFive := myArray[2:]
	fmt.Printf("The last three elements in myArray are: %v\n", threeFourFive)

	// add the number six to the end of threeFourFive
	threeFourFive = append(threeFourFive, 6)
	fmt.Printf("The last three elements in myArray are: %v\n", threeFourFive)

	slice1 := []int{1, 2, 3}
	slice2 := slice1

	fmt.Println("------------------------------------")
	fmt.Printf("slice1 memory address: %p\n", &slice1)
	fmt.Printf("slice2 memory address: %p\n", &slice2)

	fmt.Println("------------------------------------")
	array1 := [3]int{1, 2, 3}
	array2 := array1

	fmt.Printf("array1 memory address: %p\n", &array1)
	fmt.Printf("array2 memory address: %p\n", &array2)
	fmt.Println("------------------------------------")

	// slice1 and slice2 point to the same memory address
	// so when we change slice1, slice2 will also change
	slice1[0] = 4
	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)
}
