-- name: CreateDataType :one
INSERT INTO data_types (key, unit, device_id)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetDeviceDataTypes :many
SELECT *
FROM data_types
WHERE data_types.device_id = $1;