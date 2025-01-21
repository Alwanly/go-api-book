package middleware

import (
	"github.com/Alwanly/go-codebase/pkg/contract"
	"github.com/Alwanly/go-codebase/pkg/wrapper"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Recover(l *zap.Logger) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		l.Error("Unexpected error", zap.Error(err), zap.String("method", ctx.Method()), zap.String("url", ctx.Path()))
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(wrapper.JSONResult{
				Code:       fiber.StatusInternalServerError,
				StatusCode: contract.StatusCodeInternalServerError,
				Message:    "Internal server error",
			})
	}
}
