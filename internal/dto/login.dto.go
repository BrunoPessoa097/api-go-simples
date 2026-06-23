package dto

import "github.com/BrunoPessoa097/api-go-simples/internal/models"

type LoginCreate struct {
	Email string `json:"email" binding:"required,email,min=3"`
	Senha string `json:"senha" binding:"required,min=8"`
}

func ToLogin(dto *LoginCreate) *models.Login {
	return &models.Login{
		Email: dto.Email,
		Senha: dto.Senha,
	}
}
