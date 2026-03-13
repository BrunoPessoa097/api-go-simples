package repository

import (
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/mocks"
	"github.com/BrunoPessoa097/api-go-simples/internal/models"
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

// adicionar
func TestUsuarioRepositoryAdd(t *testing.T) {
	//base
	espec := models.UsuarioCriate{
		Id:        0,
		Nome:      "Bruno",
		Email:     "brunopessoa@gmail.com",
		Senha:     "1234",
		Role:      1,
		Bloqueado: false,
	}

	// iniciando
	repo := NewUsuarioRepository()

	//saida
	if saida := repo.UsuarioRepositoryAdd(espec); saida {
		assert.Equal(t, saida, true)
	}
}

// buscar por id
func TestUsuarioRepositoryById(t *testing.T) {
	repo := NewUsuarioRepository()

	if saida := repo.UsuarioRepositoryById(1); saida != nil {
		assert.Equal(t, saida.Nome, "bruno")
	} else {
		assert.Equal(t, nil, nil)
	}
}

// update
func TestUsuarioRepositoryUpdate(t *testing.T) {
	repo := NewUsuarioRepository()

	espec := models.UsuarioCriate{
		Id:        0,
		Nome:      "Bruno",
		Email:     "brunopessoa@gmail.com",
		Senha:     "1234",
		Role:      1,
		Bloqueado: false,
	}

	if saida := repo.UsuarioRepositoryUpdate(1, espec); saida {
		assert.Equal(t, false, saida)
	}
}

// delete
func TestUsuarioRepositoryDelete(t *testing.T) {
	repo := NewUsuarioRepository()

	if saida := repo.UsuarioRepositoryDelete(1); saida {
		assert.Equal(t, true, saida)
	} else {
		assert.Equal(t, false, saida)
	}
}
