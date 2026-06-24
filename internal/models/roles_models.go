package models

import "time"

type Roles struct {
	ID       int64 `gorm:"primaryKey;autoIncrement"`
	Nivel    string
	Regra    string
	Rotas    string
	DtCreate time.Time `gorm:"autoCreateTime"`
	DtUpdate time.Time `gorm:"autoUpdateTime"`
}
