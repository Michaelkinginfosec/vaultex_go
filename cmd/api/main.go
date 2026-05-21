package main

import (
	"fmt"
	"log"
	_ "vaultex/docs"
	"vaultex/internal/handlers"
	"vaultex/internal/middleware"
	"vaultex/internal/repository"
	"vaultex/internal/routes"
	"vaultex/internal/service"
	"vaultex/pkg/config"
	"vaultex/pkg/database"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Vaultex
// @version 1.0
// @description A Ledger-as-a-Service API. Platforms integrate Vaultex to manage their users' wallets and track every financial transaction via double-entry bookkeeping.
// @contact.name API Support
// @contact.url https://github.com/michaelkinginfosec
// @contact.email osundemichael7@gmail.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-api-key
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}
	pool, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}

	organizationRepo := repository.NewRepository(pool)
	walletRepo := repository.NewWalletRepository(pool)
	organizationService := service.NewService(organizationRepo)
	walletService := service.NewWalletService(walletRepo)
	organizationHandler := handlers.NewOrganizationHandler(organizationService)
	walletHandler := handlers.NewWalletHandler(walletService)
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message":  "Welcome to Vaultex API",
			"status":   "success",
			"database": "Connected",
		})
	})
	api := r.Group("/api")
	v1 := api.Group("/v1")
	authMiddleware := middleware.AuthMiddleware(pool)
	routes.OrganizationRoutes(v1, organizationHandler)
	routes.WalletRoutes(v1, walletHandler, authMiddleware)
	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Printf("Server running on :%s\n", cfg.Port)
	r.Run(":" + cfg.Port)

}
