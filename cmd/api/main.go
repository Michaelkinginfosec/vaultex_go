package main

import (
	"fmt"
	"log"
	_ "vaultex/docs"
	"vaultex/internal/handlers"
	"vaultex/internal/repository"
	"vaultex/internal/routes"
	"vaultex/internal/service"
	"vaultex/pkg/config"
	"vaultex/pkg/database"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Vaultex API
// @version 1.0
// @description This is the API documentation for Vaultex, a secure vault application.
// @contact.name API Support
// @contact.url https://github.com/michaelkinginfosec
// @contact.email osundemichael7@gmail.com
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
	userRepo := repository.NewRepository(pool)
	userService := service.NewService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
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
	routes.UsersRoutes(v1, userHandler)
	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Printf("Server running on :%s\n", cfg.Port)
	r.Run(":" + cfg.Port)

}
