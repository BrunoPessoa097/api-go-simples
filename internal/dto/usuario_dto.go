package dto

import (
	"time"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
)

// usuario criação
type UsuarioCriateDTO struct {
	Nome  string `json:"nome" binding:"required,min=3,max=20"`
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required,min=8,max=15"`
}

// usuario update
type UsuarioUpdateDTO struct {
	Nome      *string `json:"nome" binding:"min=3,max=20"`
	Email     *string `json:"email" binding:"email"`
	Senha     *string `json:"senha" binding:"min=8,max=15"`
	Role      *int64  `json:"role"`
	Bloqueado *bool   `json:"bloqueado"`
}

type UsuarioResponseDTO struct {
	ID        int32     `json:"id"`
	Nome      string    `json:"nome"`
	Email     string    `json:"email"`
	Role      int64     `json:"role"`
	Bloqueado bool      `json:"bloqueado"`
	DtCreate  time.Time `json:"dt_create"`
}

func ToModel(dto UsuarioCriateDTO) models.Usuario {
	return models.Usuario{
		Nome:      dto.Nome,
		Email:     dto.Email,
		Senha:     dto.Senha,
		Role:      1,
		Bloqueado: false,
		DtCreate:  time.Now(),
		DtUpdate:  time.Now(),
	}
}

func ToUpdate(dto UsuarioUpdateDTO) models.Usuario {
	return models.Usuario{
		Nome:      *dto.Nome,
		Email:     *dto.Email,
		Senha:     *dto.Senha,
		Role:      *dto.Role,
		Bloqueado: *dto.Bloqueado,
	}
}

func ToResponse(u models.Usuario) UsuarioResponseDTO {
	return UsuarioResponseDTO{
		ID:        u.ID,
		Nome:      u.Nome,
		Email:     u.Email,
		Role:      u.Role,
		Bloqueado: u.Bloqueado,
		DtCreate:  u.DtCreate,
	}
}

func ToResponseList(u []models.Usuario) []UsuarioResponseDTO {
	var res []UsuarioResponseDTO
	for i := range u {
		res = append(res, ToResponse(u[i]))
	}
	return res
}
