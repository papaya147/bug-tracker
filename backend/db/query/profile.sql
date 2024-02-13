-- name: CreateProfile :one
INSERT INTO profile (id, tokenId, name, email, password)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: VerifyProfile :one
UPDATE profile
SET verified = true,
    updatedAt = now()
WHERE id = $1
RETURNING *;
-- name: GetProfile :one
SELECT *
FROM profile
WHERE id = $1;
-- name: UpdateProfile :one
UPDATE profile
SET name = $1,
    updatedAt = now()
WHERE id = $2
RETURNING *;
-- name: UpdatePassword :one
UPDATE profile
SET password = $1,
    updatedAt = now()
WHERE id = $2
RETURNING *;
-- name: UpdateTokenId :one
UPDATE profile
SET tokenId = $1
WHERE id = $2
RETURNING *;
-- name: GetProfileByEmail :one
SELECT *
FROM profile
WHERE email = $1;
-- name: UpdateTokenIdByEmail :one
UPDATE profile
SET tokenId = $1
WHERE email = $2
    AND verified = true
RETURNING *;