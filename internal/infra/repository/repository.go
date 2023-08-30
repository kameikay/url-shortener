package repository

import (
	"context"
	"database/sql"

	db "github.com/kameikay/url-shortener/internal/infra/db"
)

type Repository struct {
	dbConnection *sql.DB
	*db.Queries
}

func NewRepository(database *sql.DB) *Repository {
	return &Repository{
		dbConnection: database,
		Queries:      db.New(database),
	}
}

func (r *Repository) Insert(ctx context.Context, code string, url string) error {
	err := r.Queries.InsertUrl(ctx, db.InsertUrlParams{
		Url:  url,
		Code: code,
	})
	return err
}

func (r *Repository) Find(ctx context.Context, code string) (string, error) {
	data, err := r.Queries.GetUrlByCode(ctx, code)
	if err != nil {
		return "", err
	}

	return data.Url, nil
}
