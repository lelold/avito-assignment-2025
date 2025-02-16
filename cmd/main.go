package main

import (
	"avito-assignment-2025/config"
	"avito-assignment-2025/database"
	"avito-assignment-2025/handler"
	"avito-assignment-2025/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
	log.Println("Config loaded successfully")

	database.InitDB(cfg)

	r := gin.Default()

	r.POST("/api/auth", handler.AuthHandler)

	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/sendCoin", handler.SendCoinHandler)
		auth.GET("/buy/:item", handler.BuyHandler)
		auth.GET("/info", handler.InfoHandler)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
