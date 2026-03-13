package repository

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/mocks"
	"github.com/BrunoPessoa097/api-go-simples/internal/models"
)

type UsuarioRepository struct{}

// construtor
func NewUsuarioRepository() *UsuarioRepository {
	return &UsuarioRepository{}
}

// listar os usuarios
func (r *UsuarioRepository) UsuarioRepositoryList() []models.UsuarioCriate {
	return mocks.UsuariosBD
}

// adicionar
func (r *UsuarioRepository) UsuarioRepositoryAdd(user models.UsuarioCriate) bool {
	mocksUser := mocks.UsuariosBD

	if mocksUser = append(mocksUser, user); mocksUser != nil {
		return true
	}
	return false
}

// pegar por id
func (r *UsuarioRepository) UsuarioRepositoryById(id int32) (models.UsuarioCriate, bool) {
	//mocks
	mokerUser := mocks.UsuariosBD

	//buscando
	for _, user := range mokerUser {
		if user.Id == id {
			return user, true
		}
	}

	//retorno
	return models.UsuarioCriate{}, false
}
