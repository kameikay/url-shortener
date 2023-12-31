// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: url.sql

package db

import (
	"context"
)

const getUrlByCode = `-- name: GetUrlByCode :one

SELECT id, url FROM urls WHERE code = $1
`

type GetUrlByCodeRow struct {
	ID  int32  `db:"id" json:"id"`
	Url string `db:"url" json:"url"`
}

func (q *Queries) GetUrlByCode(ctx context.Context, code string) (GetUrlByCodeRow, error) {
	row := q.db.QueryRowContext(ctx, getUrlByCode, code)
	var i GetUrlByCodeRow
	err := row.Scan(&i.ID, &i.Url)
	return i, err
}

const insertUrl = `-- name: InsertUrl :exec

INSERT INTO
    urls (id, url, code, created_at)
VALUES (DEFAULT, $1, $2, DEFAULT)
`

type InsertUrlParams struct {
	Url  string `db:"url" json:"url"`
	Code string `db:"code" json:"code"`
}

func (q *Queries) InsertUrl(ctx context.Context, arg InsertUrlParams) error {
	_, err := q.db.ExecContext(ctx, insertUrl, arg.Url, arg.Code)
	return err
}
