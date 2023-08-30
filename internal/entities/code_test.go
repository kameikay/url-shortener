package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCode(t *testing.T) {
	code := NewCode("google.com")

	assert.Equal(t, "google.com", code.Url)
	assert.Equal(t, "", code.Code)
}

func TestGenerateCode(t *testing.T) {
	code := NewCode("google.com")

	code.GenerateCode()

	assert.NotEqual(t, code.Code, "")
	assert.Equal(t, 6, len(code.Code))
}
