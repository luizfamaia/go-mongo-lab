package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUserMock(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.PUT("/users/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Usuário atualizado com sucesso"})
	})

	body := `{"name":"Novo Nome", "email":"novo@email.com"}`
	req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Usuário atualizado com sucesso")
}

func TestDeleteUserMock(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.DELETE("/users/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso"})
	})

	req, _ := http.NewRequest("DELETE", "/users/1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Usuário deletado com sucesso")
}
