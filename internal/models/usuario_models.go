package models

// model de usuário
type UsuarioCriate struct {
	Id    int32  `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}
