-- name: GetApiKey :one
SELECT *
FROM api_keys
WHERE api_keys.id = $1;