package controllers

import (
	database "golang_api/lib"
	"golang_api/structs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPost(ctx *gin.Context) {
	email, _ := ctx.Get("email")

	var list_post []structs.Post
	var response gin.H

	err := database.DB().Raw("SELECT gurus.id as guru_id, users.name as guru_name, gurus.nip as guru_nip, kelas.id as kelas_id,kelas.name as kelas_name ,posts.id as post_id, posts.type, discussions.id as discussion_id, discussions.content, posts.created_at as created_at FROM users LEFT JOIN gurus ON gurus.user_id = users.id LEFT JOIN posts ON posts.guru_id = gurus.id LEFT JOIN kelas ON kelas.id = posts.kelas_id LEFT JOIN discussions ON discussions.id = posts.type_id AND posts.type = 'discussions' WHERE users.email = ? AND discussions.id IS NOT NULL AND posts.id IS NOT NULL", email).Find(&list_post).Error

	if err != nil {
		response = gin.H{"success": false, "message": "kelas not found!", "data": nil}
	} else {
		response = gin.H{"success": true, "message": "successfully retrieve kelas!", "data": list_post}
	}

	ctx.JSON(http.StatusOK, response)
}

func CreatePost(ctx *gin.Context) {
	email, _ := ctx.Get("email")
	kelas_id := ctx.PostForm("kelas_id")
	content := ctx.PostForm("content")

	var guru structs.Guru
	var response gin.H
	var created_at = time.Now()

	err := database.DB().Raw("SELECT gurus.id as id, gurus.user_id as user_id, gurus.nip as nip, gurus.gender as gender FROM users LEFT JOIN gurus ON users.id = gurus.user_id WHERE users.email = ? LIMIT 1", email).First(&guru).Error

	if err != nil {
		response = gin.H{"success": false, "message": "guru not found!", "data": nil}
	} else {
		discussion := structs.Discussion{Content: content}
		database.DB().Create(&discussion)

		database.DB().Exec("INSERT INTO posts (guru_id, kelas_id, type_id, type, created_at) VALUES (?, ?, ?, ?, ?)", guru.ID, kelas_id, discussion.ID, "discussions", created_at)

		response = gin.H{"success": true, "message": "successfully create discussion!"}
	}

	ctx.JSON(http.StatusOK, response)

}

func DeletePost(ctx *gin.Context) {
	post_id := ctx.Query("post_id")
	discussion_id := ctx.Query("discussion_id")

	database.DB().Exec("DELETE FROM posts WHERE id = ?", post_id)
	database.DB().Exec("DELETE FROM discussions WHERE id = ?", discussion_id)

	response := gin.H{"success": true, "message": "successfully delete post!"}
	ctx.JSON(http.StatusOK, response)
}

func GetCommentPost(ctx *gin.Context) {
	post_id := ctx.Query("post_id")

	var list_comment_post []structs.CommentPost
	var response gin.H

	err := database.DB().Raw("SELECT post_comments.id as id, users.id as user_id, users.akses as user_type, users.name as user_name, IF(users.akses ='guru', gurus.nip, siswas.nisn) as user_number, posts.id as post_id, post_comments.comment as comment, post_comments.created_at as created_at FROM posts LEFT JOIN post_comments ON post_comments.post_id = posts.id LEFT JOIN users ON users.id = post_comments.user_id LEFT JOIN gurus ON gurus.user_id = users.id AND users.akses = 'guru' LEFT JOIN siswas ON siswas.user_id = users.id AND users.akses = 'siswa' WHERE posts.id = ? AND posts.id IS NOT NULL AND post_comments.id IS NOT NULL", post_id).First(&list_comment_post).Error

	if err != nil {
		response = gin.H{"success": false, "message": "comment post not found!", "data": nil}
	} else {
		response = gin.H{"success": true, "message": "successfully retrieve comment post!", "data": list_comment_post}
	}

	ctx.JSON(http.StatusOK, response)
}

func CreateCommentPost(ctx *gin.Context) {
	email, _ := ctx.Get("email")
	post_id := ctx.PostForm("post_id")
	comment := ctx.PostForm("comment")

	var guru structs.Guru
	var response gin.H
	var created_at = time.Now()

	err := database.DB().Raw("SELECT gurus.id as id, gurus.user_id as user_id, gurus.nip as nip, gurus.gender as gender FROM users LEFT JOIN gurus ON users.id = gurus.user_id WHERE users.email = ? LIMIT 1", email).First(&guru).Error

	if err != nil {
		response = gin.H{"success": false, "message": "guru not found!", "data": nil}
	} else {

		database.DB().Exec("INSERT INTO post_comments (post_id, user_id, comment, created_at) VALUES (?,?,?,?)", post_id, guru.ID, comment, created_at)

		response = gin.H{"success": true, "message": "successfully create comment!"}
	}

	ctx.JSON(http.StatusOK, response)
}
