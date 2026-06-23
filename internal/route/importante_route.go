package route

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/handlers"
	"github.com/gin-gonic/gin"
)

type Carregar struct {
	handler *handlers.Importante
}

func NewLoginRoute(h *handlers.Importante) *Carregar {
	return &Carregar{
		handler: h,
	}
}

func (l *Carregar) Login(r *gin.Engine) {
	login := r.Group("/login")
	{
		login.POST("/", l.handler.ImportanteLogin)
	}
}
