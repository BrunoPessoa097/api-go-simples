package main

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/route"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	route.InicialRota(r)

	r.Run(":8080")
}
