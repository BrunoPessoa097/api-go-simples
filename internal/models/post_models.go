package models

import "time"

// strutura de postagem

type Post struct {
	ID       int64
	IDUser   int64
	Texto    string
	DtCreate time.Time
	DtUpdate time.Time
}
