package services

import (
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
	"github.com/BrunoPessoa097/api-go-simples/internal/utils"
	"github.com/go-openapi/testify/v2/assert"
)

// listar
func TestUsuarioServiceList(t *testing.T) {
	// iniciando
	db := utils.SetupDB(t)
	repo := repository.NewUsuarioRepository(db)
	s := NewUsuarioService(repo)

	//saida e expectativa
	saida, _ := s.UsuarioServiceList()
	espec := make([]models.Usuario, len(saida))

	//saida
	assert.Equal(t, espec, saida)
}

// adicionar
func TestUsuarioServiceAdd(t *testing.T) {
	// iniciando
	db := utils.SetupDB(t)
	repo := repository.NewUsuarioRepository(db)
	s := NewUsuarioService(repo)

	espec := models.Usuario{
		Nome:      "Bruno Pess",
		Email:     "ps1@mail.com",
		Senha:     "12345678",
		Role:      1,
		Bloqueado: false,
	}

	//saida e expectativa
	err := s.UsuarioServiceAdd(espec)

	//saida
	assert.Equal(t, nil, err)
}

// byid
func TestUsuarioServiceById(t *testing.T) {
	// iniciando
	db := utils.SetupDB(t)
	repo := repository.NewUsuarioRepository(db)
	s := NewUsuarioService(repo)

	espec := models.Usuario{
		Nome:      "Bruno Pess",
		Email:     "ps1@mail.com",
		Senha:     "12345678",
		Role:      1,
		Bloqueado: false,
	}

	s.repo.UsuarioRepositoryAdd(espec)

	//saida e expectativa
	saida, _ := s.UsuarioServiceById(1)

	//saida
	assert.Equal(t, espec.Nome, saida.Nome)
}

// update
func TestUsuarioServiceUpdate(t *testing.T) {
	// iniciando
	db := utils.SetupDB(t)
	repo := repository.NewUsuarioRepository(db)
	s := NewUsuarioService(repo)

	espec := models.Usuario{
		ID:        1,
		Nome:      "Brubru",
		Email:     "bp@gmail.com",
		Senha:     "12345678",
		Role:      1,
		Bloqueado: false,
	}

	s.repo.UsuarioRepositoryAdd(espec)

	//saida e expectativa
	ok := s.UsuarioServiceUpdate(1, espec)

	assert.Equal(t, nil, ok)
}

// delete
func TestUsuarioServiceDelete(t *testing.T) {
	// iniciando
	db := utils.SetupDB(t)
	repo := repository.NewUsuarioRepository(db)
	s := NewUsuarioService(repo)

	espec := models.Usuario{
		ID:        1,
		Nome:      "Brubru",
		Email:     "bp@gmail.com",
		Senha:     "12345678",
		Role:      1,
		Bloqueado: false,
	}

	s.repo.UsuarioRepositoryAdd(espec)

	//saida e expectativa
	ok := s.UsuarioServiceDelete(1)
	assert.Equal(t, nil, ok)
}
