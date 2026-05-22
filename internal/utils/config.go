package utils

import (
	"errors"
	"os"
	"time"

	"github.com/joho/godotenv"
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
