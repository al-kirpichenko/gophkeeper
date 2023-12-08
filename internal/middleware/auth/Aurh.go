package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/al-kirpichenko/gophkeeper/internal/utils/jwt"
)

// Auth - проверка аутентификации
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("Authorization")
		log.Println(token)

		if token == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userID, err := jwt.GetUserIDFromToken(token)

		if err != nil || userID == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("userID", userID)

		ctx.Next()
	}
}
