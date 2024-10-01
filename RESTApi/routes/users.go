package routes

import (
	"net/http"

	"example.com/restapi/models"
	"example.com/restapi/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse data"})
		return 
	}	

	if user.Save() != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not create user."})	
		return	 
	}
	context.JSON(http.StatusCreated, gin.H{"message":"User created!","event":user})
}

func login(context *gin.Context) {
	var user models.User 

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse data"})
		return 
	}	

	if user.Validate() != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate"})
		return
	}

	token,err := utils.GenerateToken(user.Email,user.ID)

	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate"})
		return
	}

	context.JSON(http.StatusOK,gin.H{"message": "login successful", "token":token})
}