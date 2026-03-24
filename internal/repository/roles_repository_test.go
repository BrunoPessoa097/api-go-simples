package repository

import (
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/mocks"
	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/go-openapi/testify/v2/assert"
)

// listagem
func TestList(t *testing.T) {
	//iniciando
	mc := mocks.ListRoles
	rp := NewRolesRepository(mc)

	//recebendo
	saida := rp.RolesRepositoryList()
	espc := mc[0].Nivel

	//saido
	assert.Equal(t, espc, saida[0].Nivel)
}

// busca
func TestSearch(t *testing.T) {
	mc := mocks.ListRoles
	rp := NewRolesRepository(mc)

	espc := mc[0].Nivel
	saida := rp.RolesRepositoryBusca(espc)

	assert.Equal(t, true, saida)
}

// adicionar
func TestAdd(t *testing.T) {
	mc := mocks.ListRoles
	rp := NewRolesRepository(mc)

	espc := models.Roles{
		Nivel: "Adm2",
		Regra: "ler",
	}
	saida := rp.RolesRepositoryAdd(&espc)

	assert.Equal(t, true, saida)
}

// // byid
func TestID(t *testing.T) {
	mc := mocks.ListRoles
	rp := NewRolesRepository(mc)

	saida := rp.RolesRepositoryById(1)

	assert.Equal(t, "ADM", saida.Nivel)
}

// update
func TestUpdate(t *testing.T) {
	mc := mocks.ListRoles
	rp := NewRolesRepository(mc)

	espc := models.Roles{
		ID:    1,
		Nivel: "Adm 6",
		Regra: "ler",
	}
	saida := rp.RolesRepositoryUpdate(1, espc)

	assert.Equal(t, true, saida)
}

// delete
func TestDelete(t *testing.T) {
	mc := mocks.ListRoles
	rp := NewRolesRepository(mc)

	saida := rp.RolesRepositoryDelete(int64(2))

	assert.Equal(t, true, saida)
}
