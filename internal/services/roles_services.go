package services

import (
	"errors"

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

// buscando regra se existe
func (rs *RoleService) RoleServiceSearch(nivel string) error {
	if saida := rs.repo.RolesRepositoryBusca(nivel); saida {
		return errors.New("Nivel já cadastrado")
	}
	return nil
}

// adicionando post
func (rs *RoleService) RoleServicePost(roles models.RolesC) bool {
	//adicionando
	roles.ID = int64(len(rs.repo.Data) + 1)
	return rs.repo.RolesRepositoryAdd(&roles)
}

// buscar unico
func (rs *RoleService) RoleServiceById(id int64) *models.RolesC {
	if saida := rs.repo.RolesRepositoryById(id); saida != nil {
		return saida
	}
	return nil
}

// // deletar
// func (rs *RoleService) RoleServiceDelete(id int64) bool {
// 	for _, roles := range mocks.ListRoles {
// 		//caso exista
// 		if roles.ID == id {
// 			mocks.ListRoles = append(mocks.ListRoles[id:], mocks.ListRoles[id+1:]...)
// 			return true
// 		}
// 	}
// 	return false
// }
