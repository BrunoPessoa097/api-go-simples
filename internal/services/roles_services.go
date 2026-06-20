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
func (rs *RoleService) RoleServiceList() []models.Roles {
	roles, _ := rs.repo.RolesRepositoryList()
	return roles
}

// buscando regra se existe
func (rs *RoleService) RoleServiceSearch(nivel string) error {
	if saida := rs.repo.RolesRepositoryBusca(nivel); saida {
		return errors.New("Nivel já cadastrado")
	}
	return nil
}

// adicionando post
func (rs *RoleService) RoleServicePost(roles models.Roles) error {
	return rs.repo.RolesRepositoryAdd(&roles)
}

// buscar unico
func (rs *RoleService) RoleServiceById(id int64) (*models.Roles, error) {
	return rs.repo.RolesRepositoryById(id)
}

func (rs *RoleService) RoleServiceUpdate(id int64, role models.Roles) error {
	role.ID = id
	if err := rs.repo.RolesRepositoryUpdate(&role); err != nil {
		return err
	}
	return nil
}

// // deletar
func (rs *RoleService) RoleServiceDelete(id int64) error {
	if err := rs.repo.RolesRepositoryDelete(id); err != nil {
		return errors.New("regra não encontrada")
	}
	return nil
}
