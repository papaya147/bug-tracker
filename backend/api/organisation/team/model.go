package team

import "github.com/google/uuid"

type createTeamRequest struct {
	Name        string `json:"name" validate:"required,max=20"`
	Description string `json:"description" validate:"required,max=100"`
}

type teamResponse struct {
	Id                      uuid.UUID `json:"id"`
	OrganisationName        string    `json:"organisation_name"`
	OrganisationDescription string    `json:"organisation_description"`
	Name                    string    `json:"name"`
	Description             string    `json:"description"`
	CreatedAt               int64     `json:"created_at"`
	UpdatedAt               int64     `json:"updated_at"`
}

type teamId struct {
	Id string `json:"id" validate:"required,uuid"`
}
