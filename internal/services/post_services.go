package services

import (
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

func (p *PostService) PostServiceList() []models.Post {
	return p.Repo.Data
}
