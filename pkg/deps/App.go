package deps

import (
	"github.com/Alwanly/go-codebase/config"
	"github.com/Alwanly/go-codebase/pkg/database"
	"github.com/Alwanly/go-codebase/pkg/middleware"
	"github.com/Alwanly/go-codebase/pkg/redis"
	"github.com/Alwanly/go-codebase/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type App struct {
	Config    *config.GlobalConfig
	Logger    *zap.Logger
	DB        *database.DBService
	Redis     *redis.Service
	Auth      *middleware.AuthMiddleware
	Validator validator.IValidatorService

	// APIs
	Fiber *fiber.App
}
