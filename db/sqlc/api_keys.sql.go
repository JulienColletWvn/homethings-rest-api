// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: api_keys.sql

package db

import (
	"context"
)

const getApiKey = `-- name: GetApiKey :one
SELECT id, expiry_date
FROM api_keys
WHERE api_keys.id = $1
`

func (q *Queries) GetApiKey(ctx context.Context, id string) (ApiKey, error) {
	row := q.db.QueryRow(ctx, getApiKey, id)
	var i ApiKey
	err := row.Scan(&i.ID, &i.ExpiryDate)
	return i, err
}
