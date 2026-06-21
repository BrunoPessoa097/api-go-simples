package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
	"github.com/BrunoPessoa097/api-go-simples/internal/start"
	"github.com/BrunoPessoa097/api-go-simples/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/testify/v2/assert"
)

// tdd lista de usuarios
func TestUsuarioListHandle(t *testing.T) {
	// inicializando
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	s := start.NewStart(utils.SetupDB(t))
	handle, _ := s.UsuarioStart()

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

	s := start.NewStart(utils.SetupDB(t))
	handle, roles := s.UsuarioStart()

	// entrada
	msg := models.Usuario{
		Nome:      "Bruno Frefre",
		Email:     "brunopessoa1234@gmail.com",
		Senha:     "12345678",
		Role:      1,
		Bloqueado: false,
	}
	role1 := models.Roles{
		Nivel: "venda",
		Regra: "get,post",
	}
	roles.RolesRepositoryAdd(&role1)

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
	var resp models.Usuario
	json.Unmarshal(w.Body.Bytes(), &resp)

	//saida
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), resp.Nome)
}

// tdd pegar usuario por id
func TestUsuarioByIdHandle(t *testing.T) {
	//iniciar
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	s := start.NewStart(utils.SetupDB(t))
	handler, roles := s.UsuarioStart()

	role1 := models.Roles{
		Nivel: "venda",
		Regra: "get,post",
	}
	roles.RolesRepositoryAdd(&role1)

	//requesicao
	req := httptest.NewRequest(http.MethodGet, "/usuarios/1", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// rota leitura e escrita
	router.GET("/usuarios/:id", handler.UsuarioByIdHandle)
	router.ServeHTTP(w, req)

	// recebendo para converter para json
	var resp models.Usuario
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
	db := utils.SetupDB(t)
	s := start.NewStart(db)
	handler, roles := s.UsuarioStart()

	// entrada
	msg := models.Usuario{
		Nome:      "Bruno vv",
		Email:     "ba122@gmail.com",
		Senha:     "12345678",
		Role:      1,
		Bloqueado: false,
	}

	role1 := models.Roles{
		Nivel: "venda",
		Regra: "get,post",
	}
	roles.RolesRepositoryAdd(&role1)

	msg2 := models.Usuario{
		ID:        1,
		Nome:      "Bruno Frefre",
		Email:     "brunopessoa1234@gmail.com",
		Senha:     "12345678",
		Role:      1,
		Bloqueado: false,
	}

	repo := repository.NewUsuarioRepository(db)
	repo.UsuarioRepositoryAdd(msg2)

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
	db := utils.SetupDB(t)
	s := start.NewStart(db)
	handle, roles := s.UsuarioStart()

	repo := repository.NewUsuarioRepository(db)

	msg := models.Usuario{
		Nome:      "Bruno Frefre",
		Email:     "brunopessoa1234@gmail.com",
		Senha:     "12345678",
		Role:      1,
		Bloqueado: false,
	}
	role1 := models.Roles{
		Nivel: "venda",
		Regra: "get,post",
	}

	roles.RolesRepositoryAdd(&role1)
	repo.UsuarioRepositoryAdd(msg)
	//requisicao
	req := httptest.NewRequest(http.MethodDelete, "/usuarios/1", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	//roteamento
	route.DELETE("/usuarios/:id", handle.UsuarioDeleteHandle)
	route.ServeHTTP(w, req)

	//saida
	assert.Equal(t, http.StatusNoContent, w.Code)
}
