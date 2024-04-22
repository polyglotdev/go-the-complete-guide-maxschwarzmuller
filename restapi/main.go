package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"example.com/rest-api/db"
	"example.com/rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.Register(server)

	err := server.Run(":8080")
	if err != nil {
		errMsg := fmt.Errorf("error starting server: %v", err)
		fmt.Println(errMsg)
	}
}
