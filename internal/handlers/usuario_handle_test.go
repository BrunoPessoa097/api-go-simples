package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/testify/v2/assert"
)

func TestUsuarioListHandle(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	handle := NewUsuarioHandle()

	req := httptest.NewRequest(http.MethodGet, "/usuarios", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.GET("/usuarios", handle.UsuarioListHandle)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())
}
