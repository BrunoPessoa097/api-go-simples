package pkg

import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Nome  string `validate:"required,min=3,max=10"`
	Email string `validate:"required,email"`
}

func TestValidator(t *testing.T) {
	validate := validator.New()

	user := User{
		Nome:  "",
		Email: "email_invalido",
	}

	err := validate.Struct(user)

	p := NewPkg()
	result := p.Validator(err)

	assert.Equal(t, "campo obrigatório !", result["nome"])
	assert.Equal(t, "email inválido !", result["email"])
}
