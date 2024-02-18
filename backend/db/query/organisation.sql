-- name: CreateOrganisation :one
INSERT INTO organisation(id, name, description, owner)
VALUES ($1, $2, $3, $4)
RETURNING *;