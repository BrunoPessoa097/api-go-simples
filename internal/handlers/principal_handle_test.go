package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/testify/v2/assert"
)

// tdd inicial handle
func TestInitialHandle(t *testing.T) {
	// iniciando
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	handler := NewDefaultHandle()

	// rotas
	router.GET("/", handler.InicialHandle)

	// entradas
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// saida
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEqual(t, http.StatusNotFound, "Problemas com a rota principal")
}

// tdd não encontrado
func TestNaoEncontrado(t *testing.T) {
	// iniciando
	handler := NewDefaultHandle()

	// rotas
	router := gin.Default()
	router.NoRoute(handler.NaoEncontrada)

	// definicao das entradas
	req, _ := http.NewRequest(http.MethodGet, "/error", nil)
	w := httptest.NewRecorder()

	// entradas e saidas
	router.ServeHTTP(w, req)

	// saida
	assert.Equal(t, http.StatusNotFound, w.Code)
}
