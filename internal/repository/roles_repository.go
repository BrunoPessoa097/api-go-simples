package repository

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/models"
)

type RolesRepository struct {
	Data []models.RolesC
}

func NewRolesRepository(mr []models.RolesC) *RolesRepository {
	return &RolesRepository{
		Data: mr,
	}
}

// listagem
func (r *RolesRepository) RolesRepositoryList() []models.RolesC {
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
func (r *RolesRepository) RolesRepositoryAdd(roles *models.RolesC) bool {
	r.Data = append(r.Data, *roles)
	return true
}

// buscar if
func (r *RolesRepository) RolesRepositoryById(id int64) *models.RolesC {
	for i := range r.Data {
		if r.Data[i].ID == id {
			return &r.Data[i]
		}
	}
	return nil
}
