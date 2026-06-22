package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Sqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("social.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	log.Println("conectado com sucesso")
	return db
}
