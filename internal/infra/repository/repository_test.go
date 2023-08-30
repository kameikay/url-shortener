package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func createDb(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestNewRepository(t *testing.T) {
	db, _ := createDb(t)
	repository := NewRepository(db)
	assert.NotNil(t, repository)
}

func TestInsert(t *testing.T) {
	dbt, mock := createDb(t)
	repo := NewRepository(dbt)

	query := `INSERT INTO urls`

	testCases := []struct {
		name         string
		code         string
		url          string
		expected     error
		expectations func(mock sqlmock.Sqlmock)
	}{
		{
			name:     "success",
			code:     "123456",
			url:      "https://www.google.com",
			expected: nil,
			expectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.expectations(mock)
			err := repo.Insert(context.Background(), tc.code, tc.url)
			assert.Equal(t, tc.expected, err)
		})
	}
}

func TestFind(t *testing.T) {
	dbt, mock := createDb(t)
	repo := NewRepository(dbt)

	query := `SELECT id, url FROM urls`

	testCases := []struct {
		name         string
		code         string
		url          string
		expected     error
		expectedUrl  string
		expectations func(mock sqlmock.Sqlmock)
	}{
		{
			name:        "success",
			code:        "123456",
			url:         "https://www.google.com",
			expected:    nil,
			expectedUrl: "https://www.google.com",
			expectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"id", "url"}).AddRow(1, "https://www.google.com"))
			},
		},
		{
			name:        "error",
			code:        "123456",
			url:         "https://www.google.com",
			expected:    sql.ErrNoRows,
			expectedUrl: "",
			expectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.expectations(mock)
			url, err := repo.Find(context.Background(), tc.code)
			assert.Equal(t, tc.expected, err)
			assert.Equal(t, tc.expectedUrl, url)
		})
	}
}
