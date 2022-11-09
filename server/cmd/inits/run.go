package inits

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	// ************ Database ************
	configs, _ := envs.LoadEnv()
	db := InitDb(ctx, logger, configs)
	Migrate(logger, db)

	// ************ Service ************
	repo := repositories.New(logger, db)
	svc := InitService(logger, configs, repo)
	eps := InitEndpoints(logger, svc)
	httpHandler := InitHttpHandler(&logger, eps)

	// ************ Transport ************
	errs := make(chan error)
	go initCancel(errs)
	go initHttp(httpHandler, errs)

	logger.Sugar().Error("exit", <-errs)
}

func initCancel(errs chan error) {
	// Wait for kill signal of channel
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// This blocks until a signal is passed into the quit channel
	errs <- fmt.Errorf("%s", <-quit)
}

func initHttp(httpHandler http.Handler, errs chan error) {
	srv := &http.Server{
		Addr:    *httpAddr,
		Handler: httpHandler,
	}
	errs <- srv.ListenAndServe()
}
