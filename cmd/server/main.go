package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/net/context"

	"github.com/al-kirpichenko/gophkeeper/cmd/server/config"
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

	// через этот канал сообщим основному потоку, что соединения закрыты
	idleConnsClosed := make(chan struct{})

	sigint := make(chan os.Signal, 1)
	// регистрируем перенаправление прерываний
	signal.Notify(sigint, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)

	// запускаем горутину обработки пойманных прерываний
	go func() {
		// читаем из канала прерываний
		<-sigint

		// получили сигнал os.Interrupt, запускаем процедуру graceful shutdown
		if err := srv.Shutdown(context.Background()); err != nil {
			// ошибки закрытия Listener
			log.Printf("HTTP server Shutdown: %v", err)
		}

		// сообщаем основному потоку,
		// что все сетевые соединения обработаны и закрыты
		close(idleConnsClosed)
	}()

	run(srv)

}

func run(srv *http.Server) {

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// ошибки запуска или остановки Listener
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}
