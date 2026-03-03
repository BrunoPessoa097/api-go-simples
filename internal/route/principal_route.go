package route

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/handlers"
	"github.com/gin-gonic/gin"
)

// rota inicial
func InicialRota(r *gin.Engine) {
	// rotas padrão
	r.GET("/", handlers.InicialHandle)
	r.NoRoute(handlers.NaoEncontrada)
}
