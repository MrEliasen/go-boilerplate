package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/placeholder/boiler/internal/database"
	"github.com/placeholder/boiler/pkg/logger"
	"github.com/placeholder/boiler/web"
)

func Run() {
	logger.Logger().Info("Initialising..")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	shutdownComplete := make(chan int, 1)

	logger.Logger().Info("Connecting to DB..")
	database.Connection()

	logger.Logger().Info("Setting up API")
	api := web.New(NewApp())
	go func() {
		api.Listen()
		os.Exit(0)
	}()

	logger.Logger().Info("Server ready!")
	<-shutdownSignal
	logger.Logger().Warn("Received shutdown signal")

	go func() {
		logger.Logger().Warn("Gracefully shutting down..")
		api.Shutdown(ctx)
		shutdownComplete <- 1
	}()

	<-shutdownComplete
}
