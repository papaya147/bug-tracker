-- name: CreateBug :one
INSERT INTO bug(
        id,
        name,
        description,
        priority,
        assignedTo,
        assignedByProfile,
        assignedByTeam
    )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;
-- name: GetBug :one
SELECT *
FROM bug
WHERE id = $1;
-- name: GetActiveBugsByAssignedProfile :many
SELECT b.*
FROM bug b
    INNER JOIN team t ON b.assignedTo = t.id
    INNER JOIN teamMember tm ON t.id = tm.team
WHERE tm.profile = $1
    AND completed = FALSE
ORDER BY b.priority,
    b.status,
    b.createdAt;
-- name: GetBugsByAssigneeProfile :many
SELECT b.*
FROM bug b
    INNER JOIN team t ON b.assignedByTeam = t.id
    INNER JOIN teamMember tm ON t.id = tm.team
WHERE tm.profile = $1
ORDER BY b.priority,
    b.status,
    b.completed,
    b.createdAt;
-- name: GetBugsByAssignedTeam :many
SELECT *
FROM bug
WHERE assignedTo = $1
ORDER BY priority,
    status,
    completed,
    createdAt;
-- name: GetBugsByAssigneeTeam :many
SELECT *
FROM bug
WHERE assignedByTeam = $1
ORDER BY priority,
    status,
    completed,
    createdAt;
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
    updatedAt = NOW()
WHERE id = $5
RETURNING *;