package validator

import (
	"net/http"

	"github.com/Alwanly/go-codebase/pkg/contract"
	"github.com/Alwanly/go-codebase/pkg/logger"
	"github.com/Alwanly/go-codebase/pkg/wrapper"
	"go.uber.org/zap"
)

type ModelValidationError struct {
	Code         int
	ResponseBody wrapper.JSONResult
}

func (e *ModelValidationError) Error() string {
	return "Failed to validate request body"
}

func ValidateModel(log *zap.Logger, v IValidatorService, m interface{}) error {
	// create local logger
	l := logger.WithID(log, ContextName, "ValidateModel")

	// try validate model
	if err := v.ValidateStruct(m); err != nil {
		// log error
		l.Error(contract.ErrorValidatePayload, zap.Error(err))

		// translate error
		localizedErr := v.TranslateError(err)

		// return error
		result := wrapper.ResponseFailed(http.StatusBadRequest, contract.StatusCodeValidationFailed, contract.ErrorValidatePayload, localizedErr)
		return &ModelValidationError{
			Code:         result.Code,
			ResponseBody: result,
		}
	}

	return nil
}
