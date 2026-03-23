package services

import (
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/mocks"
	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
	"github.com/go-openapi/testify/v2/assert"
)

// listagem
func TestRoleServiceList(t *testing.T) {
	//iniciando
	mc := mocks.ListRoles
	repo := repository.NewRolesRepository(mc)
	s := NewRoleService(repo)

	//listagem
	saida := s.RoleServiceList()

	//verificação
	assert.Equal(t, "ADM", saida[0].Nivel)
}

// busca
func TestRoleServiceSearch(t *testing.T) {
	//iniciando
	mo := mocks.ListRoles
	rp := repository.NewRolesRepository(mo)
	s := NewRoleService(rp)

	// buscando algo
	saida := s.RoleServiceSearch("ADM")

	//saida
	assert.Equal(t, "Nivel já cadastrado", saida.Error())
}

// adicao
func TestRoleServicePost(t *testing.T) {
	//iniciar
	mocks := mocks.ListRoles
	repo := repository.NewRolesRepository(mocks)
	s := NewRoleService(repo)

	//modelo
	role := models.RolesC{
		Nivel: "Governador",
		Regra: "ler",
	}

	//saida
	saida := s.RoleServicePost(role)

	// saida
	assert.Equal(t, true, saida)
}

// busca por id
func TestRoleServiceById(t *testing.T) {
	//iniciando
	mocks := mocks.ListRoles
	repo := repository.NewRolesRepository(mocks)
	s := NewRoleService(repo)

	// expectativa
	esp := "ADM"
	saida := s.RoleServiceById(1)

	//saido
	assert.Equal(t, esp, saida.Nivel)
}

// update
func TestRoleServiceUpdate(t *testing.T) {
	//iniciar
	mocks := mocks.ListRoles
	repo := repository.NewRolesRepository(mocks)
	s := NewRoleService(repo)

	//modelo
	role := models.RolesC{
		ID:    1,
		Nivel: "Gov",
		Regra: "ler",
	}

	//saida
	saida := s.RoleServiceUpdate(1, role)

	// saida
	assert.Equal(t, nil, saida)
}

// delete
func TestRoleServiceDelete(t *testing.T) {
	// inicializando
	mocks := mocks.ListRoles
	repo := repository.NewRolesRepository(mocks)
	s := NewRoleService(repo)

	//queries
	saida := s.RoleServiceDelete(1)

	//saida
	assert.Equal(t, nil, saida)
}
