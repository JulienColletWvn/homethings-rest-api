-- name: CreateData :one
INSERT INTO datas (data_type_id, value)
VALUES ($1, $2)
RETURNING *;
-- name: GetAllDeviceDatas :many
SELECT *
FROM datas
    JOIN data_types ON datas.data_type_id = data_types.id
WHERE data_types.device_id = $1;
-- name: GetDeviceDatas :many
SELECT *
FROM datas
    JOIN data_types ON datas.data_type_id = data_types.id
WHERE data_types.device_id = $1
    AND data_types.key = $2;