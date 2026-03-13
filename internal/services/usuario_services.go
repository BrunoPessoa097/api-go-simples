package services

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
)

type UsuarioService struct {
	repo *repository.UsuarioRepository
}

// construtor
func NewUsuarioService(repo *repository.UsuarioRepository) *UsuarioService {
	return &UsuarioService{
		repo: repo,
	}
}

// listar
func (s *UsuarioService) UsuarioServiceList() []models.UsuarioCriate {
	return s.repo.UsuarioRepositoryList()
}

// adiciona
func (s *UsuarioService) UsuarioServiceAdd(user models.UsuarioCriate) bool {
	return s.repo.UsuarioRepositoryAdd(user)
}

// buscar por id
func (s *UsuarioService) UsuarioServiceById(id int64) *models.UsuarioCriate {
	return s.repo.UsuarioRepositoryById(int32(id))
}

// update
func (s *UsuarioService) UsuarioServiceUpdate(id int, user models.UsuarioCriate) bool {
	return s.repo.UsuarioRepositoryUpdate(int32(id), user)
}

// delete
func (s *UsuarioService) UsuarioServiceDelete(id int) bool {
	return s.repo.UsuarioRepositoryDelete(int32(id))
}
