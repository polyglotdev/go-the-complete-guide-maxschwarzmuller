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

	authenticated := server.Group("/")
	authenticated.Use(middlewares.CheckAuth)
	authenticated.POST("/events", CreateEvent)
	authenticated.PUT("/events/:id", UpdateEvent)
	authenticated.DELETE("/events/:id", DeleteEvent)
	authenticated.POST("/events/:id/register", RegisterForEvent)
	authenticated.DELETE("/events/:id/register")

	server.POST("/signup", Signup)
	server.POST("/login", Login)
}
