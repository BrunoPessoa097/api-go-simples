package services

import "github.com/BrunoPessoa097/api-go-simples/internal/models"

type UsuarioService struct{}

// construtor
func NewUsuarioService() *UsuarioService {
	return &UsuarioService{}
}

// listar
func (s *UsuarioService) UsuarioServiceList() []models.UsuarioCriate {
	return make([]models.UsuarioCriate, 0)
}

// adiciona
func (s *UsuarioService) UsuarioServiceAdd() bool {
	return true
}

// buscar por id
func (s *UsuarioService) UsuarioServiceById(id int) models.UsuarioCriate {
	return models.UsuarioCriate{
		Id:        0,
		Nome:      "Bruno",
		Email:     "brunopessoa@gmail.com",
		Senha:     "1234",
		Role:      1,
		Bloqueado: false,
	}
}

// update
func (s *UsuarioService) UsuarioServiceUpdate(id int, user models.UsuarioCriate) bool {
	return true
}

// delete
func (s *UsuarioService) UsuarioServiceDelete(id int) bool {
	return true
}
