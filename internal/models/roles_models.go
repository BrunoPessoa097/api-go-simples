package models

import "time"

type Roles struct {
	ID       int64
	Nivel    string
	Regra    string
	DtCreate time.Time
	DtUpdate time.Time
}
