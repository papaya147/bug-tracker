package teammember

import "github.com/google/uuid"

type createTeamMemberInput struct {
	TeamId uuid.UUID `json:"team_id" validate:"required" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	Email  string    `json:"email" validate:"required,email" example:"asrivatsa6@gmail.com"`
	Admin  bool      `json:"admin" example:"false"`
}

type updateTeamMemberInput struct {
	TeamId    uuid.UUID `json:"team_id" validate:"required" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	ProfileId uuid.UUID `json:"profile_id" validate:"required" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	Admin     bool      `json:"admin" example:"false"`
}

type teamMemberOutput struct {
	TeamId    uuid.UUID `json:"team_id" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	ProfileId uuid.UUID `json:"profile_id" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	Admin     bool      `json:"admin" example:"false"`
	CreatedAt int64     `json:"created_at" example:"1710579130"`
	UpdatedAt int64     `json:"updated_at" example:"1710579130"`
}

type teamId struct {
	Id string `json:"id" validate:"required,uuid" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
}

type profileOutput struct {
	Id        uuid.UUID `json:"id" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	Admin     bool      `json:"admin" example:"false"`
	Name      string    `json:"name" example:"abhinav"`
	Email     string    `json:"email" example:"asrivatsa6@gmail.com"`
	Verified  bool      `json:"verified" example:"true"`
	CreatedAt int64     `json:"created_at" example:"1710579130"`
	UpdatedAt int64     `json:"updated_at" example:"1710579130"`
}
