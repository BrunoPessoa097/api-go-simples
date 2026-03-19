package services

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/mocks"
	"github.com/BrunoPessoa097/api-go-simples/internal/models"
)

// estrutura
type RoleService struct{}

// construtor
func NewRoleService() *RoleService {
	return &RoleService{}
}

// listar regras
func (rs *RoleService) RoleServiceList() []models.RolesC {
	role := make([]models.RolesC, len(mocks.UsuariosBD)-2)

	for _, rol := range mocks.ListRoles {
		role = append(role, rol)
	}

	return role
}
