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
SELECT p.*
FROM teamMember tm
    INNER JOIN profile p ON tm.profile = p.id
WHERE tm.team = $1;