-- name: CreateTeam :one
INSERT INTO team(id, name, description, organisation)
VALUES ($1, $2, $3, $4)
RETURNING *;