package handler

import (
	"go-codebase/internal/user/repository"
	"go-codebase/internal/user/usecase"
	"go-codebase/pkg/deps"
	"go-codebase/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type (
	Handler struct {
		Logger    *zap.Logger
		Validator validator.IValidatorService
		UseCase   usecase.IUseCase
	}
)

func NewHandler(d *deps.App) *Handler {

	repository := repository.NewRepository(repository.Repository{
		DB:    d.DB,
		Redis: d.Redis,
	})
	usecase := usecase.NewUseCase(usecase.UseCase{
		Config:     d.Config,
		Logger:     d.Logger,
		Jwt:        d.Auth.Jwt,
		Repository: repository,
	})
	handler := &Handler{
		Logger:    d.Logger,
		Validator: d.Validator,
		UseCase:   usecase,
	}

	e := d.Fiber.Group("/user")
	e.Get("/example", handler.ExampleMethod())
	return handler
}

// Define methods for the handler
func (h *Handler) ExampleMethod() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement the method

		return nil
	}
}
