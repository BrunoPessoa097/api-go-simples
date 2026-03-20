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

// busca
func TestRoleServiceSearch(t *testing.T) {
	//iniciando
	s := NewRoleService()

	// buscando algo
	saida := s.RoleServiceSearch("ADM")

	//saida
	assert.Equal(t, true, saida)
}

// adicao
func TestRoleServicePost(t *testing.T) {
	//iniciar
	s := NewRoleService()

	//modelo
	role := models.RolesC{
		Nivel: "Governador",
		Regra: "ler",
	}

	//saida
	saida, _ := s.RoleServicePost(role)

	// saida
	assert.Equal(t, true, saida)
}

// busca por id
func TestRoleServiceById(t *testing.T) {
	//iniciando
	s := NewRoleService()

	// expectativa
	esp := "ADM"
	saida, _ := s.RoleServiceById(1)

	//saido
	assert.Equal(t, esp, saida.Nivel)
}

// delete
func TestRoleServiceDelete(t *testing.T) {
	// inicializando
	s := NewRoleService()

	//queries
	esp := true
	saida := s.RoleServiceDelete(1)

	//saida
	assert.Equal(t, esp, saida)
}
