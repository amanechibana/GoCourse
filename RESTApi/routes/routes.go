package routes

import (
	"example.com/restapi/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine){
	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id",getEvent) // /event/1 /event/5

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id",deleteEvent)
	authenticated.POST("/events/:id/register",registerForEvent)
	authenticated.DELETE("/events/:id/register",cancelRegister)

	server.POST("/signup",signup)
	server.POST("/login",login)
}

