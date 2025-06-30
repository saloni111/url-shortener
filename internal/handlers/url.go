package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// In-memory storage (fallback when database is not available)
var urlStore = make(map[string]string)
var analyticsStore = make(map[string]int)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RedirectHandler(c *gin.Context) {
    code := c.Param("code")

    // Hardcoded YouTube example
    if code == "5AwY56" {
        c.Redirect(http.StatusMovedPermanently, "https://www.youtube.com/watch?v=dQw4w9WgXcQ")
        return
    }

    // ...existing code for in-memory lookup...
}

// generateShortCode creates a random 6-character string
func generateShortCode() string {
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// ShortenURLRequest is the expected JSON payload
type ShortenURLRequest struct {
	URL string `json:"url" binding:"required,url"`
}

// ShortenURLHandler handles POST /shorten
func ShortenURLHandler(c *gin.Context) {
	var req ShortenURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Generate a unique short code
	var code string
	for {
		code = generateShortCode()
		if _, exists := urlStore[code]; !exists {
			break
		}
	}

	// Store in memory
	urlStore[code] = req.URL
	analyticsStore[code] = 0 // Initialize analytics counter

	c.JSON(http.StatusOK, gin.H{
		"short_url":    c.Request.Host + "/" + code,
		"code":         code,
		"original_url": req.URL,
	})
}

// RedirectHandler handles GET /:code
func RedirectHandler(c *gin.Context) {
	code := c.Param("code")

	originalURL, exists := urlStore[code]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	// Increment analytics counter
	analyticsStore[code]++

	// Redirect to the original URL
	c.Redirect(http.StatusMovedPermanently, originalURL)
}
