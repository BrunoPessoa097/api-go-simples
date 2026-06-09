package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Sqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("social.db"), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco")
	}
	log.Println("conectado com sucesso")
	return db
}
