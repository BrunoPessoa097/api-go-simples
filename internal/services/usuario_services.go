package services

import (
	"errors"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
)

type UsuarioService struct {
	repo     *repository.UsuarioRepository
	roles    *repository.RolesRepository
	repoPost *repository.PostRepository
}

// construtor
func NewUsuarioService(repo *repository.UsuarioRepository, roles *repository.RolesRepository, post *repository.PostRepository) *UsuarioService {
	return &UsuarioService{
		repo:     repo,
		roles:    roles,
		repoPost: post,
	}
}

// listar
func (s *UsuarioService) UsuarioServiceList() ([]models.Usuario, error) {
	return s.repo.UsuarioRepositoryList()
}

// adiciona
func (s *UsuarioService) UsuarioServiceAdd(user models.Usuario) error {
	if _, verifi := s.repo.UsuarioRepositorySearch(user.Nome, "", 0); verifi == true {
		return errors.New("usuário já cadastrado")
	}
	if _, verifi := s.repo.UsuarioRepositorySearch("", user.Email, 0); verifi == true {
		return errors.New("e-mail já cadastrado")
	}
	if _, err := s.roles.RolesRepositoryById(user.Role); err != nil {
		return errors.New("regra não encontrada")
	}
	return s.repo.UsuarioRepositoryAdd(&user)
}

// buscar por id
func (s *UsuarioService) UsuarioServiceById(id int32) (*models.Usuario, error) {
	return s.repo.UsuarioRepositoryById(id)
}

// update
func (s *UsuarioService) UsuarioServiceUpdate(id int32, user models.Usuario) error {
	user.ID = int32(id)
	if _, verifi := s.repo.UsuarioRepositorySearch(user.Nome, user.Email, 0); verifi != true {
		if _, err := s.roles.RolesRepositoryById(int64(user.ID)); err != nil {
			return s.repo.UsuarioRepositoryUpdate(user)
		}
		return errors.New("usuario não encontrada")
	}
	return errors.New("não pode haver dados duplicados")
}

// delete
func (s *UsuarioService) UsuarioServiceDelete(id int32) error {
	if err := s.repoPost.PostRepositorySearch(int64(id)); err != nil {
		return s.repo.UsuarioRepositoryDelete(id)
	}
	return errors.New("Há usuário com dependencia")
}
