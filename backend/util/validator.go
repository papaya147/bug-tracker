package util

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

type validationErrors struct {
	Errors []string `json:"errors"`
}

func ValidateRequest(requestPayload any) error {
	validate := validator.New()

	err := validate.Struct(requestPayload)
	var param string
	if err != nil {
		var errs validationErrors
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return errors.New("invalid json")
		}
		for _, err := range err.(validator.ValidationErrors) {
			param = fmt.Sprintf("%s: %s", err.Tag(), err.Param())
			if err.Param() == "" {
				param = err.Tag()
			}
			errs.Errors = append(errs.Errors, fmt.Sprintf("field: %s, expected %s", err.Field(), param))
		}
		return errors.New(strings.Join(errs.Errors, ", "))
	}
	return nil
}
