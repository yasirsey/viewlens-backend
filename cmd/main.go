package main

import (
	"log"
	"viewlens/config"
	"viewlens/db"
	"viewlens/internal/controllers"
	"viewlens/internal/repositories"
	"viewlens/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to the database
	database := db.Connect(cfg)

	// Initialize repositories and services
	viewRepo := repositories.NewViewRepository(database)
	viewService := services.NewViewService(viewRepo)
	viewController := controllers.NewViewController(viewService)

	// Create Gin router
	r := gin.Default()

	// Routes
	r.GET("/views", viewController.ListViews)
	r.GET("/views/:viewName", viewController.GetViewDetails)

	// Start server
	log.Printf("Server running on port %s", cfg.AppPort)
	r.Run(":" + cfg.AppPort)
}
