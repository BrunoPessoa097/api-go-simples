package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/mocks"
	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
	"github.com/BrunoPessoa097/api-go-simples/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/testify/v2/assert"
)

// listagem
func TestRolesHandlerList(t *testing.T) {
	// iniciando
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	m := mocks.ListRoles
	repo := repository.NewRolesRepository(m)
	s := services.NewRoleService(repo)
	h := NewRolesHandler(s)

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
	m := mocks.ListRoles
	repo := repository.NewRolesRepository(m)
	s := services.NewRoleService(repo)
	h := NewRolesHandler(s)

	// modelo de negocio
	role := models.RolesC{
		Nivel: "vend1",
		Regra: "get,post",
	}

	// convertendo struct para json
	body, _ := json.Marshal(&role)

	//requisição
	req := httptest.NewRequest(http.MethodPost, "/roles", bytes.NewBuffer(body))
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
	m := mocks.ListRoles
	repo := repository.NewRolesRepository(m)
	s := services.NewRoleService(repo)
	h := NewRolesHandler(s)

	// requisicao
	req := httptest.NewRequest(http.MethodGet, "/roles/1", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// saida
	r.GET("/roles/:id", h.RolesHandlerById)
	r.ServeHTTP(w, req)

	// comparar
	assert.Equal(t, http.StatusOK, w.Code)
}

// updade de regras
func TestRolesHandlerUpdate(t *testing.T) {
	// iniciando
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	m := mocks.ListRoles
	repo := repository.NewRolesRepository(m)
	s := services.NewRoleService(repo)
	h := NewRolesHandler(s)

	// modelo de negocio
	role := models.RolesC{
		ID:    1,
		Nivel: "vendedor",
		Regra: "get,post,delete,put",
	}

	// convertendo struct para json
	rolec, _ := json.Marshal(&role)

	// requisicao
	req := httptest.NewRequest(http.MethodPut, "/roles/1", bytes.NewBuffer(rolec))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// saida
	r.PUT("/roles/:id", h.RolesHandlerUpdate)
	r.ServeHTTP(w, req)

	// comparar
	assert.Equal(t, http.StatusOK, w.Code)
}

// delete
func TestRolesHandlerDelete(t *testing.T) {
	// iniciando
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	m := mocks.ListRoles
	repo := repository.NewRolesRepository(m)
	s := services.NewRoleService(repo)
	h := NewRolesHandler(s)

	//requisicao e escrita
	req := httptest.NewRequest(http.MethodDelete, "/roles/1", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// acessando a rota
	r.DELETE("/roles/:id", h.RolesHandlerDelete)
	r.ServeHTTP(w, req)

	// saida
	assert.Equal(t, http.StatusNoContent, w.Code)
}
