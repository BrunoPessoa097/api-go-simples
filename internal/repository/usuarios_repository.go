package repository

import (
	"errors"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"gorm.io/gorm"
)

type UsuarioRepository struct {
	Data *gorm.DB
}

// construtor
func NewUsuarioRepository(m *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{
		Data: m,
	}
}

// listar os usuarios
func (r *UsuarioRepository) UsuarioRepositoryList() ([]models.Usuario, error) {
	var user []models.Usuario
	err := r.Data.Find(&user).Error
	return user, err
}

// adicionar
func (r *UsuarioRepository) UsuarioRepositoryAdd(user models.Usuario) error {
	// verificando a existencia
	// if saida := r.UsuarioRepositorySearch(user.Nome, user.Email); saida {
	// 	return errors.New("Usuario e/ou E-mail já cadastrados")
	// }

	return r.Data.Create(&user).Error
}

// pegar por id
func (r *UsuarioRepository) UsuarioRepositoryById(id int32) (*models.Usuario, error) {
	//buscando
	var user models.Usuario

	err := r.Data.First(&user, id).Error

	if err != nil {
		return nil, err
	}
	//retorno
	return &user, nil
}

// update
func (r *UsuarioRepository) UsuarioRepositoryUpdate(update models.Usuario) error {
	// buscando
	// if err := r.UsuarioRepositorySearch(update.Nome, update.Email); err {
	// 	return errors.New("usuário e email já cadastrados")
	// }

	return r.Data.Save(update).Error
}

// usuario delete
func (r *UsuarioRepository) UsuarioRepositoryDelete(id int32) error {
	return r.Data.Delete(&models.Usuario{}, id).Error
}

// buscando existencia
func (r *UsuarioRepository) UsuarioRepositorySearch(nome, email string) bool {
	var user models.Usuario

	err := r.Data.Where(&models.Usuario{
		Nome:  nome,
		Email: email,
	}).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	return err == nil
}
