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
	mainHobbies = mainHobbies[1:3]
	fmt.Println("Main Hobbies Re-slice:", mainHobbies)
	fmt.Println("Main Hobbies Capacity:", cap(mainHobbies))
	fmt.Println("------------------------------------")
	courseGoals := []string{"Feel more confident in ability to write Go", "Be a more focused learner", "Be able to take on more complex projects"}
	for i, goal := range courseGoals {
		fmt.Println(i+1, goal)
	}
}
