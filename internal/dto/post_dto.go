package dto

import (
	"time"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
)

// criando post
type CreatePostDTO struct {
	IDUser int64  `json:"idUser" binding:"required,min=2,max=10"`
	Texto  string `json:"texto" binding:"required,min=2,max=50"`
}

// atualizando post
type UpdatePostDTO struct {
	IDUser *int64  `json:"idUser" binding:"required,min=2,max=10"`
	Texto  *string `json:"texto" binding:"required,min=2,max=50"`
}

// responss
type ResponsesPostDTO struct {
	IDUser int64  `json:"idUser"`
	Texto  string `json:"texto"`
}

func ToModelPostCreate(dto CreatePostDTO) *models.Post {
	return &models.Post{
		IDUser:   dto.IDUser,
		Texto:    dto.Texto,
		DtCreate: time.Now(),
		DtUpdate: time.Now(),
	}
}

func toModelPostUpdade(dto UpdatePostDTO, models models.Post) models.Post {
	if dto.Texto != nil {
		models.Texto = *dto.Texto
	}
	if dto.IDUser != nil {
		models.IDUser = *dto.IDUser
	}
	models.DtUpdate = time.Now()
	return models
}
