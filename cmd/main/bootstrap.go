package main

import (
	"encoding/json"

	"github.com/Alwanly/go-codebase/config"
	"github.com/Alwanly/go-codebase/pkg/database"
	"github.com/Alwanly/go-codebase/pkg/deps"
	"github.com/Alwanly/go-codebase/pkg/middleware"
	"github.com/Alwanly/go-codebase/pkg/redis"
	"github.com/Alwanly/go-codebase/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"go.uber.org/zap"

	_ "github.com/Alwanly/go-codebase/api"
	book_handler "github.com/Alwanly/go-codebase/internal/book/handler"
	user_handler "github.com/Alwanly/go-codebase/internal/user/handler"
)

type (
	AppDeps struct {
		Config *config.GlobalConfig
		Logger *zap.Logger
		DB     *database.DBService
		Redis  *redis.Service
		Auth   *middleware.AuthMiddleware
	}
)

var inst *deps.App

// @title Fiber Example API
// @version 1.0
// @description This is a sample Swagger example for Fiber with Basic Auth and JWT
// @host localhost:8080
// @BasePath /

// @securityDefinitions.basic BasicAuth
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization
func Bootstrap(d *AppDeps) *deps.App {
	// create http server
	e := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		ErrorHandler:          middleware.Recover(d.Logger),
	})

	// register middleware
	e.Use(cors.New())
	e.Use(recover.New())

	// create validator
	v, _ := validator.NewValidator()

	// add swagger docs if in development mode
	if d.Config.Environment == "development" {
		e.Static("/swagger.yaml", "./api/swagger.yaml")
		e.Get("/swagger/*", swagger.HandlerDefault)
	}

	// set app instance
	inst = &deps.App{
		Config:    d.Config,
		Logger:    d.Logger,
		DB:        d.DB,
		Redis:     d.Redis,
		Auth:      d.Auth,
		Fiber:     e,
		Validator: v,
	}
	database.MigrateIfNeed(inst.DB.Gorm)
	user_handler.NewHandler(inst)
	book_handler.NewHandler(inst)

	return inst
}
