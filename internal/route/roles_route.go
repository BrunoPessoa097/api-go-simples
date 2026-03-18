package route

import (
	"github.com/BrunoPessoa097/api-go-simples/internal/handlers"
	"github.com/gin-gonic/gin"
)

// estrutura
type RolesRoute struct {
	handle *handlers.RolesHandler
}

// construtor
func NewRolesRoute(h *handlers.RolesHandler) *RolesRoute {
	return &RolesRoute{handle: h}
}

func (rr *RolesRoute) RolesRoutes(ctx *gin.Engine) {
	roles := ctx.Group("/roles")
	{
		roles.GET("/", rr.handle.RolesHandlerList)
	}
}
