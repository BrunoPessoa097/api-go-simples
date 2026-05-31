package repository

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/models"
)

type PostRepository struct {
	Data []models.Post
}

func NewPostRepository(mp []models.Post) *PostRepository {
	return &PostRepository{Data: mp}
}

// listagem de post
func (p *PostRepository) PostRepositoryList() []models.Post {
	return p.Data
}
