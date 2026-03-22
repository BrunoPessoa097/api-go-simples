package repository

import (
	"errors"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
)

type UsuarioRepository struct {
	Data []models.UsuarioCriate
}

// construtor
func NewUsuarioRepository(m []models.UsuarioCriate) *UsuarioRepository {
	return &UsuarioRepository{
		Data: m,
	}
}

// listar os usuarios
func (r *UsuarioRepository) UsuarioRepositoryList() []models.UsuarioCriate {
	return r.Data
}

// adicionar
func (r *UsuarioRepository) UsuarioRepositoryAdd(user models.UsuarioCriate) error {
	// verificando a existencia
	if saida := r.UsuarioRepositorySearch(user.Nome, user.Email); saida {
		return errors.New("Usuario e/ou E-mail já cadastrados")
	}

	// registrando o usuario
	r.Data = append(r.Data, user)

	//retorno
	return nil
}

// pegar por id
func (r *UsuarioRepository) UsuarioRepositoryById(id int32) *models.UsuarioCriate {
	//buscando
	for i := range r.Data {
		if r.Data[i].Id == id {
			//return
			return &r.Data[i]
		}
	}
	//retorno
	return nil
}

// update
func (r *UsuarioRepository) UsuarioRepositoryUpdate(id int32, update models.UsuarioCriate) error {
	// buscando
	if err := r.UsuarioRepositorySearch(update.Nome, update.Email); err {
		return errors.New("usuário e email já cadastrados")
	}

	for i := range r.Data {
		if r.Data[i].Id == id {
			r.Data[i] = update
			return nil
		}
	}
	return errors.New("falha ao atualizar")
}

// usuario delete
func (r *UsuarioRepository) UsuarioRepositoryDelete(id int32) error {
	for i := range r.Data {
		if r.Data[i].Id == id {
			r.Data = append(r.Data[:i], r.Data[i+1:]...)
			return nil
		}
	}
	return errors.New("Usuario impossivel de excluir")
}

// buscando existencia
func (r *UsuarioRepository) UsuarioRepositorySearch(nome, email string) bool {
	for _, user := range r.Data {
		if user.Nome == nome || user.Email == email {
			return true
		}
	}
	return false
}
