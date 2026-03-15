package main

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/handlers"
	"github.com/BrunoPessoa097/api-go-simples/internal/route"
	"github.com/gin-gonic/gin"
)

// função principal
func main() {
	// desativando release mode
	gin.SetMode(gin.ReleaseMode)
	// iniciando servidor
	r := gin.Default()
	r.Use(gin.Recovery())

	// handles inicial
	handleDefault := handlers.NewDefaultHandle()
	// rotas padrão
	route := route.NewRotasDefault(handleDefault)

	// routeamento
	route.InicialRota(r)

	// iniciando rota
	if err := r.Run(":8080"); err != nil {
		panic("Erro no servidor")
	}
}
