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

func TestUsuarioPostHandle(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handle := NewUsuarioHandle()

	msg := models.UsuarioCriate{
		Id:        0,
		Nome:      "Bruno",
		Email:     "brunopessoa@gmail.com",
		Senha:     "1234",
		Role:      1,
		Bloqueado: false,
	}

	body, _ := json.Marshal(&msg)

	req := httptest.NewRequest(http.MethodPost, "/usuarios", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.POST("/usuarios", handle.UsuarioPostHandle)
	router.ServeHTTP(w, req)

	var resp models.UsuarioCriate

	json.Unmarshal(w.Body.Bytes(), &resp)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), resp.Nome)
}
