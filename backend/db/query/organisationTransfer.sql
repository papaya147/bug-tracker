-- name: CreateOrganisationTransfer :one
INSERT INTO organisationTransfer (organisation, fromProfile, toProfile)
VALUES ($1, $2, $3)
RETURNING *;