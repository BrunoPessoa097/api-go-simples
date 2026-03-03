package main

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/route"
	"github.com/gin-gonic/gin"
)

func main() {
	// desativando release mode
	gin.SetMode(gin.ReleaseMode)
	// iniciando servidor
	r := gin.Default()

	// routeamento
	route.InicialRota(r)

	// iniciando rota
	if err := r.Run(":8080"); err != nil {
		panic("Erro no servidor")
	}
}
