package models

import "time"

type Mensagem struct {
	ID       int64
	IDUser   string
	Texto    string
	DtCreate time.Time
	DtUpdate time.Time
}
