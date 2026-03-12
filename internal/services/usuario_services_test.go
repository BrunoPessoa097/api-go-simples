package services

import (
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/go-openapi/testify/v2/assert"
)

func TestUsuarioServiceList(t *testing.T) {
	s := NewUsuarioService()

	saida := s.UsuarioServiceList()
	espec := make([]models.UsuarioCriate, 0)

	assert.Equal(t, espec, saida)
}

func TestUsuarioServiceAdd(t *testing.T) {
	s := NewUsuarioService()

	saida := s.UsuarioServiceAdd()
	espec := true

	assert.Equal(t, espec, saida)
}

func TestUsuarioServiceById(t *testing.T) {
	s := NewUsuarioService()

	saida := s.UsuarioServiceById(1)
	espec := models.UsuarioCriate{
		Id:        0,
		Nome:      "Bruno",
		Email:     "brunopessoa@gmail.com",
		Senha:     "1234",
		Role:      1,
		Bloqueado: false,
	}

	assert.Equal(t, saida, espec)
}
