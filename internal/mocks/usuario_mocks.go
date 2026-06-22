package mocks

import "github.com/BrunoPessoa097/api-go-simples/internal/models"

var (
	UsuariosBD = []models.Usuario{
		{
			ID:        1,
			Nome:      "Bruno F",
			Email:     "brunopessoa@gmail.com",
			Senha:     "1234",
			Role:      1,
			Bloqueado: false,
		},
		{
			ID:        2,
			Nome:      "Bruno P",
			Email:     "brunopessoa@gmail.com",
			Senha:     "1234",
			Role:      1,
			Bloqueado: false,
		},
	}
)
