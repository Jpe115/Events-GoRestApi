package routes

import (
	"example.com/rest_api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/yours", middlewares.Authenticate, getYourEvents)
	server.GET("/events/:id", getEvent)

	server.POST("/events", middlewares.Authenticate, createEvent)
	server.PUT("/events/:id", middlewares.Authenticate, updateEvent)
	server.DELETE("/events/:id", middlewares.Authenticate, deleteEvent)

	server.POST("/events/:id/register", middlewares.Authenticate, registerForEvent)
	server.DELETE("/events/:id/register", middlewares.Authenticate, cancelRegistration)
	server.GET("/events/:id/registrations", middlewares.Authenticate, getRegisteredUsers)

	server.POST("/signup", signUp)
	server.POST("/login", login)
}
