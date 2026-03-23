package repository

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/models"
)

type RolesRepository struct {
	data []models.RolesC
}

func NewRolesRepository(mr []models.RolesC) *RolesRepository {
	return &RolesRepository{
		data: mr,
	}
}

func (r *RolesRepository) RolesRepositoryList() []models.RolesC {
	return r.data
}
