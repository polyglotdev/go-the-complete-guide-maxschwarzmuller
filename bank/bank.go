package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Go Bank!")
	balance := ReadBalance()

	for {
		PrintMenu()
		choice := GetChoice()
		ProcessChoice(choice, &balance)
	}
}
