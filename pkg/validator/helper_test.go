package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestValidateModel(t *testing.T) {
	log, _ := zap.NewDevelopment()
	v, _ := NewValidator()

	validStruct := TestStruct{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	err := ValidateModel(log, v, validStruct)
	assert.NoError(t, err)

	invalidStruct := TestStruct{
		Name:  "",
		Email: "invalid-email",
	}
	err = ValidateModel(log, v, invalidStruct)
	assert.Error(t, err)

	modelValidationErr, ok := err.(*ModelValidationError)
	assert.True(t, ok)
	assert.Equal(t, 400, modelValidationErr.Code)
}
