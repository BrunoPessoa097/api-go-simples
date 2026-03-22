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
		Nome:      "Bruno Pess",
		Email:     "ps1@mail.com",
		Senha:     "12345678",
		Role:      1,
		Bloqueado: false,
	}

	// iniciando
	mock := mocks.UsuariosBD
	repo := NewUsuarioRepository(mock)

	//saida
	err := repo.UsuarioRepositoryAdd(espec)
	assert.Equal(t, nil, err)
}

// buscar por id
func TestUsuarioRepositoryById(t *testing.T) {
	//iniciando
	mock := mocks.UsuariosBD
	repo := NewUsuarioRepository(mock)

	//verificar
	saida := repo.UsuarioRepositoryById(1)
	assert.Equal(t, saida.Nome, "Bruno F")
}

// // update
func TestUsuarioRepositoryUpdate(t *testing.T) {
	//iniciando
	mock := mocks.UsuariosBD
	repo := NewUsuarioRepository(mock)

	//entradas
	espec := models.UsuarioCriate{
		Id:        1,
		Nome:      "Brubru",
		Email:     "bp@gmail.com",
		Senha:     "12345678",
		Role:      1,
		Bloqueado: false,
	}

	// saida
	ok := repo.UsuarioRepositoryUpdate(1, espec)
	assert.Equal(t, nil, ok)
}

// delete
func TestUsuarioRepositoryDelete(t *testing.T) {
	//iniciando
	mock := mocks.UsuariosBD
	repo := NewUsuarioRepository(mock)

	//saida
	saida := repo.UsuarioRepositoryDelete(1)
	assert.Equal(t, nil, saida)
}

// teste de busca
func TestUsuarioRepositorySearch(t *testing.T) {
	//iniciar
	mock := mocks.UsuariosBD
	repo := NewUsuarioRepository(mock)

	//saida
	ok := repo.UsuarioRepositorySearch("Bruno F", "brunopessoa@gmail.com")
	assert.Equal(t, true, ok)
}
