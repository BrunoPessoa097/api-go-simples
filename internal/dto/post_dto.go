package dto

import (
	"time"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
)

// criando post
type CreatePostDTO struct {
	IDUser int64  `json:"idUser" binding:"required,min=1"`
	Texto  string `json:"texto" binding:"required,min=2,max=50"`
}

// atualizando post
type UpdatePostDTO struct {
	IDUser *int64  `json:"idUser" validate:"omitempty,min=2"`
	Texto  *string `json:"texto" validate:"required,min=2,max=50"`
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

func ToModelPostUpdade(dto UpdatePostDTO) models.Post {
	var post models.Post

	if dto.IDUser != nil {
		post.IDUser = *dto.IDUser
	}

	if dto.Texto != nil {
		post.Texto = *dto.Texto
	}

	post.DtUpdate = time.Now()

	return post
}

func ToModelPostListOne(p models.Post) ResponsesPostDTO {
	return ResponsesPostDTO{
		IDUser: p.IDUser,
		Texto:  p.Texto,
	}
}

func ToModelPostList(p []models.Post) []ResponsesPostDTO {
	var post []ResponsesPostDTO

	for i := range p {
		post = append(post, ToModelPostListOne(p[i]))
	}
	return post
}
