package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/testify/v2/assert"
)

func TestInitialHandle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	handler := NewDefaultHandle()

	router.GET("/", handler.InicialHandle)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEqual(t, http.StatusNotFound, "Problemas com a rota principal")
}

func TestNaoEncontrado(t *testing.T) {
	handler := NewDefaultHandle()

	router := gin.Default()

	router.NoRoute(handler.NaoEncontrada)

	req, _ := http.NewRequest(http.MethodGet, "/error", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
