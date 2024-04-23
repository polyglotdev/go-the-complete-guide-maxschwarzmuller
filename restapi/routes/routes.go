package routes

import (
	"github.com/gin-gonic/gin"

	"example.com/rest-api/middlewares"
)

// Register defines the routes for the REST API.
// it takes a pointer to a gin.Engine instance as a parameter.
func Register(server *gin.Engine) {
	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent)
	server.POST("/events", middlewares.CheckAuth, CreateEvent)
	server.PUT("/events/:id", UpdateEvent)
	server.DELETE("/events/:id", DeleteEvent)
	server.POST("/signup", Signup)
	server.POST("/login", Login)
}
