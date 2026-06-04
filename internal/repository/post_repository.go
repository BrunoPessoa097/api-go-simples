package repository

import (
	"fmt"

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

// adicionar
func (p *PostRepository) PostRepositoryAdd(post models.Post) bool {
	p.Data = append(p.Data, post)
	return true
}

// selecionar
func (p *PostRepository) PostRepositoryById(id int64) (*models.Post, error) {
	for i := range p.Data {
		if p.Data[i].ID == id {
			return &p.Data[i], nil
		}
	}
	return nil, fmt.Errorf("post não encontrado")
}

// update
func (p *PostRepository) PostRepositoryUpdate(id int64, update models.Post) bool {
	for i := range p.Data {
		if p.Data[i].ID == id {
			p.Data[i] = update
			return true
		}
	}
	return false
}
