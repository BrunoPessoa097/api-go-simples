package mocks

import "github.com/BrunoPessoa097/api-go-simples/internal/models"

var (
	UsuariosBD = []models.UsuarioCriate{
		{
			Id:        1,
			Nome:      "Bruno F",
			Email:     "brunopessoa@gmail.com",
			Senha:     "1234",
			Role:      1,
			Bloqueado: false,
		},
		{
			Id:        2,
			Nome:      "Bruno P",
			Email:     "brunopessoa@gmail.com",
			Senha:     "1234",
			Role:      1,
			Bloqueado: false,
		},
	}
)
