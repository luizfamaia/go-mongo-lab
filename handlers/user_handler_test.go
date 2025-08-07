package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// simula um JSON enviado no corpo da requisição
	body := []byte(`{"name": "Luiz", "email": "luiz@example.com"}`)
	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r := gin.Default()
	// substitua por sua função real (dependendo da sua arquitetura final)
	r.POST("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Usuário inserido com sucesso"})
	})

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Usuário inserido com sucesso")
}

func TestUpdateUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	body := []byte(`{"name": "NovoNome", "email": "novo@email.com"}`)
	req, _ := http.NewRequest(http.MethodPut, "/users/123", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r := gin.Default()
	r.PUT("/users/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Usuário atualizado com sucesso"})
	})

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Usuário atualizado com sucesso")
}

func TestDeleteUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	req, _ := http.NewRequest(http.MethodDelete, "/users/123", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.DELETE("/users/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso"})
	})

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Usuário deletado com sucesso")
}
