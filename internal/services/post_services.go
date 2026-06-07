package services

import (
	"errors"
	"time"

	"github.com/BrunoPessoa097/api-go-simples/internal/mocks"
	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
)

type PostService struct {
	Repo *repository.PostRepository
}

func NewPostService(r *repository.PostRepository) *PostService {
	return &PostService{
		Repo: r,
	}
}

// listagem
func (p *PostService) PostServiceList() []models.Post {
	return p.Repo.Data
}

// adicionar
func (p *PostService) PostServiceAdd(post models.Post) bool {
	tam := len(mocks.ListPost) - 1
	id := mocks.ListPost[tam].ID + 1

	posts := models.Post{
		ID:       id,
		IDUser:   post.IDUser,
		Texto:    post.Texto,
		DtCreate: time.Now(),
		DtUpdate: time.Now(),
	}
	return p.Repo.PostRepositoryAdd(posts)
}

func (p *PostService) PostRepositoryId(id int64) (*models.Post, error) {
	return p.Repo.PostRepositoryById(id)
}

func (p *PostService) PostServiceUpdate(id int64, post models.Post) error {
	if ok := p.Repo.PostRepositoryUpdate(id, post); ok {
		return nil
	}
	return errors.New("Problemas ao atualizar")
}
