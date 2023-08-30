package usecases

import (
	"context"

	"github.com/kameikay/url-shortener/internal/dtos"
	"github.com/kameikay/url-shortener/internal/entities"
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

func (uc *GenerateCodeUseCase) Execute(ctx context.Context, input dtos.CreateCodeInputDTO) (string, error) {
	codeEntity := entities.NewCode(input.Url)
	codeEntity.GenerateCode()

	err := uc.repository.Insert(ctx, codeEntity.Code, codeEntity.Url)
	if err != nil {
		return "", err
	}

	return codeEntity.Code, nil
}
