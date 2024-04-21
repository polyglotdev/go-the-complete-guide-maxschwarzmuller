package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	for i, value := range nums {
		nums[i] = value * 2
	}
	fmt.Println(nums)
	fmt.Println("------------------------------")
	otherNums := []int{2, 4, 6, 8, 10}
	fmt.Println(otherNums)
	for index := 0; index < len(otherNums); index++ {
		otherNums[index] = otherNums[index] * 10
	}
	fmt.Println(otherNums)
}
