-- name: CreateTeamMember :one
INSERT INTO teamMember (team, profile, admin)
VALUES ($1, $2, $3)
RETURNING *;