package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/al-kirpichenko/gophkeeper/internal/models"
)

// Register - хэндлер регистрации нового пользователя
func (s *Server) Register(ctx *gin.Context) {

	var auth *models.Auth
	if err := ctx.ShouldBindJSON(&auth); err != nil {

		s.Log.Info("Dont read json")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	err := s.AuthService.CreateUser(auth)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "the new user has been successfully registered!"})
}
