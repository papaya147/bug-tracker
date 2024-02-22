package util

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
)

func ValidateRequest(requestPayload any) error {
	validate := validator.New()

	err := validate.Struct(requestPayload)
	var param string
	if err != nil {
		var errs []*ErrorDetail
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return NewError("The request body was not of valid json", errors.New("invalid json"))
		}
		for _, err := range err.(validator.ValidationErrors) {
			param = fmt.Sprintf("%s: %s", err.Tag(), err.Param())
			if err.Param() == "" {
				param = err.Tag()
			}
			errs = append(errs, &ErrorDetail{
				Location: fmt.Sprintf("body.%s", err.Field()),
				Message:  fmt.Sprintf("field: %s, expected %s", err.Field(), param),
				Value:    err.Value(),
			})
		}
		apiError := NewError("The request body did not meet some requirements")
		for _, err := range errs {
			apiError = apiError.AddError(err)
		}
		return apiError
	}
	return nil
}
