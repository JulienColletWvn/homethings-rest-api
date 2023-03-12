-- name: CreateData :one
INSERT INTO datas (data_type_id, value)
VALUES ($1, $2)
RETURNING *;
-- name: GetLastData :one
SELECT datas.created_at,
    datas."value",
    data_types."key",
    data_types.unit,
    devices."name",
    devices."location"
FROM datas
    JOIN data_types ON data_types.id = datas.data_type_id
    JOIN devices ON devices.id = data_types.device_id
WHERE data_types.device_id = $1
    AND data_types.key = $2
ORDER BY datas.created_at DESC
LIMIT 1;
-- name: GetDatas :many
SELECT datas.created_at,
    datas."value",
    data_types."key",
    data_types.unit,
    devices."name",
    devices."location"
FROM datas
    JOIN data_types ON data_types.id = datas.data_type_id
    JOIN devices ON devices.id = data_types.device_id
WHERE data_types.device_id = $1
    AND data_types.key = $2
    AND datas.created_at BETWEEN $3 AND $4
ORDER BY datas.created_at ASC;