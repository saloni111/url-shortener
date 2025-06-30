package main

import (
	"log"
	"net/http"
	"os"
	"url-shortener/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin to release mode in production
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create Gin router
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "url-shortener",
			"version": "1.0.0",
			"mode":    "in-memory",
		})
	})

	// URL shortener endpoint
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	router.POST("/shorten", handlers.ShortenURLHandler)

	// Analytics endpoint (must come before redirect to avoid conflicts)
	router.GET("/analytics/:code", handlers.AnalyticsHandler)

	// Redirect endpoint (must come last)
	router.GET("/:code", handlers.RedirectHandler)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s (in-memory mode)", port)
	log.Fatal(router.Run(":" + port))
}

// getEnv gets environment variable with fallback
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
