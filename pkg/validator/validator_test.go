package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

func TestNewValidator(t *testing.T) {
	v, err := NewValidator()
	assert.NoError(t, err)
	assert.NotNil(t, v)
}

func TestValidateStruct(t *testing.T) {
	v, _ := NewValidator()

	validStruct := TestStruct{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	err := v.ValidateStruct(validStruct)
	assert.NoError(t, err)

	invalidStruct := TestStruct{
		Name:  "",
		Email: "invalid-email",
	}
	err = v.ValidateStruct(invalidStruct)
	assert.Error(t, err)
}

func TestTranslateError(t *testing.T) {
	v, _ := NewValidator()

	invalidStruct := TestStruct{
		Name:  "",
		Email: "invalid-email",
	}
	err := v.ValidateStruct(invalidStruct)
	assert.Error(t, err)

	translatedErrors := v.TranslateError(err)
	assert.Len(t, translatedErrors, 2)
	assert.Equal(t, "Name", translatedErrors[0].Field)
	assert.Equal(t, "Email", translatedErrors[1].Field)
}

func TestTranslateToLocale(t *testing.T) {
	v, _ := NewValidator()

	invalidStruct := TestStruct{
		Name:  "",
		Email: "invalid-email",
	}
	err := v.ValidateStruct(invalidStruct)
	assert.Error(t, err)

	translatedErrors := v.TranslateToLocale(err, "en")
	assert.Len(t, translatedErrors, 2)
	assert.Equal(t, "Name", translatedErrors[0].Field)
	assert.Equal(t, "Email", translatedErrors[1].Field)
}
