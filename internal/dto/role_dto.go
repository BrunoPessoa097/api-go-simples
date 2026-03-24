package dto

import (
	"time"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
)

// usuario criação
type RolesCreateDTO struct {
	Nivel string `json:"nivel" binding:"required,min=2,max=10"`
	Regra string `json:"regra" binding:"required,min=2,max=50"`
}

// usuario update
type RolesUpdateDTO struct {
	Nivel *string `json:"nivel" binding:"min=2,max=10"`
	Regra *string `json:"regra" binding:"min=2,max=50"`
}

type RolesResponseDTO struct {
	ID    int64  `json:"id"`
	Nivel string `json:"nivel"`
	Regra string `json:"regra"`
}

func ToModelRoles(dto RolesCreateDTO) models.RolesC {
	return models.RolesC{
		Nivel:    dto.Nivel,
		Regra:    dto.Regra,
		DtCreate: time.Now(),
		DtUpdate: time.Now(),
	}
}

func ToUpdateRoles(dto RolesUpdateDTO, model models.RolesC) models.RolesC {
	if dto.Nivel != nil {
		model.Nivel = *dto.Nivel
	}

	if dto.Regra != nil {
		model.Regra = *dto.Regra
	}

	model.DtUpdate = time.Now()

	return model
}

func ToResponseRoles(u models.RolesC) RolesResponseDTO {
	return RolesResponseDTO{
		ID:    u.ID,
		Nivel: u.Nivel,
		Regra: u.Regra,
	}
}

func ToResponseRolesList(u []models.RolesC) []RolesResponseDTO {
	var roles []RolesResponseDTO

	for i := range u {
		roles = append(roles, ToResponseRoles(u[i]))
	}
	return roles
}
