package services

import (
	"errors"

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

// adicionando post
func (rs *RoleService) RoleServicePost(roles models.RolesC) (bool, error) {
	//recebendo os valores
	role := roles

	// verificando a existencia
	ok := rs.RoleServiceSearch(role.Nivel)

	//saida de erros
	if ok {
		return false, errors.New("Regra já cadastrada")
	}

	//adicionando id
	id := len(mocks.ListRoles) - 1
	role.ID = int64(id)

	//adicionando
	mocks.ListRoles = append(mocks.ListRoles, role)

	// retorno
	return true, nil
}

// buscando regra se existe
func (rs *RoleService) RoleServiceSearch(nome string) bool {
	for _, role := range mocks.ListRoles {
		//retorno caso exista
		if role.Nivel == nome {
			return true
		}
	}
	return false
}
