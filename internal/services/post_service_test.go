package services

import (
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/mocks"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
	"github.com/go-openapi/testify/v2/assert"
)

func TestPostList(t *testing.T) {
	mocks := mocks.ListPost
	rpost := repository.NewPostRepository(mocks)
	spost := NewPostService(rpost)

	saida := spost.PostServiceList()
	assert.Equal(t, saida[0].Texto, "Como é viver no Brasil?")
}
