package routes

import (
	"vaultex/internal/handlers"

	"github.com/gin-gonic/gin"
)

func WalletRoutes(r *gin.RouterGroup, handler *handlers.WalletHandler, authMiddleware gin.HandlerFunc) {
	group := r.Group("/wallets")
	group.Use(authMiddleware)
	{
		group.POST("", handler.CreateWallet)
		group.GET("/:external_user_id", handler.GetWalletByOrganizationIDAndExternalUserID)
	}
}
