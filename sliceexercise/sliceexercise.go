package main

import "fmt"

func main() {
	hobbies := [3]string{"learning", "euphonium", "writing code"}
	for i, hobby := range hobbies {
		fmt.Println(i+1, hobby)
	}
	fmt.Println("------------------------------------")
	fmt.Println("First Hobby:", hobbies[0])
	fmt.Println("Second and Third Hobby:", hobbies[1:])

	fmt.Println("------------------------------------")
	mainHobbies := hobbies[:2]
	fmt.Println("Main Hobbies:", mainHobbies)
	fmt.Println("Main Hobbies Capacity:", cap(mainHobbies))
	fmt.Println("------------------------------------")
	otherHobbies := mainHobbies[1:]
	fmt.Println("Other Hobbies:", otherHobbies)
}
