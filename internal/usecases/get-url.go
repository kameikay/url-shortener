package usecases

import (
	"context"

	"github.com/kameikay/url-shortener/internal/infra/repository"
)

type GetUrlUseCase struct {
	repository repository.RepositoryInterface
}

func NewGetUrlUseCase(repository repository.RepositoryInterface) *GetUrlUseCase {
	return &GetUrlUseCase{
		repository: repository,
	}
}

func (uc *GetUrlUseCase) Execute(ctx context.Context, code string) (string, error) {
	url, err := uc.repository.Find(ctx, code)
	if err != nil {
		return "", err
	}

	return url, nil
}
