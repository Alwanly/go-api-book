package deps

import (
	"go-codebase/pkg/config"
	"go-codebase/pkg/database"
	"go-codebase/pkg/middleware"
	"go-codebase/pkg/redis"
	"go-codebase/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type App struct {
	Config    *config.GlobalConfig
	Logger    *zap.Logger
	DB        *database.DBService
	Validator *validator.ValidatorService
	Redis     *redis.RedisService
	Auth      *middleware.AuthMiddleware

	// APIs
	Fiber *fiber.App
}
