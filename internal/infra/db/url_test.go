package db

import (
	"context"
	"testing"

	"github.com/kameikay/url-shortener/tests/integration"
	"github.com/stretchr/testify/assert"
)

func NewDbQueries(t *testing.T) *Queries {
	queries := New(integration.TestDB)
	return queries
}

func TestInsertUrl(t *testing.T) {
	queries := NewDbQueries(t)
	err := queries.InsertUrl(context.Background(), InsertUrlParams{
		Url:  "https://www.google.com",
		Code: "123456",
	})
	assert.NoError(t, err)
}

func TestGetUrlByCode(t *testing.T) {
	queries := NewDbQueries(t)
	err := queries.InsertUrl(context.Background(), InsertUrlParams{
		Url:  "https://www.google.com",
		Code: "654321",
	})
	data, err := queries.GetUrlByCode(context.Background(), "654321")
	assert.NoError(t, err)
	assert.Equal(t, "https://www.google.com", data.Url)
}
