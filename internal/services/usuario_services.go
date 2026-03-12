package services

import "github.com/BrunoPessoa097/api-go-simples/internal/models"

type UsuarioService struct{}

func NewUsuarioService() *UsuarioService {
	return &UsuarioService{}
}

func (s *UsuarioService) UsuarioServiceList() []models.UsuarioCriate {
	return make([]models.UsuarioCriate, 0)
}

func (s *UsuarioService) UsuarioServiceAdd() bool {
	return true
}

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
