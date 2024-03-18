package util

import (
	"errors"
	"net/http"
)

type ErrorDetailer interface {
	ErrorDetail() *ErrorDetail
}

type ErrorDetail struct {
	Message  string `json:"message,omitempty"`
	Location string `json:"location,omitempty"`
	Value    any    `json:"value,omitempty"`
}

func (e *ErrorDetail) ErrorDetail() *ErrorDetail {
	return e
}

func (e *ErrorDetail) Error() string {
	return e.Message
}

type ErrorModel struct {
	Type     string         `json:"type,omitempty"`
	Title    string         `json:"title,omitempty"`
	Status   int            `json:"status,omitempty"`
	Detail   string         `json:"detail,omitempty"`
	Instance string         `json:"instance,omitempty"`
	Errors   []*ErrorDetail `json:"errors,omitempty"`
}

// Error implements ApiError.
func (e *ErrorModel) Error() string {
	return e.Model().Detail
}

// Model implements ApiError.
func (e *ErrorModel) Model() ErrorModel {
	return *e
}

// Write implements ApiError.
func (e *ErrorModel) Write(w http.ResponseWriter) {
	ErrorJson(w, e)
}

type ApiError interface {
	Error() string
	AddError(error) ApiError
	Model() ErrorModel
	Write(http.ResponseWriter)
}

func NewError(message string, errs ...error) ApiError {
	if len(errs) == 0 {
		return &ErrorModel{
			Detail: message,
		}
	}

	if converted, ok := errs[0].(ApiError); ok {
		return converted
	}

	firstErr := errs[0]
	details, ok := CustomErrors[firstErr]
	if !ok {
		details = CustomErrors[ErrGeneric]
	}

	detailText := details.Detail
	if message != "" {
		detailText = message
	}

	var errorDetails []*ErrorDetail
	for _, err := range errs {
		if converted, ok := err.(ErrorDetailer); ok {
			errorDetails = append(errorDetails, converted.ErrorDetail())
		} else {
			if err == nil {
				continue
			}
			errorDetails = append(errorDetails, &ErrorDetail{Message: err.Error()})
		}
	}

	return &ErrorModel{
		Title:  http.StatusText(details.Status),
		Status: details.Status,
		Detail: detailText,
		Errors: errorDetails,
	}
}

func (a *ErrorModel) AddError(err error) ApiError {
	if a.Errors == nil {
		a = NewError(a.Detail, err).(*ErrorModel)
		return a
	}

	if converted, ok := err.(ErrorDetailer); ok {
		a.Errors = append(a.Errors, converted.ErrorDetail())
	} else {
		if err == nil {
			return a
		}
		a.Errors = append(a.Errors, &ErrorDetail{Message: err.Error()})
	}
	return a
}

func NewErrorAndWrite(w http.ResponseWriter, errs ...error) {
	NewError("", errs...).Write(http.ResponseWriter(w))
}

var (
	ErrGeneric                 = errors.New("something went wrong")
	ErrInternal                = errors.New("internal error")
	ErrDatabase                = errors.New("database error")
	ErrEmailExists             = errors.New("email exists")
	ErrInvalidToken            = errors.New("invalid token")
	ErrProfileNotFound         = errors.New("profile not found or not verified")
	ErrProfileAlreadyVerified  = errors.New("profile already verified")
	ErrProfileNotVerified      = errors.New("profile not verified")
	ErrWrongPassword           = errors.New("wrong password")
	ErrEntityExists            = errors.New("entity already exists")
	ErrEntityDoesNotExist      = errors.New("entity does not exist")
	ErrTransferAlreadyExists   = errors.New("transfer already exists")
	ErrCannotTransferToSelf    = errors.New("cannot transfer to self")
	ErrUnauthorised            = errors.New("unauthorised")
	ErrTeamMemberAlreadyExists = errors.New("team member already exists")
	ErrInvalidCookie           = errors.New("invalid cookie")
	ErrDifferentOrganisation   = errors.New("teams not part of same organisation")
)

type CustomErrorDetails struct {
	Detail string
	Status int
}

var CustomErrors = map[error]CustomErrorDetails{
	ErrGeneric:                 {Detail: "Something went wrong", Status: http.StatusBadRequest},
	ErrInternal:                {Detail: "Internal server error", Status: http.StatusInternalServerError},
	ErrDatabase:                {Detail: "Database error", Status: http.StatusInternalServerError},
	ErrEmailExists:             {Detail: "Email already exists", Status: http.StatusBadRequest},
	ErrInvalidToken:            {Detail: "Invalid token", Status: http.StatusUnauthorized},
	ErrProfileNotFound:         {Detail: "Profile not found or not verified", Status: http.StatusBadRequest},
	ErrProfileAlreadyVerified:  {Detail: "Profile already verified", Status: http.StatusBadRequest},
	ErrProfileNotVerified:      {Detail: "Profile not verified", Status: http.StatusBadRequest},
	ErrWrongPassword:           {Detail: "Wrong password", Status: http.StatusBadRequest},
	ErrEntityExists:            {Detail: "Entity already exists", Status: http.StatusBadRequest},
	ErrEntityDoesNotExist:      {Detail: "Entity does not exist", Status: http.StatusBadRequest},
	ErrTransferAlreadyExists:   {Detail: "Transfer already exists", Status: http.StatusBadRequest},
	ErrCannotTransferToSelf:    {Detail: "Cannot transfer to self", Status: http.StatusBadRequest},
	ErrUnauthorised:            {Detail: "Unauthorised", Status: http.StatusForbidden},
	ErrTeamMemberAlreadyExists: {Detail: "Team member already exists", Status: http.StatusBadRequest},
	ErrInvalidCookie:           {Detail: "Invalid cookie", Status: http.StatusBadRequest},
}
