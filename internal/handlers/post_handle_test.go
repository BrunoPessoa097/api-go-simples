package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BrunoPessoa097/api-go-simples/internal/models"
	"github.com/BrunoPessoa097/api-go-simples/internal/repository"
	"github.com/BrunoPessoa097/api-go-simples/internal/services"
	"github.com/BrunoPessoa097/api-go-simples/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/testify/v2/assert"
)

// func inicial() (*gin.Engine, *PostHandlers) {
// 	gin.SetMode(gin.TestMode)
// 	r := gin.Default()
// 	mocks := mocks.ListPost
// 	rp := repository.NewPostRepository(mocks)
// 	sp := services.NewPostService(rp)
// 	return r, NewPostHandlers(sp)
// }

// teste listagem
func TestPostHandlerList(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	db := utils.SetupDB(t)
	pr := repository.NewPostRepository(db)
	ru := repository.NewUsuarioRepository(db)
	ps := services.NewPostService(pr, ru)
	p := NewPostHandlers(ps)

	req := httptest.NewRequest(http.MethodGet, "/posts", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.GET("/posts", p.PostHandlersList)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// inserir postagem
func TestPostHandlerPost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	db := utils.SetupDB(t)
	pr := repository.NewPostRepository(db)
	ru := repository.NewUsuarioRepository(db)
	ps := services.NewPostService(pr, ru)
	p := NewPostHandlers(ps)

	post := models.Post{
		IDUser: 11,
		Texto:  "Vamos nessa Brasil",
	}

	body, _ := json.Marshal(&post)

	req := httptest.NewRequest(http.MethodPost, "/posts", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.POST("/posts", p.PostHandlersPost)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// selecionar postagem
func TestPostHandlerById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	db := utils.SetupDB(t)
	pr := repository.NewPostRepository(db)
	ru := repository.NewUsuarioRepository(db)
	ps := services.NewPostService(pr, ru)
	rr := repository.NewRolesRepository(db)
	p := NewPostHandlers(ps)

	role1 := models.Roles{
		Nivel: "venda",
		Regra: "get,post",
	}
	rr.RolesRepositoryAdd(&role1)

	user := models.Usuario{
		Nome:      "Bruno Frefre",
		Email:     "brunopessoa1234@gmail.com",
		Senha:     "12345678",
		Role:      role1.ID,
		Bloqueado: false,
	}
	ru.UsuarioRepositoryAdd(&user)

	post := models.Post{
		IDUser: int64(user.ID),
		Texto:  "Vamos nessa Brasil",
	}

	p.Service.PostServiceAdd(&post)

	req := httptest.NewRequest(http.MethodGet, "/posts/1", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.GET("/posts/:id", p.PostHandlersById)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// atualizar postagem
func TestPostHandlerUpdate(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	db := utils.SetupDB(t)
	pr := repository.NewPostRepository(db)
	ru := repository.NewUsuarioRepository(db)
	ps := services.NewPostService(pr, ru)
	rr := repository.NewRolesRepository(db)
	p := NewPostHandlers(ps)

	role1 := models.Roles{
		Nivel: "venda",
		Regra: "get,post",
	}
	rr.RolesRepositoryAdd(&role1)

	user := models.Usuario{
		Nome:      "Bruno Frefre",
		Email:     "brunopessoa1234@gmail.com",
		Senha:     "12345678",
		Role:      role1.ID,
		Bloqueado: false,
	}
	ru.UsuarioRepositoryAdd(&user)

	post := models.Post{
		IDUser: int64(user.ID),
		Texto:  "Vamos nessa Brasil",
	}

	p.Service.PostServiceAdd(&post)

	postUpdate := models.Post{
		ID:     1,
		IDUser: 12,
		Texto:  "Bem Vindo ao Marrocos",
	}

	update, _ := json.Marshal(&postUpdate)

	req := httptest.NewRequest(http.MethodPatch, "/posts/1", bytes.NewBuffer(update))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.PATCH("/posts/:id", p.PostHandlersUpdate)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// delete de postagem
func TestPostHandlerDelete(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	db := utils.SetupDB(t)
	pr := repository.NewPostRepository(db)
	ru := repository.NewUsuarioRepository(db)
	ps := services.NewPostService(pr, ru)
	rr := repository.NewRolesRepository(db)
	p := NewPostHandlers(ps)

	role1 := models.Roles{
		Nivel: "venda",
		Regra: "get,post",
	}
	rr.RolesRepositoryAdd(&role1)

	user := models.Usuario{
		Nome:      "Bruno Frefre",
		Email:     "brunopessoa1234@gmail.com",
		Senha:     "12345678",
		Role:      role1.ID,
		Bloqueado: false,
	}
	ru.UsuarioRepositoryAdd(&user)

	post := models.Post{
		IDUser: int64(user.ID),
		Texto:  "Vamos nessa Brasil",
	}

	p.Service.PostServiceAdd(&post)

	req := httptest.NewRequest(http.MethodDelete, "/posts/1", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.DELETE("/posts/:id", p.PostHandlersDelete)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
