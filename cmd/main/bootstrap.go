package main

import (
	"encoding/json"
	"go-codebase/pkg/config"
	"go-codebase/pkg/database"
	"go-codebase/pkg/deps"
	"go-codebase/pkg/middleware"
	"go-codebase/pkg/redis"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

type (
	AppDeps struct {
		Config *config.GlobalConfig
		Logger *zap.Logger
		DB     database.IDBService
		Redis  redis.IRedisService
	}
)

var inst *deps.App

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

	// set app instance
	inst = &deps.App{
		Config: d.Config,
		Logger: d.Logger,
		DB:     d.DB,
		Redis:  d.Redis,
		Fiber:  e,
	}

	return inst
}
