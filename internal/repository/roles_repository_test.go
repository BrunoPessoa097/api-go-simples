package repository

import (
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/utils"
	"github.com/go-openapi/testify/v2/assert"
)

// listagem
func TestList(t *testing.T) {
	//iniciando
	db := utils.SetupDB(t)
	rp := NewRolesRepository(db)

	//recebendo
	saida, _ := rp.RolesRepositoryList()
	tam := len(saida)
	assert.Equal(t, tam, len(saida))
}

// busca
func TestSearch(t *testing.T) {
	db := utils.SetupDB(t)
	rp := NewRolesRepository(db)

	espc := "saida"
	saida := rp.RolesRepositoryBusca(espc)
	assert.Equal(t, false, saida)
}

// adicionar
func TestAdd(t *testing.T) {
	db := utils.SetupDB(t)
	rp := NewRolesRepository(db)

	espc := models.Roles{
		Nivel: "Adm2",
		Regra: "ler",
	}
	saida := rp.RolesRepositoryAdd(&espc)
	assert.Equal(t, nil, saida)
}

// byid
func TestID(t *testing.T) {
	db := utils.SetupDB(t)
	rp := NewRolesRepository(db)

	role := models.Roles{
		Nivel: "ADM",
		Regra: "ler",
	}

	err := db.Create(&role).Error
	if err != nil {
		t.Fatal(err)
	}

	saida, err := rp.RolesRepositoryById(role.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "ADM", saida.Nivel)
	assert.Equal(t, "ler", saida.Regra)
}

// update
func TestUpdate(t *testing.T) {
	db := utils.SetupDB(t)
	rp := NewRolesRepository(db)

	role := models.Roles{
		Nivel: "ADM",
		Regra: "ler",
	}

	err := db.Create(&role).Error
	if err != nil {
		t.Fatal(err)
	}

	espc := models.Roles{
		ID:    1,
		Nivel: "Adm 6",
		Regra: "ler",
	}
	saida := rp.RolesRepositoryUpdate(&espc)
	assert.Equal(t, nil, saida)
}

// delete
func TestDelete(t *testing.T) {
	db := utils.SetupDB(t)
	rp := NewRolesRepository(db)

	role := models.Roles{
		Nivel: "ADM",
		Regra: "ler",
	}

	err := db.Create(&role).Error
	if err != nil {
		t.Fatal(err)
	}

	saida := rp.RolesRepositoryDelete(int64(1))
	assert.Equal(t, nil, saida)
}
