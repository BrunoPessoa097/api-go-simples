package route

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/handlers"
	"github.com/gin-gonic/gin"
)

type PostRoutes struct {
	handlers *handlers.PostHandlers
}

func NewPostRoutas(h *handlers.PostHandlers) *PostRoutes {
	return &PostRoutes{handlers: h}
}

func (p *PostRoutes) PostRotas(rg *gin.Engine) {
	post := rg.Group("/posts")
	{
		post.GET("/", p.handlers.PostHandlersList)
		post.POST("/", p.handlers.PostHandlersPost)
		post.GET("/:id", p.handlers.PostHandlersById)
		post.PATCH("/:id", p.handlers.PostHandlersUpdate)
		post.DELETE("/:id", p.handlers.PostHandlersDelete)
	}
}
