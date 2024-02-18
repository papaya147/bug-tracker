-- name: CreateOrganisation :one
INSERT INTO organisation(id, name, description, owner)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetOrganisation :one
SELECT *
FROM organisation
WHERE owner = $1
LIMIT 1;
-- name: UpdateOrganisation :one
UPDATE organisation
SET name = $1,
    description = $2,
    updatedAt = now()
WHERE owner = $3
RETURNING *;
-- name: UpdateOrganisationOwner :one
UPDATE organisation
SET owner = $1,
    updatedAt = now()
WHERE id = $2
RETURNING *;