package services

import (
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/mocks"
	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
	"github.com/go-openapi/testify/v2/assert"
)

func iniciarPs() *PostService {
	mocks := mocks.ListPost
	rpost := repository.NewPostRepository(mocks)
	return NewPostService(rpost)
}

// listar
func TestPostList(t *testing.T) {
	spost := iniciarPs()

	saida := spost.PostServiceList()
	assert.Equal(t, saida[0].Texto, "Como é viver no Brasil?")
}

// adicionar
func TestPostAdd(t *testing.T) {
	spost := iniciarPs()
	post := models.Post{
		IDUser: 1,
		Texto:  "Vamos nessa Brasil",
	}

	saida := spost.PostServiceAdd(post)
	assert.Equal(t, true, saida)
}

// buscar por id
func TestPostId(t *testing.T) {
	spost := iniciarPs()
	id := int64(1)

	saida, _ := spost.PostRepositoryId(id)
	assert.Equal(t, "Como é viver no Brasil?", saida.Texto)
}

func TestPostUpdate(t *testing.T) {
	spost := iniciarPs()
	id := int64(1)
	post := models.Post{
		ID:     1,
		IDUser: 1,
		Texto:  "Como é viver no Jamaica?",
	}

	saida := spost.PostServiceUpdate(id, post)
	assert.Equal(t, nil, saida)
}
