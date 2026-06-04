package repository

import (
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/mocks"
	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/go-openapi/testify/v2/assert"
)

func Init() *PostRepository {
	mock := mocks.ListPost
	return NewPostRepository(mock)
}

func TestListPost(t *testing.T) {
	post := Init()
	saida := post.PostRepositoryList()

	assert.Equal(t, saida[0].Texto, "Como é viver no Brasil?")
}

func TestAddPost(t *testing.T) {
	r := Init()
	post := models.Post{
		IDUser: 1,
		Texto:  "Vamos nessa Brasil",
	}
	saida := r.PostRepositoryAdd(post)
	assert.Equal(t, true, saida)
}
