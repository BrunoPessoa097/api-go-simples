package mocks

import "github.com/BrunoPessoa097/api-go-simples/internal/models"

var ListPost = []models.Post{
	{
		ID:     1,
		IDUser: 1,
		Texto:  "Como é viver no Brasil?",
	},
	{
		ID:     2,
		IDUser: 2,
		Texto:  "Como é viver no na Inglaterra?",
	},
	{
		ID:     3,
		IDUser: 1,
		Texto:  "Como é viver no Portugal?",
	},
}
