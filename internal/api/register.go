package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/al-kirpichenko/gophkeeper/internal/database"
	"github.com/al-kirpichenko/gophkeeper/internal/models"
	"github.com/al-kirpichenko/gophkeeper/internal/services"
)

// Register - регистрация нового пользователя
func (s *Server) Register(ctx *gin.Context) {

	var auth *models.Auth
	if err := ctx.ShouldBindJSON(&auth); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	hashedPassword, err := services.HashPassword(auth.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	newUser := models.User{
		Login:    auth.Login,
		Password: hashedPassword,
	}

	err = database.CreateUser(s.DB, &newUser)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	token, err := services.GenerateToken(newUser.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token, "message": "the user is registered"})
}
