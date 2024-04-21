package main

import "fmt"

func main() {
	result := factorial(5)
	fmt.Printf("result: %d\n", result)

	result2 := factorial(10)
	fmt.Printf("result2: %d\n", result2)

	result3 := factorial(15)
	fmt.Printf("result3: %d\n", result3)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}
