package repository

import (
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/mocks"
	"github.com/go-openapi/testify/v2/assert"
)

// teste de list
func TestUsuarioRepositoryList(t *testing.T) {
	// contrutor
	repo := NewUsuarioRepository()

	//saida
	esp := mocks.UsuariosBD

	// saida
	if saida := repo.UsuarioRepositoryList(); saida != nil {
		assert.Equal(t, saida, esp)
	}

}
