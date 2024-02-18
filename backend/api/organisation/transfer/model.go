package transfer

import "github.com/google/uuid"

type transferRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type transferResponse struct {
	Id           uuid.UUID `json:"id"`
	Organisation uuid.UUID `json:"organisation"`
	FromProfile  uuid.UUID `json:"from_profile"`
	ToProfile    uuid.UUID `json:"to_profile"`
	Completed    bool      `json:"completed"`
	CreatedAt    int64     `json:"created_at"`
}

type transferId struct {
	Id string `validate:"required,uuid"`
}

type transferResponseStatus struct {
	Id     string `validate:"required,uuid"`
	Status string `validate:"required,oneof=true false"`
}
