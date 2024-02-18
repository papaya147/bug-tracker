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