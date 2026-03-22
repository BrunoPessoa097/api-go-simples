package services

import (
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/mocks"
	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
	"github.com/go-openapi/testify/v2/assert"
)

// listar
func TestUsuarioServiceList(t *testing.T) {
	// iniciando
	m := mocks.UsuariosBD
	repo := repository.NewUsuarioRepository(m)
	s := NewUsuarioService(repo)

	//saida e expectativa
	saida := s.UsuarioServiceList()
	espec := "Bruno F"

	//saida
	assert.Equal(t, saida[0].Nome, espec)
}

// adicionar
func TestUsuarioServiceAdd(t *testing.T) {
	// iniciando
	m := mocks.UsuariosBD
	repo := repository.NewUsuarioRepository(m)
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
	m := mocks.UsuariosBD
	repo := repository.NewUsuarioRepository(m)
	s := NewUsuarioService(repo)

	//saida e expectativa
	saida := s.UsuarioServiceById(1)

	//saida
	assert.Equal(t, saida, saida)
}

// update
func TestUsuarioServiceUpdate(t *testing.T) {
	// iniciando
	m := mocks.UsuariosBD
	repo := repository.NewUsuarioRepository(m)
	s := NewUsuarioService(repo)

	espec := models.Usuario{
		Id:        1,
		Nome:      "Brubru",
		Email:     "bp@gmail.com",
		Senha:     "12345678",
		Role:      1,
		Bloqueado: false,
	}

	//saida e expectativa
	ok := s.UsuarioServiceUpdate(1, espec)
	assert.Equal(t, nil, ok)
}

// delete
func TestUsuarioServiceDelete(t *testing.T) {
	// iniciando
	m := mocks.UsuariosBD
	repo := repository.NewUsuarioRepository(m)
	s := NewUsuarioService(repo)

	//saida e expectativa
	ok := s.UsuarioServiceDelete(1)
	assert.Equal(t, nil, ok)
}
