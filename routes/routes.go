package routes

import (
	"net/http"

	"go-mongo-lab/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handlers.UserHandler) *gin.Engine {
	r := gin.Default()

	// Health (liveness/readiness simples)
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Versão 1 da API
	v1 := r.Group("/v1")
	{
		v1.POST("/users", h.CreateUser)
		v1.PUT("/users/:id", h.UpdateUser)
		v1.DELETE("/users/:id", h.DeleteUser)
	}

	return r
}
