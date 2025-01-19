package utils

import (
	"go-codebase/pkg/logger"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

const ContextName = "Binding"

type (
	Binder struct {
		l   *zap.Logger
		ctx *fiber.Ctx
		m   interface{}
	}

	BindingSource func(*Binder) error

	ModelBindingError struct {
		Code         int
		ResponseBody JSONResult
	}
)

func (e *ModelBindingError) Error() string {
	return "Failed to bind request body"
}

func BindFromBody() BindingSource {
	return func(b *Binder) error {
		if err := b.ctx.BodyParser(b.m); err != nil {
			b.l.Debug("Error when binding from body", zap.Error(err))
			return err
		}

		return nil
	}
}

func BindFromQuery() BindingSource {
	return func(b *Binder) error {
		if err := b.ctx.QueryParser(b.m); err != nil {
			b.l.Debug("Error when binding from query string", zap.Error(err))
			return err
		}

		return nil
	}
}

func BindFromParams() BindingSource {
	return func(b *Binder) error {
		if err := b.ctx.ParamsParser(b.m); err != nil {
			b.l.Debug("Error when binding from path params", zap.Error(err))
			return err
		}

		return nil
	}
}

func BindFromHeaders() BindingSource {
	return func(b *Binder) error {
		if err := b.ctx.ReqHeaderParser(b.m); err != nil {
			b.l.Debug("Error when binding from request headers", zap.Error(err))
			return err
		}

		return nil
	}
}

func BindModel(log *zap.Logger, c *fiber.Ctx, m interface{}, sources ...BindingSource) error {
	// create local logger
	l := logger.WithId(log, ContextName, "BindModel")

	// create binder instance
	binder := &Binder{
		l:   l,
		ctx: c,
		m:   m,
	}

	// process data binding
	for _, source := range sources {
		// execute binding
		if err := source(binder); err != nil {
			result := ResponseFailed(http.StatusBadRequest, StatusCodeBindingFailed, ErrorValidatePayload, nil)
			return &ModelBindingError{
				Code:         result.Code,
				ResponseBody: result,
			}
		}
	}
	return nil
}
