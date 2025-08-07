package routes

import (
	"go-mongo-lab/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handlers.UserHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/users", h.CreateUser)
	r.PUT("/users/:id", h.UpdateUser)
	r.DELETE("/users/:id", h.DeleteUser)

	return r
}
