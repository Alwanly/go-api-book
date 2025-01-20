package main

import (
	"context"
	"fmt"
	"go-codebase/pkg/config"
	"go-codebase/pkg/database"
	"go-codebase/pkg/logger"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {

	// load config
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		panic(err)
	}

	// Setup dependencies
	globalLogger := logger.NewLogger(cfg.ServiceName, cfg.LogLevel)
	l := logger.WithId(globalLogger, "server", "main")
	l.Info("Starting application")

	dbConfig := database.DBServiceOpts{
		Debug:  cfg.Debug,
		Logger: globalLogger,
	}

	database.SetPostgresUri(cfg.PostgresUri)

	db, err := database.NewPostgres(&dbConfig)
	if err != nil {
		l.Error("Cannot create database", zap.Error(err))
		panic(err)
	}

	// Create app
	app := Bootstrap(&AppDeps{
		Config: cfg,
		Logger: globalLogger,
		DB:     db,
	})

	// Register health check

	//--------------------- Bootstrap Application ---------------------

	ctx, cancel := context.WithCancel(context.Background())
	g, gCtx := errgroup.WithContext(ctx)

	// run http server
	g.Go(func() error {
		l.Info("Starting server...", zap.Int("port", cfg.Port))
		return app.Fiber.Listen(fmt.Sprintf(":%d", cfg.Port))
	})

	// graceful shutdown
	g.Go(func() error {

		<-gCtx.Done()
		l.Info("Gracefully shutting down...")

		l.Info("Server gracefully shutdown")
		if err := app.Fiber.Shutdown(); err != nil {
			l.Error("Cannot shutdown server", zap.Error(err))
			return err
		}

		l.Info("Closing database connection")
		if err := app.DB.Close(); err != nil {
			l.Error("Cannot close database connection", zap.Error(err))
			return err
		}

		return nil
	})

	// listen for interrupt signal
	go func() {
		c := make(chan os.Signal, 1)

		signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		l.Info("Listening for OS signal...")
		<-c

		// cancel context
		l.Info("Received OS signal, canceling context...")
		cancel()
	}()

	// wait for all goroutines to finish
	if err := g.Wait(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
