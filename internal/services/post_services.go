package services

import (
	"errors"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
)

type PostService struct {
	Repo     *repository.PostRepository
	RepoUser *repository.UsuarioRepository
}

func NewPostService(r *repository.PostRepository, u *repository.UsuarioRepository) *PostService {
	return &PostService{
		Repo:     r,
		RepoUser: u,
	}
}

// listagem
func (p *PostService) PostServiceList() ([]models.Post, error) {
	return p.Repo.PostRepositoryList()
}

// adicionar
func (p *PostService) PostServiceAdd(post *models.Post) error {
	if _, err := p.RepoUser.UsuarioRepositoryById(int32(post.IDUser)); err != nil {
		return err
	}

	return p.Repo.PostRepositoryAdd(post)
}

// buscar por id
func (p *PostService) PostRepositoryId(id int64) (*models.Post, error) {
	return p.Repo.PostRepositoryById(id)
}

// atualizar post
func (p *PostService) PostServiceUpdate(id int64, post models.Post) error {
	post.ID = id
	if err := p.Repo.PostRepositoryUpdate(&post); err != nil {
		return errors.New("Problemas ao atualizar")
	}
	return nil
}

// delete
func (p *PostService) PostServiceDelete(id int64) error {
	if err := p.Repo.PostRepositoryDelete(id); err != nil {
		return errors.New("Erros ao deletar")
	}
	return nil
}
