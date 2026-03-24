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
func (rs *RoleService) RoleServicePost(roles models.Roles) bool {
	//adicionando
	roles.ID = int64(len(rs.repo.Data) + 1)
	return rs.repo.RolesRepositoryAdd(&roles)
}

// buscar unico
func (rs *RoleService) RoleServiceById(id int64) *models.Roles {
	if saida := rs.repo.RolesRepositoryById(id); saida != nil {
		return saida
	}
	return nil
}

func (rs *RoleService) RoleServiceUpdate(id int64, role models.Roles) error {
	role.ID = id
	if err := rs.repo.RolesRepositoryUpdate(role.ID, role); err {
		return nil
	}
	return errors.New("Erro ao atualizar a regra")
}

// // deletar
func (rs *RoleService) RoleServiceDelete(id int64) error {
	if ok := rs.repo.RolesRepositoryDelete(id); ok {
		return nil
	}
	return errors.New("regra não encontrada")
}
