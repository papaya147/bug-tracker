package organisation

import "github.com/google/uuid"

type createOrganisationInput struct {
	Name        string `json:"name" validate:"required,max=20" example:"buggy org"`
	Description string `json:"description" validate:"required,max=100" example:"The organisation for bugs!"`
}

type organisationOutput struct {
	ID          uuid.UUID `json:"id" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	Name        string    `json:"name" example:"buggy org"`
	Description string    `json:"description" example:"The organisation for bugs!"`
	Owner       uuid.UUID `json:"owner" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	CreatedAt   int64     `json:"created_at" example:"1710579130"`
	UpdatedAt   int64     `json:"updated_at" example:"1710579130"`
}
