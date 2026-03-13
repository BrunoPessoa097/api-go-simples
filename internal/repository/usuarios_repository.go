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
func (r *UsuarioRepository) UsuarioRepositoryById(id int32) *models.UsuarioCriate {
	//mocks
	mokerUser := mocks.UsuariosBD

	//buscando
	for _, user := range mokerUser {
		if user.Id == id {

			//return
			return &user
		}
	}

	//retorno
	return nil
}

// update
func (r *UsuarioRepository) UsuarioRepositoryUpdate(id int32, update models.UsuarioCriate) bool {
	//mocks
	mokerUser := mocks.UsuariosBD
	userUp := update

	// buscando
	for i, user := range mokerUser {
		if user.Id == id {
			mokerUser[i] = userUp
			return true
		}
	}

	return false
}

func (r *UsuarioRepository) UsuarioRepositoryDelete(id int32) bool {
	for _, user := range mocks.UsuariosBD {
		if user.Id == id {
			mocks.UsuariosBD = append(mocks.UsuariosBD[:id], mocks.UsuariosBD[id+1:]...)
			return true
		}
	}
	return false
}
