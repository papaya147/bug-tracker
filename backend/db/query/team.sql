-- name: CreateTeam :one
INSERT INTO team(id, name, description, organisation)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetOrganisationTeams :many
SELECT *
FROM team
WHERE organisation = $1;
-- name: UpdateTeam :one
UPDATE team
SET name = $1,
    description = $2,
    updatedAt = now()
WHERE id = $3
    AND organisation = $4
RETURNING *;
-- name: GetTeamOrganisation :one
SELECT organisation
FROM team
WHERE id = $1;