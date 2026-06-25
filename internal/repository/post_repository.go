package repository

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	Data *gorm.DB
}

func NewPostRepository(mp *gorm.DB) *PostRepository {
	return &PostRepository{Data: mp}
}

// listagem de post
func (p *PostRepository) PostRepositoryList() ([]models.Post, error) {
	var posts []models.Post
	err := p.Data.Find(&posts).Error
	return posts, err
}

// adicionar
func (p *PostRepository) PostRepositoryAdd(post *models.Post) error {
	return p.Data.Create(post).Error
}

// selecionar
func (p *PostRepository) PostRepositoryById(id int64) (*models.Post, error) {
	var post models.Post

	err := p.Data.First(&post, id).Error

	if err != nil {
		return nil, err
	}

	return &post, nil
}

// update
func (p *PostRepository) PostRepositoryUpdate(update *models.Post) error {
	return p.Data.Save(update).Error
}

// delete
func (p *PostRepository) PostRepositoryDelete(id int64) error {
	return p.Data.Delete(&models.Post{}, id).Error
}

// pesquisar por post/user
func (p *PostRepository) PostRepositorySearch(IDUser int64) error {
	var post models.Post

	return p.Data.Where(&models.Post{IDUser: IDUser}).First(&post).Error
}
