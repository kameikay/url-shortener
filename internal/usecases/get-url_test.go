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

type GetUrlUseCaseSuite struct {
	suite.Suite
	ctrl          *gomock.Controller
	GetUrlUseCase *GetUrlUseCase
	repository    *mock.MockRepositoryInterface
	mock          sqlmock.Sqlmock
}

func TestGetUrlUseCaseStart(t *testing.T) {
	suite.Run(t, new(GetUrlUseCaseSuite))
}

func (suite *GetUrlUseCaseSuite) GetUrlUseCaseSuiteDown() {
	defer suite.ctrl.Finish()
}

func (suite *GetUrlUseCaseSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	_, dbMock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	suite.repository = mock.NewMockRepositoryInterface(suite.ctrl)
	suite.GetUrlUseCase = NewGetUrlUseCase(suite.repository)
	suite.mock = dbMock
}

func (suite *GetUrlUseCaseSuite) TestExecute() {
	testCases := []struct {
		name          string
		expectedError error
		expectations  func(repo *mock.MockRepositoryInterface)
	}{
		{
			name:          "success",
			expectedError: nil,
			expectations: func(repo *mock.MockRepositoryInterface) {
				repo.EXPECT().Find(context.Background(), "123456").Return("url.com", nil)
			},
		},
		{
			name:          "error",
			expectedError: errors.New("error"),
			expectations: func(repo *mock.MockRepositoryInterface) {
				repo.EXPECT().Find(context.Background(), "123456").Return("", errors.New("error"))
			},
		},
	}

	for _, tc := range testCases {
		suite.T().Run(tc.name, func(t *testing.T) {
			tc.expectations(suite.repository)

			_, err := suite.GetUrlUseCase.Execute(context.Background(), "123456")

			suite.Equal(tc.expectedError, err)
		})
	}
}
