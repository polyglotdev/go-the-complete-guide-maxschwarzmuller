package main

import "fmt"

func main() {
	result := factorial(5)
	fmt.Printf("result: %d\n", result)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}
