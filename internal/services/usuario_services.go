package services

import (
	"errors"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
)

type UsuarioService struct {
	repo  *repository.UsuarioRepository
	roles *repository.RolesRepository
}

// construtor
func NewUsuarioService(repo *repository.UsuarioRepository, roles *repository.RolesRepository) *UsuarioService {
	return &UsuarioService{
		repo:  repo,
		roles: roles,
	}
}

// listar
func (s *UsuarioService) UsuarioServiceList() ([]models.Usuario, error) {
	return s.repo.UsuarioRepositoryList()
}

// adiciona
func (s *UsuarioService) UsuarioServiceAdd(user models.Usuario) error {
	if verifi := s.repo.UsuarioRepositorySearch(user.Nome, user.Email); verifi != true {
		if _, err := s.roles.RolesRepositoryById(user.Role); err == nil {
			return s.repo.UsuarioRepositoryAdd(&user)
		}
		return errors.New("regra não encontrada")
	}

	return errors.New("usuário já cadastrado")
}

// buscar por id
func (s *UsuarioService) UsuarioServiceById(id int32) (*models.Usuario, error) {
	return s.repo.UsuarioRepositoryById(id)
}

// update
func (s *UsuarioService) UsuarioServiceUpdate(id int32, user models.Usuario) error {
	user.ID = int32(id)
	if verifi := s.repo.UsuarioRepositorySearch(user.Nome, user.Email); verifi != true {
		if _, err := s.roles.RolesRepositoryById(int64(user.ID)); err != nil {
			return s.repo.UsuarioRepositoryUpdate(user)
		}
		return errors.New("usuario não encontrada")
	}
	return errors.New("não pode haver duplicados")
}

// delete
func (s *UsuarioService) UsuarioServiceDelete(id int32) error {
	return s.repo.UsuarioRepositoryDelete(id)
}
