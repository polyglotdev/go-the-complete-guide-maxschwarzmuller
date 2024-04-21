package main

import (
	"fmt"
)

func main() {
	result := factorial(5)
	fmt.Printf("result: %d\n", result)

	result2 := factorial(10)
	fmt.Printf("result2: %d\n", result2)

	result3 := factorial(15)
	fmt.Printf("result3: %d\n", result3)

	fmt.Println("-------------------------------")
	sum := sumup(1, 2, 3, 4, 5)
	fmt.Printf("Sum: %d\n", sum)

	fmt.Println("-------------------------------")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum2 := sumup(nums[0], nums...)
	fmt.Printf("Sum2: %d\n", sum2)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func sumup(_ int, numbers ...int) int {
	sum := 0

	for _, v := range numbers {
		sum += v
	}
	return sum
}
