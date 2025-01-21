package utils

import (
	"math"
	"net/http"

	"github.com/Alwanly/go-codebase/pkg/common"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type JSONResult struct {
	Code       int               `json:"-"`
	StatusCode common.StatusCode `json:"statusCode"`
	Message    string            `json:"message"`
	Meta       *PaginationMeta   `json:"meta,omitempty"`
	Data       interface{}       `json:"data"`
}

type PaginationMeta struct {
	Page            int         `json:"page"`
	TotalData       int         `json:"totalData"`
	TotalPage       int         `json:"totalPage"`
	TotalDataOnPage int         `json:"totalDataOnPage"`
	MetaData        interface{} `json:"metadata,omitempty"`
}

func ResponseSuccess(code int, data interface{}) JSONResult {
	return JSONResult{
		Code:       code,
		StatusCode: common.StatusCodeSuccess,
		Message:    "Success",
		Data:       data,
	}
}

func ResponseFailed(httpCode int, statusCode common.StatusCode, message string, data interface{}) JSONResult {
	return JSONResult{
		Code:       httpCode,
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}

func ResponsePagination(page int, limit int, count int, total int, data interface{}, metaData interface{}) JSONResult {
	return JSONResult{
		Code:       http.StatusOK,
		StatusCode: common.StatusCodeSuccess,
		Message:    "Success",
		Data:       data,
		Meta: &PaginationMeta{
			Page:            page,
			TotalData:       total,
			TotalDataOnPage: count,
			TotalPage:       int(math.Ceil(float64(total) / float64(limit))),
			MetaData:        metaData,
		},
	}
}

func ResponseRecover(l *zap.Logger) fiber.ErrorHandler {

	return func(ctx *fiber.Ctx, err error) error {
		l.Error("Unexpected error", zap.Error(err), zap.String("method", ctx.Method()), zap.String("url", ctx.Path()))
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(JSONResult{
				Code:       fiber.StatusInternalServerError,
				StatusCode: common.StatusCodeInternalServerError,
				Message:    "Internal server error",
			})
	}
}
