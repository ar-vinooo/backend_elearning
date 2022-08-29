package middlewares

import (
	"golang_api/controllers"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(401, gin.H{"message": "unauthoried"})
			ctx.Abort()
			return
		}
		email, err := controllers.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(401, gin.H{"message": "unauthoried"})
			ctx.Abort()
			return
		}
		ctx.Set("email", email)
		ctx.Next()
	}
}
