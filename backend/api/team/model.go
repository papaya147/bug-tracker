package team

import "github.com/google/uuid"

type teamOutput struct {
	Id                      uuid.UUID `json:"id" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	OrganisationName        string    `json:"organisation_name" example:"buggy org"`
	OrganisationDescription string    `json:"organisation_description" example:"The organisation for bugs!"`
	Name                    string    `json:"name" example:"testing team"`
	Description             string    `json:"description" example:"The Buggy testing team!"`
	Admin                   bool      `json:"admin" example:"false"`
}
