package routes

import (
	"vaultex/internal/handlers"

	"github.com/gin-gonic/gin"
)

func OrganizationRoutes(r *gin.RouterGroup, handler *handlers.OrganizationHandler) {
	group := r.Group("/organizations")
	{
		group.POST("", handler.CreateOrganization)
		group.POST("/login", handler.Login)
		// group.GET("/:id", handler.FindOrganizationByAPIKey)
		group.GET("", handler.FindOrganizationByEmail)
	}
}
