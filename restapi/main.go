package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/rest-api/models"
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
	events := models.GetEvents()
	c.JSON(http.StatusOK, events)
}
