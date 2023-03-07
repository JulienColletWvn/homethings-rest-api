-- name: CreateDevice :one
INSERT INTO devices (id, name, location)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetDevices :many
SELECT *
FROM devices;
-- name: GetDevice :one
SELECT *
FROM devices
WHERE devices.id = $1;