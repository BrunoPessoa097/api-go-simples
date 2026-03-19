package services

import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

// listagem
func TestRoleServiceList(t *testing.T) {
	//iniciando
	s := NewRoleService()

	//listagem
	saida := s.RoleServiceList()

	//verificação
	assert.Equal(t, "ADM", saida[0].Nivel)
}
