package router

import (
	"github.com/al-kirpichenko/gophkeeper/internal/api"

	"github.com/gin-gonic/gin"
)

func Router(server *api.Server) *gin.Engine {

	r := gin.Default()

	r.POST("/api/user/register", server.Register)
	r.POST("/api/user/login", server.Login)

	return r
}
