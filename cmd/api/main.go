package main

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/db"
	"github.com/BrunoPessoa097/api-go-simples/internal/handlers"
	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/route"
	"github.com/gin-gonic/gin"
)

// função principal
func main() {
	// desativando release mode
	gin.SetMode(gin.ReleaseMode)
	// iniciando servidor
	r := gin.Default()

	db := db.Sqlite()
	db.AutoMigrate(&models.Roles{}, &models.Usuario{})

	// handles inicial
	handleDefault := handlers.NewDefaultHandle()
	// rotas padrão
	route := route.NewRotasDefault(handleDefault, db)

	// routeamento
	route.InicialRota(r)

	// iniciando rota
	if err := r.Run(":8080"); err != nil {
		panic("Erro no servidor")
	}
}
