-- name: CreateOrganisationTransfer :one
INSERT INTO organisationTransfer (id, organisation, fromProfile, toProfile)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetActiveOrganisationTransfer :one
SELECT *
FROM organisationTransfer
WHERE organisation = $1
    AND completed = false
LIMIT 1;
-- name: GetIncomingOrganisationTransfers :many
SELECT ot.*,
    o.name as orgname,
    o.description,
    fp.name as profilename
FROM organisationTransfer ot
    INNER JOIN organisation o ON ot.organisation = o.id
    INNER JOIN profile fp ON ot.fromProfile = fp.id
WHERE toProfile = $1
    AND completed = false;
-- name: GetOutgoingOrganisationTransfers :many
SELECT *
FROM organisationTransfer
WHERE fromProfile = $1
    AND completed = false;
-- name: DeleteOrganisationTransfer :one
DELETE FROM organisationTransfer
WHERE id = $1
    AND completed = false
RETURNING *;
-- name: CompleteOrganisationTransfer :one
UPDATE organisationTransfer
SET completed = true
WHERE id = $1
    AND toProfile = $2
RETURNING *;