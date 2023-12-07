package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/al-kirpichenko/gophkeeper/internal/models"
)

// CreateSecret - создание секрета
func (s *Server) CreateSecret(ctx *gin.Context) {

	var secret models.Secret

	err := json.NewDecoder(ctx.Request.Body).Decode(&secret)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

}
