package repository

import (
	"errors"

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
func (r *UsuarioRepository) UsuarioRepositoryAdd(user models.UsuarioCriate) (error, bool) {
	mocksUser := mocks.UsuariosBD

	// verificando a existencia
	if saida := r.UsuarioRepositorySearch(user.Nome, user.Email); saida {
		return errors.New("Usuario e/ou E-mail já cadastrados"), true
	}

	// registrando o usuario
	mocksUser = append(mocksUser, user)

	//retorno
	return nil, false
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

	// saida := r.UsuarioRepositorySearch()

	// buscando
	for i, user := range mokerUser {
		if user.Id == id {
			mokerUser[i] = userUp
			return true
		}
	}

	return false
}

// usuario delete
func (r *UsuarioRepository) UsuarioRepositoryDelete(id int32) bool {
	for _, user := range mocks.UsuariosBD {
		if user.Id == id {
			mocks.UsuariosBD = append(mocks.UsuariosBD[:id], mocks.UsuariosBD[id+1:]...)
			return true
		}
	}
	return false
}

// buscando existencia
func (r *UsuarioRepository) UsuarioRepositorySearch(nome, email string) bool {
	for _, user := range mocks.UsuariosBD {
		if user.Nome == nome || user.Email == email {
			return true
		}
	}
	return false
}
