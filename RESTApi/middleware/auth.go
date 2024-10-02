package middleware

import (
	"net/http"

	"example.com/restapi/models"
	"example.com/restapi/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == ""{
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token) 

	if err != nil{
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message":"Could not parse data"})
		return
	}

	context.Next()
}