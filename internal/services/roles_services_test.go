package services

import (
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/go-openapi/testify/v2/assert"
)

// listagem
func TestRoleServiceList(t *testing.T) {
	//iniciando
	s := NewRoleService()

	//listagem
	saida := s.RoleServiceList()

	//verificação
	assert.Equal(t, "ADM", saida[0].Nivel)
}

func TestRoleServiceSearch(t *testing.T) {
	//iniciando
	s := NewRoleService()

	// buscando algo
	saida := s.RoleServiceSearch("ADM")

	//saida
	assert.Equal(t, true, saida)
}

func TestRoleServicePost(t *testing.T) {
	s := NewRoleService()

	role := models.RolesC{
		Nivel: "Governador",
		Regra: "ler",
	}

	saida, _ := s.RoleServicePost(role)

	assert.Equal(t, true, saida)
}
