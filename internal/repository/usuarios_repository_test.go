package repository

import (
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/utils"
	"github.com/go-openapi/testify/v2/assert"
)

// teste de list
func TestUsuarioRepositoryList(t *testing.T) {
	// contrutor
	db := utils.SetupDB(t)
	repo := NewUsuarioRepository(db)

	// saida
	saida, _ := repo.UsuarioRepositoryList()
	tam := len(saida)
	assert.Equal(t, tam, len(saida))
}

// adicionar
func TestUsuarioRepositoryAdd(t *testing.T) {
	//base

	// iniciando
	db := utils.SetupDB(t)
	repo := NewUsuarioRepository(db)

	espec := models.Usuario{
		Nome:      "Bruno Pess",
		Email:     "ps1@mail.com",
		Senha:     "12345678",
		Role:      1,
		Bloqueado: false,
	}
	//saida
	err := repo.UsuarioRepositoryAdd(&espec)

	assert.Equal(t, nil, err)
}

// buscar por id
func TestUsuarioRepositoryById(t *testing.T) {
	//iniciando
	db := utils.SetupDB(t)
	repo := NewUsuarioRepository(db)

	espec := models.Usuario{
		Nome:      "Bruno Pessoa",
		Email:     "ps1@mail.com",
		Senha:     "12345678",
		Role:      1,
		Bloqueado: false,
	}

	err := db.Create(&espec).Error
	if err != nil {
		t.Fatal(err)
	}

	//verificar
	saida, err := repo.UsuarioRepositoryById(espec.ID)

	assert.Equal(t, "Bruno Pessoa", saida.Nome)
}

// update
func TestUsuarioRepositoryUpdate(t *testing.T) {
	//iniciando
	db := utils.SetupDB(t)
	repo := NewUsuarioRepository(db)

	//entradas
	espec := models.Usuario{
		ID:        1,
		Nome:      "Brubru",
		Email:     "bp@gmail.com",
		Senha:     "12345678",
		Role:      1,
		Bloqueado: false,
	}

	criar := models.Usuario{
		Nome:      "Bruno Pessoa",
		Email:     "ps1@mail.com",
		Senha:     "12345678",
		Role:      1,
		Bloqueado: false,
	}

	err := db.Create(&criar).Error
	if err != nil {
		t.Fatal(err)
	}

	// saida
	ok := repo.UsuarioRepositoryUpdate(espec)
	assert.Equal(t, nil, ok)
}

// delete
func TestUsuarioRepositoryDelete(t *testing.T) {
	//iniciando
	db := utils.SetupDB(t)
	repo := NewUsuarioRepository(db)
	criar := models.Usuario{
		Nome:      "Bruno Pessoa",
		Email:     "ps1@mail.com",
		Senha:     "12345678",
		Role:      1,
		Bloqueado: false,
	}

	err := db.Create(&criar).Error
	if err != nil {
		t.Fatal(err)
	}

	//saida
	saida := repo.UsuarioRepositoryDelete(1)
	assert.Equal(t, nil, saida)
}

// teste de busca
func TestUsuarioRepositorySearch(t *testing.T) {
	db := utils.SetupDB(t)
	repo := NewUsuarioRepository(db)

	criar := models.Usuario{
		Nome:      "Bruno Pessoa",
		Email:     "ps1@mail.com",
		Senha:     "12345678",
		Role:      1,
		Bloqueado: false,
	}

	err := db.Create(&criar).Error
	if err != nil {
		t.Fatal(err)
	}

	//saida
	ok := repo.UsuarioRepositorySearch("Bruno Pessoa", "ps1@mail.com")
	assert.Equal(t, true, ok)
}
