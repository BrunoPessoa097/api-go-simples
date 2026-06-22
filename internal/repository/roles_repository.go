package repository

import (
	"errors"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"gorm.io/gorm"
)

type RolesRepository struct {
	Data *gorm.DB
}

func NewRolesRepository(mr *gorm.DB) *RolesRepository {
	return &RolesRepository{
		Data: mr,
	}
}

// listagem
func (r *RolesRepository) RolesRepositoryList() ([]models.Roles, error) {
	var roles []models.Roles
	err := r.Data.Find(&roles).Error

	return roles, err
}

// buscar
func (r *RolesRepository) RolesRepositoryBusca(role string) bool {
	var roles models.Roles

	err := r.Data.
		Where("Nivel = ?", role).
		First(&roles).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	return err == nil
}

// adicionar
func (r *RolesRepository) RolesRepositoryAdd(roles *models.Roles) error {
	return r.Data.Create(roles).Error
}

// buscar id
func (r *RolesRepository) RolesRepositoryById(id int64) (*models.Roles, error) {
	var role models.Roles

	err := r.Data.First(&role, id).Error

	if err != nil {
		return nil, err
	}

	return &role, nil
}

// update
func (r *RolesRepository) RolesRepositoryUpdate(role *models.Roles) error {
	return r.Data.Model(&models.Post{}).Where("id", role.ID).Updates(role).Error
}

// delete
func (r *RolesRepository) RolesRepositoryDelete(id int64) error {
	return r.Data.Delete(&models.Roles{}, id).Error
}
