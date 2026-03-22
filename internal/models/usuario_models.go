package models

import "time"

// model de usuário
type Usuario struct {
	Id        int32
	Nome      string
	Email     string
	Senha     string
	Role      int64
	Bloqueado bool
	DtCreate  time.Time
	DtUpdate  time.Time
}
