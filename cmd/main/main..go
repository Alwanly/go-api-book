package main

import (
	"context"
	"fmt"

	"os"
	"os/signal"
	"syscall"

	"github.com/Alwanly/go-codebase/pkg/authentication"
	"github.com/Alwanly/go-codebase/pkg/config"
	"github.com/Alwanly/go-codebase/pkg/database"
	"github.com/Alwanly/go-codebase/pkg/logger"
	"github.com/Alwanly/go-codebase/pkg/middleware"
	"github.com/Alwanly/go-codebase/pkg/redis"
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
	globalLogger := logger.NewLogger(cfg.ServiceName, cfg.LogLevel,
		logger.WithPrettyPrint(),
	)
	l := logger.WithID(globalLogger, "server", "main")
	l.Info("Starting application")

	// Setup database
	dbConfig := database.DBServiceOpts{
		Debug:                      cfg.Debug,
		Logger:                     globalLogger,
		PostgresURI:                &cfg.PostgresURI,
		PostgresMaxOpenConnections: cfg.PostgresMaxOpenConnections,
		PostgresMaxIdleConnections: cfg.PostgresMaxIdleConnections,
	}

	db, err := database.NewPostgres(&dbConfig)
	if err != nil {
		l.Error("Cannot create database", zap.Error(err))
		panic(err)
	}

	// Setup redis
	redisConfig := redis.Opts{
		Logger:   globalLogger,
		RedisURI: &cfg.RedisURI,
	}
	redis, err := redis.NewRedis(&redisConfig)
	if err != nil {
		l.Error("Cannot create redis", zap.Error(err))
		panic(err)
	}

	// Setup middleware
	jwtConfig := middleware.SetJwtAuth(&authentication.JWTConfig{
		PrivateKey:     cfg.PrivateKey,
		PublicKey:      cfg.PublicKey,
		Audience:       cfg.JwtAudience,
		Issuer:         cfg.JwtIssuer,
		ExpirationTime: cfg.JwtExpirationTime,
		RefreshTime:    cfg.JwtRefreshTime,
	})
	basicAuthConfig := middleware.SetBasicAuth(&authentication.BasicAuthTConfig{
		Username: cfg.BasicAuthUsername,
		Password: cfg.BasicAuthPassword,
	})

	authMiddleware := middleware.NewAuthMiddleware(jwtConfig, basicAuthConfig)
	if authMiddleware == nil {
		l.Error("Cannot create auth middleware")
		panic("Cannot create auth middleware")
	}

	// Create app
	app := Bootstrap(&AppDeps{
		Config: &cfg,
		Logger: globalLogger,
		DB:     db,
		Redis:  redis,
		Auth:   authMiddleware,
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
