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

func (p *PostRepository) PostRepositoryAdd(post models.Post) bool {
	p.Data = append(p.Data, post)
	return true
}
