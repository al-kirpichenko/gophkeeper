package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/al-kirpichenko/gophkeeper/internal/models"
)

// Login авторизация пользователя
func (s *Server) Login(ctx *gin.Context) {

	var auth *models.Auth
	if err := ctx.ShouldBindJSON(&auth); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	token, err := s.AuthService.Login(auth)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.Writer.Header().Set("Authorization", token)
	ctx.JSON(http.StatusCreated, gin.H{"message": "the user has logged in"})

}
