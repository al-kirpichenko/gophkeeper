package main

import (
	"net/http"

	"github.com/al-kirpichenko/gophkeeper/cmd/config"
	"github.com/al-kirpichenko/gophkeeper/internal/api"
	"github.com/al-kirpichenko/gophkeeper/internal/router"
)

func main() {

	cfg := config.NewConfig()

	server := api.NewServer(cfg)

	srv := &http.Server{
		Addr:    cfg.Host,
		Handler: router.Router(server),
	}

}
