-- name: CreateTeamMember :one
INSERT INTO teamMember (team, profile, admin)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetTeamMember :one
SELECT *
FROM teamMember
WHERE team = $1
    AND profile = $2
LIMIT 1;
-- name: GetAllTeamMembers :many
SELECT p.*,
    tm.admin
FROM teamMember tm
    INNER JOIN profile p ON tm.profile = p.id
WHERE tm.team = $1;
-- name: UpdateTeamMember :one
UPDATE teamMember
SET admin = $1,
    updatedAt = now()
WHERE team = $2
    AND profile = $3
RETURNING *;
-- name: GetTeams :many
SELECT t.name as teamName,
    t.description as teamDescription,
    t.id as teamId,
    o.name as orgName,
    o.description as orgDescription,
    tm.admin
FROM teamMember tm
    INNER JOIN team t ON tm.team = t.id
    INNER JOIN organisation o ON o.id = t.organisation
WHERE tm.profile = $1;
-- name: GetAssignableOrganisations :many
SELECT DISTINCT o.*
FROM teamMember tm
    INNER JOIN team t ON tm.team = t.id
    INNER JOIN organisation o ON o.id = t.organisation
WHERE tm.profile = $1;