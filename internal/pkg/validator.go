package pkg

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type pkg struct{}

func NewPkg() *pkg {
	return &pkg{}
}

func (p *pkg) Validator(err error) map[string]string {
	validationErros, ok := err.(validator.ValidationErrors)
	if !ok {
		return map[string]string{
			"error": err.Error(),
		}
	}

	errors := make(map[string]string)

	for _, fieldError := range validationErros {
		field := strings.ToLower(fieldError.Field())

		switch fieldError.Tag() {
		case "required":
			errors[field] = "campo obrigatório !"

		case "min":
			errors[field] = "mínimo " + fieldError.Param() + " caracteres."
		case "max":
			errors[field] = "máximo " + fieldError.Param() + "caracteres."

		case "email":
			errors[field] = "email inválido !"
		}
	}
	return errors
}
