package main

import (
	"example.com/restapi/db"
	"github.com/gin-gonic/gin"
	"example.com/restapi/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()  //returns a pointer to server

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost: 8080
}

