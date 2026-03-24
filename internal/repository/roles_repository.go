package repository

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/models"
)

type RolesRepository struct {
	Data []models.Roles
}

func NewRolesRepository(mr []models.Roles) *RolesRepository {
	return &RolesRepository{
		Data: mr,
	}
}

// listagem
func (r *RolesRepository) RolesRepositoryList() []models.Roles {
	return r.Data
}

// buscar
func (r *RolesRepository) RolesRepositoryBusca(role string) bool {
	for i := range r.Data {
		if r.Data[i].Nivel == role {
			return true
		}
	}
	return false
}

// adicionar
func (r *RolesRepository) RolesRepositoryAdd(roles *models.Roles) bool {
	r.Data = append(r.Data, *roles)
	return true
}

// buscar id
func (r *RolesRepository) RolesRepositoryById(id int64) *models.Roles {
	for i := range r.Data {
		if r.Data[i].ID == id {
			return &r.Data[i]
		}
	}
	return nil
}

// update
func (r *RolesRepository) RolesRepositoryUpdate(id int64, role models.Roles) bool {
	for i := range r.Data {
		if r.Data[i].ID == id {
			r.Data[i] = role
			return true
		}
	}
	return false
}

// delete
func (r *RolesRepository) RolesRepositoryDelete(id int64) bool {
	for i := range r.Data {
		//caso exista
		if r.Data[i].ID == id {
			r.Data = append(r.Data[:i], r.Data[i+1:]...)
			return true
		}
	}
	return false
}
