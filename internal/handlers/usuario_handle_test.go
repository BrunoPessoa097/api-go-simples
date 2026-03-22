package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/start"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/testify/v2/assert"
)

// tdd lista de usuarios
func TestUsuarioListHandle(t *testing.T) {
	// inicializando
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	s := start.NewStart()
	handle := s.UsuarioStart()

	// requisição
	req := httptest.NewRequest(http.MethodGet, "/usuarios", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// rotas
	router.GET("/usuarios", handle.UsuarioListHandle)
	router.ServeHTTP(w, req)

	// saida
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())
}

// tdd post de usuarios
func TestUsuarioPostHandle(t *testing.T) {
	// iniciando
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	s := start.NewStart()
	handle := s.UsuarioStart()

	// entrada
	msg := models.UsuarioCriate{
		Nome:      "Bruno Frefre",
		Email:     "brunopessoa1234@gmail.com",
		Senha:     "12345678",
		Role:      1,
		Bloqueado: false,
	}

	// convertendo de json para bytes
	body, _ := json.Marshal(&msg)

	// requesicao
	req := httptest.NewRequest(http.MethodPost, "/usuarios", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// rotas e leitura
	router.POST("/usuarios", handle.UsuarioPostHandle)
	router.ServeHTTP(w, req)

	//resposta para converter em bites para json
	var resp models.UsuarioCriate
	json.Unmarshal(w.Body.Bytes(), &resp)

	//saida
	assert.Equal(t, http.StatusCreated, w.Code)
	// assert.Contains(t, w.Body.String(), resp.Nome)
}

// tdd pegar usuario por id
func TestUsuarioByIdHandle(t *testing.T) {
	//iniciar
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	s := start.NewStart()
	handler := s.UsuarioStart()

	//requesicao
	req := httptest.NewRequest(http.MethodGet, "/usuarios/:id", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// rota leitura e escrita
	router.GET("/usuarios/:id", handler.UsuarioByIdHandle)
	router.ServeHTTP(w, req)

	// recebendo para converter para json
	var resp models.UsuarioCriate
	json.Unmarshal(w.Body.Bytes(), &resp)

	//saida
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), resp.Nome)
}

// tdd update de usuario
func TestUsuarioUpdateHandle(t *testing.T) {
	//inicializando
	gin.SetMode(gin.TestMode)
	route := gin.Default()
	s := start.NewStart()
	handler := s.UsuarioStart()

	// entrada
	msg := models.UsuarioCriate{
		Nome:      "Bruno vv",
		Email:     "ba122@gmail.com",
		Senha:     "12345678",
		Role:      1,
		Bloqueado: false,
	}

	//convertendo para json
	user, _ := json.Marshal(&msg)

	//requisicao
	req := httptest.NewRequest(http.MethodPut, "/usuarios/1", bytes.NewBuffer(user))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// leitura e escrita
	route.PUT("/usuarios/:id", handler.UsuarioUpdateHandle)
	route.ServeHTTP(w, req)

	//saida
	assert.Equal(t, 200, w.Code)
}

// tdd deletar usuarios
func TestUsuarioDeleteHandle(t *testing.T) {
	// iniciando
	gin.SetMode(gin.TestMode)
	route := gin.Default()
	s := start.NewStart()
	handle := s.UsuarioStart()

	//requisicao
	req := httptest.NewRequest(http.MethodDelete, "/usuarios/:1", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	//roteamento
	route.DELETE("/usuarios/:id", handle.UsuarioDeleteHandle)
	route.ServeHTTP(w, req)

	//saida
	assert.Equal(t, http.StatusNoContent, w.Code)
}
