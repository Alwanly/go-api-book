package deps

import (
	"go-codebase/pkg/config"
	"go-codebase/pkg/database"
	"go-codebase/pkg/redis"
	"go-codebase/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type App struct {
	Config    *config.GlobalConfig
	Logger    *zap.Logger
	DB        database.IDBService
	Validator validator.IValidatorService
	Redis     redis.IRedisService

	// APIs
	Fiber *fiber.App
}
