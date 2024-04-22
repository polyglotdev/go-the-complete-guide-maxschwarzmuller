package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)

	err := server.Run(":8080")
	if err != nil {
		errMsg := fmt.Errorf("error starting server: %v", err)
		fmt.Println(errMsg)
	}
}

func getEvents(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GET events",
	})
}
