package inits

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	envs "github.com/bagasunix/ginclean/pkg/env"
	"github.com/bagasunix/ginclean/server/domains/data/repositories"
)

var (
	httpAddr = flag.String("http.addr", ":3000", "HTTP listen address")
)

func Run() {
	flag.Parse()
	ctx := context.Background()

	logger := InitLogger()
	configs, _ := envs.LoadEnv()
	db := InitDb(ctx, logger, configs)
	Migrate(logger, db)
	repo := repositories.New(logger, db)
	svc := InitService(logger, repo)
	eps := InitEndpoints(logger, svc)

	httpHandler := InitHttpHandler(logger, eps)

	srv := &http.Server{
		Addr:    *httpAddr,
		Handler: httpHandler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	// Wait for kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}
