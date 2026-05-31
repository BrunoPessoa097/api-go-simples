package repository

import (
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/mocks"
	"github.com/go-openapi/testify/v2/assert"
)

func TestListPost(t *testing.T) {
	mock := mocks.ListPost
	post := NewPostRepository(mock)

	saida := post.PostRepositoryList()

	assert.Equal(t, saida[0].Texto, "Como é viver no Brasil?")
}
