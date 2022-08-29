package controllers

import (
	"errors"
	database "golang_api/lib"
	"golang_api/structs"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("test")

type Claims struct {
	Email string
	jwt.RegisteredClaims
}

func Login(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	remember, _ := strconv.ParseBool(ctx.PostForm("remember"))

	var user structs.User
	var response gin.H

	result := database.DB().Table("users").Where("email = ?", email).First(&user).Error

	if result != nil {
		response = gin.H{"success": false, "message": "email or password wrong!"}
	}

	match := CheckPasswordHash(password, user.Password)

	if match {
		expirationTime := time.Now().AddDate(0, 0, 7)
		if !remember {
			expirationTime = time.Now().AddDate(0, 1, 0)
		}

		claims := &Claims{
			Email: user.Email,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, _ := token.SignedString(jwtKey)

		response = gin.H{"success": true, "message": "successfully loggedin", "token": tokenString}
	} else {
		response = gin.H{"success": false, "message": "email or password wrong!"}

	}

	ctx.JSON(http.StatusOK, response)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidateToken(signedToken string) (email string, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	email = claims.Email
	return
}
