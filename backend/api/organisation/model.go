package organisation

import "github.com/google/uuid"

type createOrganisationRequest struct {
	Name        string `json:"name" validate:"required,max=20"`
	Description string `json:"description" validate:"required,max=100"`
}

type organisationResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Owner       uuid.UUID `json:"owner"`
	CreatedAt   int64     `json:"created_at"`
	UpdatedAt   int64     `json:"updated_at"`
}
