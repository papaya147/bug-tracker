// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: teamMember.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createTeamMember = `-- name: CreateTeamMember :one
INSERT INTO teamMember (team, profile, admin)
VALUES ($1, $2, $3)
RETURNING team, profile, admin, createdat, updatedat
`

type CreateTeamMemberParams struct {
	Team    uuid.UUID `json:"team"`
	Profile uuid.UUID `json:"profile"`
	Admin   bool      `json:"admin"`
}

func (q *Queries) CreateTeamMember(ctx context.Context, arg CreateTeamMemberParams) (Teammember, error) {
	row := q.db.QueryRow(ctx, createTeamMember, arg.Team, arg.Profile, arg.Admin)
	var i Teammember
	err := row.Scan(
		&i.Team,
		&i.Profile,
		&i.Admin,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}

const getAllTeamMembers = `-- name: GetAllTeamMembers :many
SELECT p.id, p.tokenid, p.name, p.email, p.password, p.verified, p.createdat, p.updatedat,
    tm.admin
FROM teamMember tm
    INNER JOIN profile p ON tm.profile = p.id
WHERE tm.team = $1
`

type GetAllTeamMembersRow struct {
	ID        uuid.UUID `json:"id"`
	Tokenid   uuid.UUID `json:"tokenid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Verified  bool      `json:"verified"`
	Createdat time.Time `json:"createdat"`
	Updatedat time.Time `json:"updatedat"`
	Admin     bool      `json:"admin"`
}

func (q *Queries) GetAllTeamMembers(ctx context.Context, team uuid.UUID) ([]GetAllTeamMembersRow, error) {
	rows, err := q.db.Query(ctx, getAllTeamMembers, team)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllTeamMembersRow{}
	for rows.Next() {
		var i GetAllTeamMembersRow
		if err := rows.Scan(
			&i.ID,
			&i.Tokenid,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.Verified,
			&i.Createdat,
			&i.Updatedat,
			&i.Admin,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTeamMember = `-- name: GetTeamMember :one
SELECT team, profile, admin, createdat, updatedat
FROM teamMember
WHERE team = $1
    AND profile = $2
LIMIT 1
`

type GetTeamMemberParams struct {
	Team    uuid.UUID `json:"team"`
	Profile uuid.UUID `json:"profile"`
}

func (q *Queries) GetTeamMember(ctx context.Context, arg GetTeamMemberParams) (Teammember, error) {
	row := q.db.QueryRow(ctx, getTeamMember, arg.Team, arg.Profile)
	var i Teammember
	err := row.Scan(
		&i.Team,
		&i.Profile,
		&i.Admin,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}

const getTeams = `-- name: GetTeams :many
SELECT t.name as teamName,
    t.description as teamDescription,
    t.id as teamId,
    o.name as orgName,
    o.description as orgDescription,
    tm.admin
FROM teamMember tm
    INNER JOIN team t ON tm.team = t.id
    INNER JOIN organisation o ON o.id = t.organisation
WHERE tm.profile = $1
`

type GetTeamsRow struct {
	Teamname        string    `json:"teamname"`
	Teamdescription string    `json:"teamdescription"`
	Teamid          uuid.UUID `json:"teamid"`
	Orgname         string    `json:"orgname"`
	Orgdescription  string    `json:"orgdescription"`
	Admin           bool      `json:"admin"`
}

func (q *Queries) GetTeams(ctx context.Context, profile uuid.UUID) ([]GetTeamsRow, error) {
	rows, err := q.db.Query(ctx, getTeams, profile)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetTeamsRow{}
	for rows.Next() {
		var i GetTeamsRow
		if err := rows.Scan(
			&i.Teamname,
			&i.Teamdescription,
			&i.Teamid,
			&i.Orgname,
			&i.Orgdescription,
			&i.Admin,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTeamMember = `-- name: UpdateTeamMember :one
UPDATE teamMember
SET admin = $1,
    updatedAt = now()
WHERE team = $2
    AND profile = $3
RETURNING team, profile, admin, createdat, updatedat
`

type UpdateTeamMemberParams struct {
	Admin   bool      `json:"admin"`
	Team    uuid.UUID `json:"team"`
	Profile uuid.UUID `json:"profile"`
}

func (q *Queries) UpdateTeamMember(ctx context.Context, arg UpdateTeamMemberParams) (Teammember, error) {
	row := q.db.QueryRow(ctx, updateTeamMember, arg.Admin, arg.Team, arg.Profile)
	var i Teammember
	err := row.Scan(
		&i.Team,
		&i.Profile,
		&i.Admin,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}
