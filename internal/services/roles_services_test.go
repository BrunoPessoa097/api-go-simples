package services

import (
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
	"github.com/BrunoPessoa097/api-go-simples/internal/utils"
	"github.com/go-openapi/testify/v2/assert"
)

// listagem
func TestRoleServiceList(t *testing.T) {
	//iniciando
	db := utils.SetupDB(t)
	repo := repository.NewRolesRepository(db)
	s := NewRoleService(repo)

	//listagem
	saida := s.RoleServiceList()

	espec := make([]models.Roles, len(saida))

	//verificação
	assert.Equal(t, espec, saida)
}

// busca
func TestRoleServiceSearch(t *testing.T) {
	//iniciando
	db := utils.SetupDB(t)
	repo := repository.NewRolesRepository(db)
	s := NewRoleService(repo)

	// buscando algo
	saida := s.RoleServiceSearch("ADM")

	//saida
	assert.Equal(t, nil, saida)
}

// adicao
func TestRoleServicePost(t *testing.T) {
	//iniciar
	db := utils.SetupDB(t)
	repo := repository.NewRolesRepository(db)
	s := NewRoleService(repo)

	//modelo
	role := models.Roles{
		Nivel: "Governador",
		Regra: "ler",
	}

	//saida
	saida := s.RoleServicePost(role)

	// saida
	assert.Equal(t, nil, saida)
}

// busca por id
func TestRoleServiceById(t *testing.T) {
	//iniciando
	db := utils.SetupDB(t)
	repo := repository.NewRolesRepository(db)
	s := NewRoleService(repo)

	role := models.Roles{
		Nivel: "Governador",
		Regra: "ler",
	}

	s.RoleServicePost(role)

	// expectativa
	saida, _ := s.RoleServiceById(1)

	//saido
	assert.Equal(t, role.Nivel, saida.Nivel)
}

// update
func TestRoleServiceUpdate(t *testing.T) {
	//iniciar
	db := utils.SetupDB(t)
	repo := repository.NewRolesRepository(db)
	s := NewRoleService(repo)

	role1 := models.Roles{
		Nivel: "Governador",
		Regra: "ler",
	}

	s.RoleServicePost(role1)

	//modelo
	role := models.Roles{
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
	db := utils.SetupDB(t)
	repo := repository.NewRolesRepository(db)
	s := NewRoleService(repo)

	role1 := models.Roles{
		Nivel: "Governador",
		Regra: "ler",
	}

	s.RoleServicePost(role1)

	//queries
	saida := s.RoleServiceDelete(1)
	//saida
	assert.Equal(t, nil, saida)
}
