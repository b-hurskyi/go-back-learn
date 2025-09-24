package routes

import (
	"github.com/b-hurskyi/go-back-learn/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEventById)
	authenticated.DELETE("/events/:id", deleteEventById)

	server.POST("/signup", signUp)
	server.POST("/login", login)
}
