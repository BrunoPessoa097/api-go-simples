package models

import "time"

type RolesC struct {
	ID       int64  `json:"id"`
	Nivel    string `json:"nivel" binding:"required,min=2,max=10"`
	Regra    string `json:"regra" binding:"required,min=2,max=50"`
	DtCreate time.Time
	DtUpdate time.Time
}
