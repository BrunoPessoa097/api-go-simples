package route

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/handlers"
	"github.com/gin-gonic/gin"
)

func InicialRota(r *gin.Engine) {
	r.GET("/", handlers.InicialHandle)
	r.NoRoute(handlers.NaoEncontrada)
}
