package utils

import (
	"net/http"

	"github.com/Alwanly/go-codebase/pkg/common"
	"github.com/Alwanly/go-codebase/pkg/logger"
	"github.com/Alwanly/go-codebase/pkg/validator"
	"go.uber.org/zap"
)

type ModelValidationError struct {
	Code         int
	ResponseBody JSONResult
}

func (e *ModelValidationError) Error() string {
	return "Failed to validate request body"
}

func ValidateModel(log *zap.Logger, v validator.IValidatorService, m interface{}) error {
	// create local logger
	l := logger.WithID(log, ContextName, "ValidateModelV2")

	// try validate model
	if err := v.ValidateStruct(m); err != nil {
		// log error
		l.Error(common.ErrorValidatePayload, zap.Error(err))

		// translate error
		localizedErr := v.TranslateError(err)

		// return error
		result := ResponseFailed(http.StatusBadRequest, common.StatusCodeValidationFailed, common.ErrorValidatePayload, localizedErr)
		return &ModelValidationError{
			Code:         result.Code,
			ResponseBody: result,
		}
	}

	return nil
}
