package main

import (
	"log"

	"example.com/structs/user"
)

func main() {
	elijah, err := user.New()
	if err != nil {
		log.Fatal(err)
	}
	elijah.PrintUser()

	ezra, err := user.New()
	if err != nil {
		log.Fatal(err)
	}
	ezra.PrintUser()

	admin := user.NewAdmin("admin@admin.com", "admin", "Admin", "Admin", "01/01/2000")

	admin.PrintUser()
}
