package mocks

import "github.com/BrunoPessoa097/api-go-simples/internal/models"

var ListRoles = []models.Roles{
	{
		ID:    1,
		Nivel: "ADM",
		Regra: "get,post,delete,put",
	},
	{
		ID:    2,
		Nivel: "Gerente",
		Regra: "get,post,delete,put",
	},
	{
		ID:    3,
		Nivel: "publico",
		Regra: "get",
	},
}
