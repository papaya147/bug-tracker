package teammember

import "github.com/google/uuid"

type createTeamMemberRequest struct {
	TeamId uuid.UUID `json:"team_id" validate:"required"`
	Email  string    `json:"email" validate:"required,email"`
	Admin  bool      `json:"admin"`
}

type updateTeamMemberRequest struct {
	TeamId    uuid.UUID `json:"team_id" validate:"required"`
	ProfileId uuid.UUID `json:"profile_id" validate:"required"`
	Admin     bool      `json:"admin"`
}

type teamMemberResponse struct {
	TeamId    uuid.UUID `json:"team_id"`
	ProfileId uuid.UUID `json:"profile_id"`
	Admin     bool      `json:"admin"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
}

type teamId struct {
	Id string `json:"id" validate:"required,uuid"`
}

type profileResponse struct {
	Id        uuid.UUID `json:"id"`
	Admin     bool      `json:"admin"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Verified  bool      `json:"verified"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
}
