package models

import "time"

// model de usuário
type UsuarioCriate struct {
	Id        int32  `json:"id"`
	Nome      string `json:"nome"`
	Email     string `json:"email"`
	Senha     string `json:"senha"`
	Role      int64  `json:"role"`
	Bloqueado bool   `json:"bloqueado"`
	DtCreate  time.Time
	DtUpdate  time.Time
}

type UsuarioUpdate struct {
	Nome      string `json:"nome"`
	Email     string `json:"email"`
	Senha     string `json:"senha"`
	Role      int64  `json:"role"`
	Bloqueado bool   `json:"bloqueado"`
	DtCreate  time.Time
	DtUpdate  time.Time
}
