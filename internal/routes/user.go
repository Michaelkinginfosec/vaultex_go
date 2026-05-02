package routes

import (
	"vaultex/internal/handlers"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(r *gin.RouterGroup, handler *handlers.UserHandler) {
	group := r.Group("/users")
	{
		group.POST("", handler.CreateUser)
		group.GET("/:id", handler.GetUser)
		group.GET("/user/:email", handler.FindUserByEmail)
		group.GET("", handler.GetAllUsers)
		group.PUT("/:id", handler.UpdateUser)
		group.DELETE("/:id", handler.DeleteUser)
	}
}
