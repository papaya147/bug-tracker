-- name: CreateProfile :one
INSERT INTO profile (id, name, email, password)
VALUES ($1, $2, $3, $4)
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