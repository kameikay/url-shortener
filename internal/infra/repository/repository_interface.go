package repository

import "context"

type RepositoryInterface interface {
	Insert(ctx context.Context, code string, url string) error
	Find(ctx context.Context, code string) (string, error)
}
