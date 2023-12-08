package api

import (
	"log/slog"
	"os"

	"github.com/al-kirpichenko/gophkeeper/cmd/server/config"
	"github.com/al-kirpichenko/gophkeeper/internal/services/auth"
	"github.com/al-kirpichenko/gophkeeper/internal/services/keeper"
	"github.com/al-kirpichenko/gophkeeper/internal/storage"
)

type Server struct {
	Config        *config.Config
	Storage       *storage.Storage
	AuthService   *auth.Auth
	Log           *slog.Logger
	KeeperService *keeper.Keeper
}

func NewServer(config *config.Config) *Server {

	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
	store := storage.NewStorage(config.DataBaseURI)

	return &Server{
		Config:        config,
		Storage:       store,
		Log:           log,
		AuthService:   auth.NewAuth(log, store, config.TokenTTL),
		KeeperService: keeper.NewKeeper(log, store),
	}
}
