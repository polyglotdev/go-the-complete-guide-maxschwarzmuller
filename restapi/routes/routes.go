package routes

import "github.com/gin-gonic/gin"

// Register defines the routes for the REST API.
// it takes a pointer to a gin.Engine instance as a parameter.
func Register(server *gin.Engine) {
	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent)
	server.POST("/events", CreateEvent)
	server.PUT("/events/:id", UpdateEvent)
}
