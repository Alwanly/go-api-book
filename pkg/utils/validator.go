package utils

import (
	"go-codebase/pkg/logger"
	"go-codebase/pkg/validator"
	"net/http"

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
	l := logger.WithId(log, ContextName, "ValidateModelV2")

	// try validate model
	if err := v.ValidateStruct(m); err != nil {
		// log error
		l.Error(ErrorValidatePayload, zap.Error(err))

		// translate error
		localizedErr := v.TranslateError(err)

		// return error
		result := ResponseFailed(http.StatusBadRequest, StatusCodeValidationFailed, ErrorValidatePayload, localizedErr)
		return &ModelValidationError{
			Code:         result.Code,
			ResponseBody: result,
		}
	}

	return nil
}
