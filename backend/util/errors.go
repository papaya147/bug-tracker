package util

import (
	"errors"
	"net/http"
)

var (
	ErrInternal    = errors.New("internal error")
	ErrDatabase    = errors.New("database error")
	ErrEmailExists = errors.New("email exists")
)

var CustomErrors = map[error]int{
	ErrInternal:    http.StatusInternalServerError,
	ErrDatabase:    http.StatusInternalServerError,
	ErrEmailExists: http.StatusBadRequest,
}
