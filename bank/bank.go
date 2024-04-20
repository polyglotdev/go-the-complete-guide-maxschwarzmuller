package main

import "fmt"

func main() {
	welcome, _ := greeting()
	fmt.Println(welcome)
	menu := menu()
	fmt.Println(menu)
}

func greeting() (string, string) {
	welcome := "Welcome to the Go bank!"
	instructions := "Please select an option from the menu below:"

	return fmt.Sprintf("%s\n%s", welcome, instructions), welcome
}

func menu() string {
	menu := `
1. Check balance
2. Deposit funds
3. Withdraw funds
4. Exit
`
	return menu
}
