package main

import (
	"encoding/json"

	"github.com/Alwanly/go-codebase/pkg/config"
	"github.com/Alwanly/go-codebase/pkg/database"
	"github.com/Alwanly/go-codebase/pkg/deps"
	"github.com/Alwanly/go-codebase/pkg/middleware"
	"github.com/Alwanly/go-codebase/pkg/redis"
	"github.com/Alwanly/go-codebase/pkg/utils"
	"github.com/Alwanly/go-codebase/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"

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

func Bootstrap(d *AppDeps) *deps.App {
	// create http server
	e := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		ErrorHandler:          utils.ResponseRecover(d.Logger),
	})

	// register middleware
	e.Use(cors.New())
	e.Use(recover.New())

	// create validator
	v, _ := validator.NewValidator()

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

	user_handler.NewHandler(inst)
	return inst
}
