package controllers

import (
	database "golang_api/lib"
	"golang_api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetKelas(ctx *gin.Context) {
	email, _ := ctx.Get("email")

	var list_kelas []structs.Kelas
	var response gin.H

	err := database.DB().Raw("SELECT kelas.id, kelas.name FROM users LEFT JOIN gurus ON gurus.user_id = users.id LEFT JOIN kelas ON kelas.guru_id = gurus.id WHERE users.email = ? AND kelas.deleted_at IS NULL", email).Find(&list_kelas).Error

	if err != nil {
		response = gin.H{"success": false, "message": "kelas not found!", "data": nil}
	} else {
		response = gin.H{"success": true, "message": "successfully retrieve kelas!", "data": list_kelas}
	}

	ctx.JSON(http.StatusOK, response)

}
