package utils

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// data so servidor
var DtHr = time.Now().Format("01-02-2006 15:04:05")

// carregar o dotenv
func Dotenv(variaveis string) (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return "", errors.New("Erro ao carregar as variáveis do .env")
	}

	return os.Getenv(variaveis), nil
}

func SetupDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatal(err)
	}

	if err := db.AutoMigrate(&models.Roles{}, &models.Usuario{}); err != nil {
		t.Fatal(err)
	}

	return db
}
