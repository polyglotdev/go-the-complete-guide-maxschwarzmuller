package main

import (
	"log"

	"example.com/structs/user"
)

func main() {
	elijah, err := user.CreateUser()
	if err != nil {
		log.Fatal(err)
	}
	elijah.PrintUser()

	ezra, err := user.CreateUser()
	if err != nil {
		log.Fatal(err)
	}
	ezra.PrintUser()
}
