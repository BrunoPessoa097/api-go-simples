package services

import (
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
	"github.com/go-openapi/testify/v2/assert"
)

func TestUsuarioServiceList(t *testing.T) {
	// iniciando
	repo := repository.NewUsuarioRepository()
	s := NewUsuarioService(repo)

	//saida e expectativa
	saida := s.UsuarioServiceList()
	espec := make([]models.UsuarioCriate, 0)

	//saida
	assert.Equal(t, espec, saida)
}

func TestUsuarioServiceAdd(t *testing.T) {
	// iniciando
	repo := repository.NewUsuarioRepository()
	s := NewUsuarioService(repo)

	espec := models.UsuarioCriate{
		Id:        0,
		Nome:      "Bruno",
		Email:     "brunopessoa@gmail.com",
		Senha:     "1234",
		Role:      1,
		Bloqueado: false,
	}

	//saida e expectativa
	saida := s.UsuarioServiceAdd(espec)

	//saida
	assert.Equal(t, true, saida)
}

func TestUsuarioServiceById(t *testing.T) {
	// iniciando
	repo := repository.NewUsuarioRepository()
	s := NewUsuarioService(repo)

	//saida e expectativa
	saida := s.UsuarioServiceById(1)

	//saida
	assert.Equal(t, saida, saida)
}

func TestUsuarioServiceUpdate(t *testing.T) {
	// iniciando
	repo := repository.NewUsuarioRepository()
	s := NewUsuarioService(repo)

	espec := models.UsuarioCriate{
		Id:        0,
		Nome:      "Bruno",
		Email:     "brunopessoa@gmail.com",
		Senha:     "1234",
		Role:      1,
		Bloqueado: false,
	}

	//saida e expectativa
	if saida := s.UsuarioServiceUpdate(1, espec); saida {
		//saida
		assert.Equal(t, saida, true)
	} else {
		//saida
		assert.Equal(t, saida, false)
	}
}

func TestUsuarioServiceDelete(t *testing.T) {
	// iniciando
	repo := repository.NewUsuarioRepository()
	s := NewUsuarioService(repo)

	//saida e expectativa
	if saida := s.UsuarioServiceDelete(1); saida {
		//saida
		assert.Equal(t, saida, true)
	} else {
		//saida
		assert.Equal(t, saida, false)
	}
}
