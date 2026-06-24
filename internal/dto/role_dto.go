package dto

import (
	"time"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
)

// usuario criação
type RolesCreateDTO struct {
	Nivel string `json:"nivel" binding:"required,min=2,max=10"`
	Rotas string `json:"rotas" binding:"required,min=2,max=100"`
	Regra string `json:"regra" binding:"required,min=2,max=50"`
}

// usuario update
type RolesUpdateDTO struct {
	Nivel *string `json:"nivel" binding:"min=2,max=10"`
	Rotas *string `json:"rotas" binding:"min=2,max=100"`
	Regra *string `json:"regra" binding:"min=2,max=50"`
}

type RolesResponseDTO struct {
	ID    int64  `json:"id"`
	Nivel string `json:"nivel"`
	Rotas string `json:"rotas"`
	Regra string `json:"regra"`
}

func ToModelRoles(dto RolesCreateDTO) models.Roles {
	return models.Roles{
		Nivel:    dto.Nivel,
		Regra:    dto.Regra,
		Rotas:    dto.Rotas,
		DtCreate: time.Now(),
		DtUpdate: time.Now(),
	}
}

func ToUpdateRoles(dto RolesUpdateDTO, model models.Roles) models.Roles {
	if dto.Nivel != nil {
		model.Nivel = *dto.Nivel
	}

	if dto.Regra != nil {
		model.Regra = *dto.Regra
	}

	if dto.Rotas != nil {
		model.Rotas = *dto.Rotas
	}

	model.DtUpdate = time.Now()

	return model
}

func ToResponseRoles(u models.Roles) RolesResponseDTO {
	return RolesResponseDTO{
		ID:    u.ID,
		Nivel: u.Nivel,
		Rotas: u.Rotas,
		Regra: u.Regra,
	}
}

func ToResponseRolesList(u []models.Roles) []RolesResponseDTO {
	var roles []RolesResponseDTO

	for i := range u {
		roles = append(roles, ToResponseRoles(u[i]))
	}
	return roles
}
