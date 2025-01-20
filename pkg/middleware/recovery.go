package middleware

import (
	"go-codebase/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Recover(l *zap.Logger) fiber.ErrorHandler {

	return func(ctx *fiber.Ctx, err error) error {
		l.Error("Unexpected error", zap.Error(err), zap.String("method", ctx.Method()), zap.String("url", ctx.Path()))
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(utils.JSONResult{
				Code:       fiber.StatusInternalServerError,
				StatusCode: utils.StatusCodeInternalServerError,
				Message:    "Internal server error",
			})
	}
}
