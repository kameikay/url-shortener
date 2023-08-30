package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/kameikay/url-shortener/internal/infra/repository"
)

type GenerateCodeUseCase struct {
	repository repository.RepositoryInterface
}

func NewGenerateCodeUseCase(repository repository.RepositoryInterface) *GenerateCodeUseCase {
	return &GenerateCodeUseCase{
		repository: repository,
	}
}

func (uc *GenerateCodeUseCase) Execute(ctx context.Context, url string) (string, error) {
	id := uuid.New()
	code := id.String()[0:6]

	err := uc.repository.Insert(ctx, code, url)
	if err != nil {
		return "", err
	}

	return code, nil
}
