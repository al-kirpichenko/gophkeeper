package api

import (
	"gorm.io/gorm"

	"github.com/al-kirpichenko/gophkeeper/cmd/config"
	"github.com/al-kirpichenko/gophkeeper/internal/database"
)

type Server struct {
	Config *config.Config
	DB     *gorm.DB
}

func NewServer(config *config.Config) *Server {

	return &Server{
		Config: config,
		DB:     database.InitDB(config.DataBaseURI),
	}
}
