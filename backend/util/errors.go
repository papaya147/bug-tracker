package util

import (
	"errors"
	"net/http"
)

var (
	ErrInternal               = errors.New("internal error")
	ErrDatabase               = errors.New("database error")
	ErrEmailExists            = errors.New("email exists")
	ErrInvalidToken           = errors.New("invalid token")
	ErrUserNotFound           = errors.New("user not found or not verified")
	ErrProfileAlreadyVerified = errors.New("profile already verified")
	ErrWrongPassword          = errors.New("wrong password")
)

var CustomErrors = map[error]int{
	ErrInternal:               http.StatusInternalServerError,
	ErrDatabase:               http.StatusInternalServerError,
	ErrEmailExists:            http.StatusBadRequest,
	ErrInvalidToken:           http.StatusBadRequest,
	ErrUserNotFound:           http.StatusBadRequest,
	ErrProfileAlreadyVerified: http.StatusBadRequest,
	ErrWrongPassword:          http.StatusBadRequest,
}
