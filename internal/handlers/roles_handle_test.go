package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/testify/v2/assert"
)

func TestRolesHandlerList(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	h := NewRolesHandler()

	req := httptest.NewRequest(http.MethodGet, "/roles", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.GET("/roles", h.RolesHandlerList)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
