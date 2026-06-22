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
	if saida, _ := p.RepoUser.UsuarioRepositoryById(int32(post.IDUser)); saida != nil {
		return p.Repo.PostRepositoryAdd(post)
	}

	return errors.New("usuario não encontrado")
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
	if _, err := p.Repo.PostRepositoryById(id); err != nil {
		return errors.New("Post não encontrado")
	}
	if err := p.Repo.PostRepositoryDelete(id); err != nil {
		return errors.New("Erro ao deletar")
	}
	return nil
}
