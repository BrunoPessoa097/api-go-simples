package repository

import (
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/mocks"
	"github.com/go-openapi/testify/v2/assert"
)

// teste de list
func TestUsuarioRepositoryList(t *testing.T) {
	repo := NewUsuarioRepository()

	saida := repo.UsuarioRepositoryList()
	esp := mocks.UsuariosBD

	assert.Equal(t, saida, esp)
}
