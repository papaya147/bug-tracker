-- name: CreateBug :one
INSERT INTO bug(
        id,
        name,
        description,
        status,
        priority,
        assignedTo,
        assignedBy
    )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;
-- name: GetBug :one
SELECT *
FROM bug b
WHERE id = $1;
-- name: GetActiveBugsByProfile :many
SELECT b.*
FROM bug b
    INNER JOIN team t ON b.assignedTo = t.id
    INNER JOIN teamMember tm ON t.id = tm.team
WHERE tm.profile = $1
    AND completed = FALSE
ORDER BY b.priority DESC,
    b.status DESC;
-- name: GetBugsByAssignedTeam :many
SELECT *
FROM bug
WHERE assignedTo = $1
ORDER BY priority DESC,
    status DESC,
    completed DESC;
-- name: GetBugsByAsigneeTeam :many
SELECT b.*
FROM bug b
    INNER JOIN teamMember tm ON b.assignedBy = tm.profile
WHERE tm.team = $1
ORDER BY b.priority DESC,
    b.status DESC,
    b.completed DESC;
-- name: CloseBug :one
UPDATE bug
SET completed = TRUE,
    closedBy = $1,
    remarks = $2,
    closedAt = NOW()
WHERE id = $3
RETURNING *;
-- name: DeleteBug :one
DELETE FROM bug
WHERE id = $1
RETURNING *;
-- name: UpdateBug :one
UPDATE bug
SET name = $1,
    description = $2,
    status = $3,
    priority = $4,
    assignedTo = $5,
    updatedAt = NOW()
WHERE id = $6
RETURNING *;