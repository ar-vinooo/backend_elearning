package main

import (
	"fmt"
	"golang_api/controllers"
	"golang_api/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	result, _ := controllers.HashPassword("arvino")

	fmt.Println(result)

	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())

	router.POST("/api/login", controllers.Login)

	user := router.Group("/api").Use(middlewares.Auth())
	{
		user.GET("/user", controllers.GetUser)
	}

	guru := router.Group("/api/guru").Use(middlewares.Auth())
	{
		guru.GET("/kelas", controllers.GetKelas)

		guru.GET("/post", controllers.GetPost)
		guru.POST("/post", controllers.CreatePost)
		guru.DELETE("/post", controllers.DeletePost)
		guru.GET("/post/comment", controllers.GetCommentPost)
		guru.POST("/post/comment", controllers.CreateCommentPost)
	}

	router.Run()

}
