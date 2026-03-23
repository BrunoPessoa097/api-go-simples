package services

import (
	"errors"

	"github.com/BrunoPessoa097/api-go-simples/internal/mocks"
	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
)

// estrutura
type RoleService struct {
	repo *repository.RolesRepository
}

// construtor
func NewRoleService(r *repository.RolesRepository) *RoleService {
	return &RoleService{repo: r}
}

// listar regras
func (rs *RoleService) RoleServiceList() []models.RolesC {
	return rs.repo.RolesRepositoryList()
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
	id := len(mocks.ListRoles) + 1
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

// buscar unico
func (rs *RoleService) RoleServiceById(id int64) (*models.RolesC, error) {
	//pegar um model
	var role models.RolesC
	//pegando a pessoa
	for _, roles := range mocks.ListRoles {
		if roles.ID == id {
			role = roles
			return &role, nil
		}

	}
	//return
	return nil, errors.New("regra não encontrado")
}

// deletar
func (rs *RoleService) RoleServiceDelete(id int64) bool {
	for _, roles := range mocks.ListRoles {
		//caso exista
		if roles.ID == id {
			mocks.ListRoles = append(mocks.ListRoles[id:], mocks.ListRoles[id+1:]...)
			return true
		}
	}
	return false
}
