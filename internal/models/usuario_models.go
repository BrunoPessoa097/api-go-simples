package models

import "time"

// model de usuário
type Usuario struct {
	ID        int32 `gorm:"primaryKey;autoIncrement"`
	Nome      string
	Email     string
	Senha     string
	Role      int64
	Bloqueado bool
	DtCreate  time.Time `gorm:"autoCreateTime"`
	DtUpdate  time.Time `gorm:"autoUpdateTime"`
}
