package repository

import (
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/mocks"
	"github.com/go-openapi/testify/v2/assert"
)

func TestList(t *testing.T) {
	mc := mocks.ListRoles
	rp := NewRolesRepository(mc)

	saida := rp.RolesRepositoryList()
	espc := mc[0].Nivel

	assert.Equal(t, espc, saida[0].Nivel)
}
