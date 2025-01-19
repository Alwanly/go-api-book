package validator

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string      `json:"field"`
	Value   interface{} `json:"value"`
	Message string      `json:"message"`
}

// ValidatorService is a struct that contains a validator and a translator.
type ValidatorService struct {
	Validate          *validator.Validate
	Translator        ut.Translator
	DefaultLocale     string
	TranslateLanguage map[string]ut.Translator
}

// IValidatorService is an interface that defines the methods for validating and translating errors.
type IValidatorService interface {
	// Validates the contents of a struct
	//
	// Parameters:
	//    - input: struct to be validated
	//
	// Returns:
	//    - error if any with all the validation errors
	ValidateStruct(input interface{}) error

	// Translates validation errors into human readable messages
	//
	// Parameters:
	//    - err: validation error
	//
	// Returns:
	//    - array of ValidationError
	TranslateError(err error) []ValidationError

	// Translates validation errors into locale specified
	//
	// Parameters:
	//    - err: validation error
	//		- locale: locale code
	//
	// Returns:
	//    - array of ValidationError
	TranslateToLocale(err error, locale string) []ValidationError
}
