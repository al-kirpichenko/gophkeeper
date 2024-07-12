package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReadSecret - получить секрет по имени
func (s *Server) ReadSecret(ctx *gin.Context) {

	var (
		req struct {
			Title string `json:"title"`
		}
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	userID, _ := ctx.Get("userID")

	secret, err := s.KeeperService.ReadSecret(req.Title, userID.(uint))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	response, err := json.Marshal(secret)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.Data(http.StatusOK, "application/json", response)

}
