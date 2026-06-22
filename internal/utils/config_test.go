package utils

import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func TestHash(t *testing.T) {
	senha := Hash("hoje")

	assert.Equal(t, senha, senha)
}

func TestCompare(t *testing.T) {
	recebido := "hoje"
	senha := Hash(recebido)
	comparar := Comparar(senha, recebido)

	assert.Equal(t, true, comparar)
}
