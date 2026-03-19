package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/testify/v2/assert"
)

// listagem
func TestRolesHandlerList(t *testing.T) {
	// iniciando
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	h := NewRolesHandler()

	// requisicao tipo json
	req := httptest.NewRequest(http.MethodGet, "/roles", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// rota com saida
	r.GET("/roles", h.RolesHandlerList)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// adicionar regras
func TestRolesHandlerPost(t *testing.T) {
	// iniciando
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	h := NewRolesHandler()

	// modelo de negocio
	role := models.RolesC{
		ID:    3,
		Nivel: "vendedor",
		Regra: "get,post,delete,put",
	}

	// convertendo struct para json
	rolec, _ := json.Marshal(&role)

	//requisição
	req := httptest.NewRequest(http.MethodPost, "/roles", bytes.NewBuffer(rolec))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// rota
	r.POST("/roles", h.RolesHandlerPost)
	r.ServeHTTP(w, req)

	// saida
	assert.Equal(t, http.StatusCreated, w.Code)
}

// buscar por id
func TestRolesHandlerById(t *testing.T) {
	// iniciando
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	h := NewRolesHandler()

	// requisicao
	req := httptest.NewRequest(http.MethodGet, "/roles/:id", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// saida
	r.GET("/roles/:id", h.RolesHandlerById)
	r.ServeHTTP(w, req)

	// comparar
	assert.Equal(t, http.StatusOK, w.Code)
}
