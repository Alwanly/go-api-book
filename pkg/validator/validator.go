package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func NewValidator() (IValidatorService, error) {
	// create validator
	v := validator.New()

	// register custom validators

	// register english translator
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	// register english translator
	_ = en_translations.RegisterDefaultTranslations(v, trans)

	return &Service{
		Validate:      v,
		Translator:    trans,
		DefaultLocale: "en",
		TranslateLanguage: map[string]ut.Translator{
			"en": trans,
		},
	}, nil
}

func (s *Service) ValidateStruct(input interface{}) error {
	return s.Validate.Struct(input)
}

func (s *Service) TranslateError(err error) []ValidationError {
	// check if err is nil
	if err == nil {
		return nil
	}

	// check if err is validator.ValidationErrors
	validatorErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil
	}

	// translate each error
	errors := []ValidationError{}
	for _, e := range validatorErrs {
		errors = append(errors, ValidationError{
			Field:   e.Field(),
			Value:   e.Value(),
			Message: e.Translate(s.Translator),
		})
	}

	return errors
}

func (s *Service) TranslateToLocale(err error, locale string) []ValidationError {
	// check if err is nil
	if err == nil {
		return nil
	}

	// check if err is validator.ValidationErrors
	validatorErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil
	}

	_, ok = s.TranslateLanguage[locale]
	if !ok {
		locale = s.DefaultLocale
	}

	// translate each error
	errors := []ValidationError{}
	for _, e := range validatorErrs {
		errors = append(errors, ValidationError{
			Field:   e.Field(),
			Value:   e.Value(),
			Message: e.Translate(s.TranslateLanguage[locale]),
		})
	}

	return errors
}
