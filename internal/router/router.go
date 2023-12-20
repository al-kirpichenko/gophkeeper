package router

import (
	"github.com/al-kirpichenko/gophkeeper/internal/api"
	middleware "github.com/al-kirpichenko/gophkeeper/internal/middleware/auth"

	"github.com/gin-gonic/gin"
)

func Router(server *api.Server) *gin.Engine {

	r := gin.Default()

	auth := r.Group("/")

	{
		auth.Use(middleware.Auth())
		auth.POST("/api/secret/create", server.CreateSecret)
		auth.POST("/api/secret/read", server.ReadSecret)
	}
	r.POST("/api/user/register", server.Register)
	r.POST("/api/user/login", server.Login)

	return r
}
