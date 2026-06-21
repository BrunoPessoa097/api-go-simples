package services

import (
	"fmt"
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
	"github.com/BrunoPessoa097/api-go-simples/internal/utils"
	"github.com/go-openapi/testify/v2/assert"
)

// listar
func TestPostList(t *testing.T) {
	db := utils.SetupDB(t)
	ruser := repository.NewUsuarioRepository(db)
	rpost := repository.NewPostRepository(db)
	spost := NewPostService(rpost, ruser)

	saida, _ := spost.PostServiceList()
	tam := make([]models.Post, len(saida))

	assert.Equal(t, tam, saida)
}

// adicionar
func TestPostAdd(t *testing.T) {
	db := utils.SetupDB(t)

	ruser := repository.NewUsuarioRepository(db)
	rpost := repository.NewPostRepository(db)
	spost := NewPostService(rpost, ruser)

	user := models.Usuario{
		Nome:  "Bruno",
		Email: "bruno@test.com",
	}

	ruser.UsuarioRepositoryAdd(&user)

	post := models.Post{
		IDUser: int64(user.ID),
		Texto:  "Vamos nessa Brasil",
	}

	saida := spost.PostServiceAdd(&post)

	assert.Equal(t, nil, saida)
}

// buscar por id
func TestPostId(t *testing.T) {
	db := utils.SetupDB(t)
	ruser := repository.NewUsuarioRepository(db)
	rpost := repository.NewPostRepository(db)
	spost := NewPostService(rpost, ruser)

	user := models.Usuario{
		Nome:  "Bruno",
		Email: "bruno@test.com",
	}

	ruser.UsuarioRepositoryAdd(&user)

	post := models.Post{
		IDUser: int64(user.ID),
		Texto:  "Vamos nessa Brasil",
	}
	spost.PostServiceAdd(&post)

	saida, _ := spost.PostRepositoryId(post.ID)

	assert.Equal(t, post.Texto, saida.Texto)
}

// update
func TestPostUpdate(t *testing.T) {
	db := utils.SetupDB(t)
	ruser := repository.NewUsuarioRepository(db)
	rpost := repository.NewPostRepository(db)
	spost := NewPostService(rpost, ruser)

	user := models.Usuario{
		Nome:  "Bruno",
		Email: "bruno@test.com",
	}

	ruser.UsuarioRepositoryAdd(&user)

	post := models.Post{
		IDUser: int64(user.ID),
		Texto:  "Vamos nessa Brasil",
	}

	spost.PostServiceAdd(&post)

	postUpdate := models.Post{
		ID:     post.ID,
		IDUser: int64(user.ID),
		Texto:  "Como é viver no Jamaica?",
	}

	saida := spost.PostServiceUpdate(post.ID, postUpdate)

	assert.Equal(t, nil, saida)
}

// delete
func TestPostDelete(t *testing.T) {
	db := utils.SetupDB(t)
	ruser := repository.NewUsuarioRepository(db)
	rpost := repository.NewPostRepository(db)
	spost := NewPostService(rpost, ruser)

	user := models.Usuario{
		Nome:  "Bruno",
		Email: "bruno@test.com",
	}

	ruser.UsuarioRepositoryAdd(&user)

	post := models.Post{
		IDUser: int64(user.ID),
		Texto:  "Vamos nessa Brasil",
	}

	spost.PostServiceAdd(&post)
	fmt.Println(post.ID)

	saida := spost.PostServiceDelete(post.ID)

	assert.Equal(t, nil, saida)
}
