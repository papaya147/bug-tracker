package transfer

import "github.com/google/uuid"

type transferInput struct {
	Email string `json:"email" validate:"required,email" example:"asrivatsa6@gmail.com"`
}

type transferOutput struct {
	Id           uuid.UUID `json:"id" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	Organisation uuid.UUID `json:"organisation" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	FromProfile  uuid.UUID `json:"from_profile" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	ToProfile    uuid.UUID `json:"to_profile" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	Completed    bool      `json:"completed" example:"false"`
	CreatedAt    int64     `json:"created_at" example:"1710579130"`
}

type transferId struct {
	Id string `validate:"required,uuid" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
}

type transferResponseInput struct {
	Id     string `validate:"required,uuid" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	Status string `validate:"required,oneof=true false" example:"true"`
}
