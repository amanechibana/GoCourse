package main

import (
	"net/http"
	"example.com/restapi/db"
	"example.com/restapi/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id",getEvent) // /event/1 /event/5
	server.POST("/events", createEvent)

	server.Run(":8080") // localhost: 8080
}

func getEvents(context *gin.Context) {
	events, err :=  models.GetAllEvents()
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch events."})
		return
	}
	
	context.JSON(http.StatusOK,events)
}

func getEvent(context *gin.Context){
	eventID,err := strconv.ParseInt(context.Param("id"),10,64)
	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not fetch event."})
		return
	}

	event,err := models.GetEventByID(eventID)

	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch event."})
		return		
	}

	context.JSON(http.StatusOK,event)	
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse data"})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()

	if err!= nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not create events."})	
		return	 
	}
	context.JSON(http.StatusCreated, gin.H{"message":"Event created!","event":event})
}