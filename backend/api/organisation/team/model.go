package team

import "github.com/google/uuid"

type createTeamInput struct {
	Name        string `json:"name" validate:"required,max=20" example:"testing team"`
	Description string `json:"description" validate:"required,max=100" example:"The Buggy testing team!"`
}

type teamOutput struct {
	Id                      uuid.UUID `json:"id" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	OrganisationName        string    `json:"organisation_name" example:"buggy org"`
	OrganisationDescription string    `json:"organisation_description" example:"The organisation for bugs!"`
	Name                    string    `json:"name" example:"testing team"`
	Description             string    `json:"description" example:"The Buggy testing team!"`
	CreatedAt               int64     `json:"created_at" example:"1710579130"`
	UpdatedAt               int64     `json:"updated_at" example:"1710579130"`
}

type teamId struct {
	Id string `json:"id" validate:"required,uuid" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
}
