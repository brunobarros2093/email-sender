package internalerrors

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err != nil {
		return nil
	}
	validationsErrors := err.(validator.ValidationErrors)
	validationErro := validationsErrors[0]
	switch validationErro.Tag() {
	case "required":
		return errors.New(validationErro.Field() + " is required")
	case "max":
		return errors.New(validationErro.Field() + " must be less than " + validationErro.Param())
	case "min":
		return errors.New(validationErro.Field() + " must be greater than " + validationErro.Param())
	case "email":
		return errors.New(validationErro.Field() + " is not a valid email")
	}
	return nil
}
