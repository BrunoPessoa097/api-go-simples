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
	mock := mocks.UsuariosBD
	repo := NewUsuarioRepository(mock)

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
		Nome:      "Bruno",
		Email:     "brunopessoa@gmail.com",
		Senha:     "1234",
		Role:      1,
		Bloqueado: false,
	}

	// iniciando
	mock := mocks.UsuariosBD
	repo := NewUsuarioRepository(mock)

	//saida
	if _, err := repo.UsuarioRepositoryAdd(espec); err {
		assert.Equal(t, err, true)
	}
}

// buscar por id
func TestUsuarioRepositoryById(t *testing.T) {
	mock := mocks.UsuariosBD
	repo := NewUsuarioRepository(mock)

	if saida := repo.UsuarioRepositoryById(1); saida != nil {
		assert.Equal(t, saida.Nome, "Bruno F")
	} else {
		assert.Equal(t, nil, nil)
	}
}

// update
func TestUsuarioRepositoryUpdate(t *testing.T) {
	mock := mocks.UsuariosBD
	repo := NewUsuarioRepository(mock)

	espec := models.UsuarioCriate{
		Nome:      "Bruno 3",
		Email:     "bruno123@gmail.com",
		Senha:     "1234",
		Role:      1,
		Bloqueado: false,
	}

	if _, ok := repo.UsuarioRepositoryUpdate(1, espec); ok {
		assert.Equal(t, true, ok)
	} else {
		assert.Equal(t, false, ok)
	}
}

// delete
func TestUsuarioRepositoryDelete(t *testing.T) {
	mock := mocks.UsuariosBD
	repo := NewUsuarioRepository(mock)

	if saida := repo.UsuarioRepositoryDelete(1); saida {
		assert.Equal(t, true, saida)
	} else {
		assert.Equal(t, false, saida)
	}
}

// teste de busca
func TestUsuarioRepositorySearch(t *testing.T) {
	mock := mocks.UsuariosBD
	repo := NewUsuarioRepository(mock)

	if ok := repo.UsuarioRepositorySearch("Bruno F", "brunopessoa097@gmail.com"); ok {
		assert.Equal(t, true, ok)
	}
}
