package controllers

import (
	database "golang_api/lib"
	"golang_api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	email, _ := ctx.Get("email")

	var user structs.User
	var response gin.H

	err := database.DB().Table("users").Where("email = ?", email).First(&user).Error
	if err != nil {
		response = gin.H{"success": false, "message": "user not found!", "data": nil}
	} else {
		response = gin.H{"success": true, "message": "successfully retrieve user!", "data": &user}
	}

	ctx.JSON(http.StatusOK, response)
}
