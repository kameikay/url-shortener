package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	mock "github.com/kameikay/url-shortener/internal/infra/repository/mocks"
	"github.com/stretchr/testify/suite"
)

type GenerateCodeUseCaseSuite struct {
	suite.Suite
	ctrl                *gomock.Controller
	GenerateCodeUseCase *GenerateCodeUseCase
	repository          *mock.MockRepositoryInterface
	mock                sqlmock.Sqlmock
}

func TestGenerateCodeUseCaseStart(t *testing.T) {
	suite.Run(t, new(GenerateCodeUseCaseSuite))
}

func (suite *GenerateCodeUseCaseSuite) GenerateCodeUseCaseSuiteDown() {
	defer suite.ctrl.Finish()
}

func (suite *GenerateCodeUseCaseSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	_, dbMock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	suite.repository = mock.NewMockRepositoryInterface(suite.ctrl)
	suite.GenerateCodeUseCase = NewGenerateCodeUseCase(suite.repository)
	suite.mock = dbMock
}

func (suite *GenerateCodeUseCaseSuite) TestExecute() {
	testCases := []struct {
		name          string
		expectedError error
		expectations  func(repo *mock.MockRepositoryInterface)
	}{
		{
			name:          "success",
			expectedError: nil,
			expectations: func(repo *mock.MockRepositoryInterface) {
				repo.EXPECT().Insert(context.Background(), gomock.Any(), "google.com").Return(nil)
			},
		},
		{
			name:          "error",
			expectedError: errors.New("error"),
			expectations: func(repo *mock.MockRepositoryInterface) {
				repo.EXPECT().Insert(context.Background(), gomock.Any(), "google.com").Return(errors.New("error"))
			},
		},
	}

	for _, tc := range testCases {
		suite.T().Run(tc.name, func(t *testing.T) {
			tc.expectations(suite.repository)

			_, err := suite.GenerateCodeUseCase.Execute(context.Background(), "google.com")

			suite.Equal(tc.expectedError, err)
		})
	}
}
