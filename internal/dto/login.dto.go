package dto

type LoginCreate struct {
	Email string `json:"email" binding:"required"`
	Senha string `json:"senha" binding:"required"`
}
