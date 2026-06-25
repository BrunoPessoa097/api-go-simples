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
	repoU := repository.NewUsuarioRepository(db)
	s := NewRoleService(repo, repoU)

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
	repoU := repository.NewUsuarioRepository(db)
	s := NewRoleService(repo, repoU)

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
	repoU := repository.NewUsuarioRepository(db)
	s := NewRoleService(repo, repoU)

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
	repoU := repository.NewUsuarioRepository(db)
	s := NewRoleService(repo, repoU)

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
	repoU := repository.NewUsuarioRepository(db)
	s := NewRoleService(repo, repoU)

	role1 := models.Roles{
		Nivel: "Governador1",
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
	saida := s.RoleServiceUpdate(role1.ID, role)
	// saida
	assert.NoError(t, saida)
}

// delete
func TestRoleServiceDelete(t *testing.T) {
	// inicializando
	db := utils.SetupDB(t)
	repo := repository.NewRolesRepository(db)
	repoU := repository.NewUsuarioRepository(db)
	s := NewRoleService(repo, repoU)

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
