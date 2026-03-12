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
