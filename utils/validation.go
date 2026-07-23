package utils

import (
	"github.com/go-playground/validator/v10"
)

// function to send meaningful error messages for field validation
func Validate(errs validator.ValidationErrors) any {
	errMessages := []string{}

	for _, err := range errs {
		switch err.Tag() {
		case "required":
			errMessages = append(errMessages, err.Field() + " is required field")
		case "email":
			errMessages = append(errMessages, err.Field() + " field has invalid email format")
		case "len":
			errMessages = append(errMessages, err.Field() + " field must be exactly " + err.Param() + " characters long")
		case "min":
			errMessages = append(errMessages, err.Field() + " field must be at least " + err.Param() + " characters long")
		case "digits":
			errMessages = append(errMessages, err.Field() + " field must only contain numbers")
		}
	}

	return errMessages
}