package handler

import (
	"github.com/Alwanly/go-codebase/internal/user/repository"
	"github.com/Alwanly/go-codebase/internal/user/schema"
	"github.com/Alwanly/go-codebase/internal/user/usecase"
	"github.com/Alwanly/go-codebase/pkg/deps"
	"github.com/Alwanly/go-codebase/pkg/logger"
	"github.com/Alwanly/go-codebase/pkg/utils"
	"github.com/Alwanly/go-codebase/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

const ContextName = "Internal.User.Handler"

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

	e := d.Fiber.Group("/auth/v1")
	e.Post("/login", d.Auth.BasicAuth(), handler.Login)
	e.Post("/register", d.Auth.BasicAuth(), handler.Register)
	e.Get("/profile", d.Auth.JwtAuth(), handler.Profile)
	return handler
}

// Login
func (h *Handler) Login(c *fiber.Ctx) error {
	l := logger.WithID(h.Logger, ContextName, "Login")

	// bind model
	model := &schema.AuthLoginRequest{}
	if err := utils.BindModel(l, c, model, utils.BindFromBody()); err != nil {
		perr := err.(*utils.ModelBindingError)
		return c.Status(perr.Code).JSON(perr.ResponseBody)
	}

	// validate model
	if err := utils.ValidateModel(l, h.Validator, model); err != nil {
		perr := err.(*utils.ModelValidationError)
		return c.Status(perr.Code).JSON(perr.ResponseBody)
	}

	// process request
	response := h.UseCase.Auth(c.UserContext(), model)
	return c.Status(response.Code).JSON(response)
}

// Register
func (h *Handler) Register(c *fiber.Ctx) error {
	l := logger.WithID(h.Logger, ContextName, "Register")

	// bind model
	model := &schema.AuthRegisterRequest{}
	if err := utils.BindModel(l, c, model, utils.BindFromBody()); err != nil {
		perr := err.(*utils.ModelBindingError)
		return c.Status(perr.Code).JSON(perr.ResponseBody)
	}

	// validate model
	if err := utils.ValidateModel(l, h.Validator, model); err != nil {
		perr := err.(*utils.ModelValidationError)
		return c.Status(perr.Code).JSON(perr.ResponseBody)
	}

	// process request
	response := h.UseCase.Register(c.UserContext(), model)
	return c.Status(response.Code).JSON(response)
}

// Register
func (h *Handler) Profile(c *fiber.Ctx) error {
	l := logger.WithID(h.Logger, ContextName, "Register")

	// bind model
	model := &schema.ProfileRequest{}
	if err := utils.BindModel(l, c, model); err != nil {
		perr := err.(*utils.ModelBindingError)
		return c.Status(perr.Code).JSON(perr.ResponseBody)
	}

	// validate model
	if err := utils.ValidateModel(l, h.Validator, model); err != nil {
		perr := err.(*utils.ModelValidationError)
		return c.Status(perr.Code).JSON(perr.ResponseBody)
	}

	// process request
	response := h.UseCase.Profile(c.UserContext(), model)
	return c.Status(response.Code).JSON(response)
}
