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

	myArray4 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("myArray4 before slice:", myArray4)
	myArray4Slice := myArray4[:]
	// append the number six to the end of myArray4Slice
	myArray4Slice = append(myArray4Slice, 6)
	fmt.Println("myArray4 after slice:", myArray4)
	fmt.Println("myArray4Slice after slice:", myArray4Slice)

	fmt.Println("------------------------------------")
	featuredPrices := []float64{100.25, 200.98, 300.67}
	tester := featuredPrices[0:2]
	fmt.Println("featuredPrices before append:", featuredPrices)
	fmt.Println("tester before append:", tester)
	tester[0] = 455.23
	fmt.Println("featuredPrices after append:", featuredPrices)
	fmt.Println("tester after append:", tester)

	fmt.Println("------------------------------------")
	checkArray := [5]int{1, 2, 3, 4, 5}
	checkArray2 := checkArray
	fmt.Println("checkArray: ", checkArray)
	fmt.Println("checkArray2: ", checkArray2)
	checkArray2[0] = 100
	fmt.Println("checkArray: ", checkArray)
	fmt.Println("checkArray2: ", checkArray2)

	// Convert the array into a slice
	checkArraySlice := checkArray[:]

	fmt.Println("------------------------------------")
	fmt.Printf("checkArray Header:\nLength: %d\nCapacity: %d\nMemory Address: %v\n", len(checkArray), cap(checkArray), &checkArray)

	fmt.Println("------------------------------------")
	fmt.Printf("checkArray2 Header:\nLength: %d\nCapacity: %d\nMemory Address: %v\n", len(checkArray2), cap(checkArray2), &checkArray2)

	fmt.Println("------------------------------------")
	fmt.Printf("checkArraySlice Header:\nLength: %d\nCapacity: %d\nMemory Address: %v\n", len(checkArraySlice), cap(checkArraySlice), &checkArraySlice)

	fmt.Println("------------------------------------")
	fmt.Printf("checkArray Header:\nLength: %d\nCapacity: %d\nMemory Address: %v\n", len(checkArray), cap(checkArray), &checkArray)

	fmt.Println("------------------------------------")
	fmt.Printf("checkArray2 Header:\nLength: %d\nCapacity: %d\nMemory Address: %v\n", len(checkArray2), cap(checkArray2), &checkArray2)

	fmt.Println("------------------------------------")
	ticketPrices := []float64{100.25, 200.98, 300.67}
	fmt.Println(ticketPrices[0:1])
	updatedTicketPrices := append(ticketPrices, 400.50)
	fmt.Println("Ticket prices: ", ticketPrices)
	fmt.Println("Updated prices: ", updatedTicketPrices)
	updatedTicketPrices[0] = 43600.75
	fmt.Println("------------------------------------")
	fmt.Println("Ticket prices: ", ticketPrices)
	fmt.Println("Updated prices: ", updatedTicketPrices)

	fmt.Println("------------------------------------")
	dArray := [5]int{1, 2, 3, 4, 5}
	newDArray := dArray
	fmt.Println("dArray before change:", dArray)
	fmt.Println("newDArray before change:", newDArray)
	newDArray[0] = 100
	fmt.Println("dArray after change:", dArray)
	fmt.Println("newDArray after change:", newDArray)
	fmt.Printf("dArray memory address: %p\n", &dArray)
	fmt.Printf("newDArray memory address: %p\n", &newDArray)
}
