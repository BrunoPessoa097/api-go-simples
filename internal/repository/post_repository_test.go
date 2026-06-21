package repository

import (
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/utils"
	"github.com/go-openapi/testify/v2/assert"
)

// listagem
func TestListPost(t *testing.T) {
	db := utils.SetupDB(t)
	post := NewPostRepository(db)
	saida, _ := post.PostRepositoryList()
	tam := make([]models.Post, len(saida))

	assert.Equal(t, tam, saida)
}

// adicionar
func TestAddPost(t *testing.T) {
	db := utils.SetupDB(t)
	post := NewPostRepository(db)

	//postagem
	postagem := models.Post{
		IDUser: 1,
		Texto:  "Vamos nessa Brasil",
	}

	saida := post.PostRepositoryAdd(postagem)

	assert.Equal(t, nil, saida)
}

// id
func TestIdPost(t *testing.T) {
	db := utils.SetupDB(t)
	post := NewPostRepository(db)

	//criar postagem
	postagem := models.Post{
		IDUser: 1,
		Texto:  "Vamos nessa Brasil",
	}

	//criando postagem
	err := db.Create(&postagem).Error
	if err != nil {
		t.Fatal(err)
	}

	//
	saida, _ := post.PostRepositoryById(postagem.ID)

	assert.Equal(t, postagem.Texto, saida.Texto)
}

func TestUpdatePost(t *testing.T) {
	db := utils.SetupDB(t)
	post := NewPostRepository(db)

	postagem := models.Post{
		IDUser: 1,
		Texto:  "Vamos nessa Brasil",
	}

	err := db.Create(&postagem).Error
	if err != nil {
		t.Fatal(err)
	}

	text := models.Post{ID: postagem.ID, IDUser: 1, Texto: "vamos falar da Itália"}

	saida := post.PostRepositoryUpdate(text)

	assert.Equal(t, nil, saida)
}

func TestDeletePost(t *testing.T) {
	db := utils.SetupDB(t)
	post := NewPostRepository(db)

	postagem := models.Post{
		IDUser: 1,
		Texto:  "Vamos nessa Brasil",
	}

	err := db.Create(&postagem).Error
	if err != nil {
		t.Fatal(err)
	}

	saida := post.PostRepositoryDelete(postagem.ID)

	assert.Equal(t, nil, saida)
}
