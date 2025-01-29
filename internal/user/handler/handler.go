package handler

import (
	"github.com/Alwanly/go-codebase/internal/user/repository"
	"github.com/Alwanly/go-codebase/internal/user/schema"
	"github.com/Alwanly/go-codebase/internal/user/usecase"
	"github.com/Alwanly/go-codebase/pkg/binding"
	"github.com/Alwanly/go-codebase/pkg/deps"
	"github.com/Alwanly/go-codebase/pkg/logger"
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

// @Summary User Login
// @Description Authenticate a user and return a token
// @ID user-login
// @Accept json
// @Produce json
// @Param login body schema.AuthLoginRequest true "Login request"
// @Security BasicAuth
// @Success 200 {object} schema.AuthLoginResponse
// @Router /auth/v1/login [post]
func (h *Handler) Login(c *fiber.Ctx) error {
	l := logger.WithID(h.Logger, ContextName, "Login")

	// bind model
	model := &schema.AuthLoginRequest{}
	if err := binding.BindModel(l, c, model, binding.BindFromBody()); err != nil {
		perr := err.(*binding.ModelBindingError)
		return c.Status(perr.Code).JSON(perr.ResponseBody)
	}

	// validate model
	if err := validator.ValidateModel(l, h.Validator, model); err != nil {
		perr := err.(*validator.ModelValidationError)
		return c.Status(perr.Code).JSON(perr.ResponseBody)
	}

	// process request
	response := h.UseCase.Auth(c.UserContext(), model)
	return c.Status(response.Code).JSON(response)
}

// @Summary User Registration
// @Description Register a new user
// @ID user-register
// @Accept json
// @Produce json
// @Param register body schema.AuthRegisterRequest true "Register request"
// @Security BasicAuth
// @Success 201 {object} schema.AuthRegisterResponse
// @Router /auth/v1/register [post]
func (h *Handler) Register(c *fiber.Ctx) error {
	l := logger.WithID(h.Logger, ContextName, "Register")

	// bind model
	model := &schema.AuthRegisterRequest{}
	if err := binding.BindModel(l, c, model, binding.BindFromBody()); err != nil {
		perr := err.(*binding.ModelBindingError)
		return c.Status(perr.Code).JSON(perr.ResponseBody)
	}

	// validate model
	if err := validator.ValidateModel(l, h.Validator, model); err != nil {
		perr := err.(*validator.ModelValidationError)
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
	if err := binding.BindModel(l, c, model); err != nil {
		perr := err.(*binding.ModelBindingError)
		return c.Status(perr.Code).JSON(perr.ResponseBody)
	}

	// validate model
	if err := validator.ValidateModel(l, h.Validator, model); err != nil {
		perr := err.(*validator.ModelValidationError)
		return c.Status(perr.Code).JSON(perr.ResponseBody)
	}

	// process request
	response := h.UseCase.Profile(c.UserContext(), model)
	return c.Status(response.Code).JSON(response)
}
