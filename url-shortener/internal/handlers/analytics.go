package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AnalyticsHandler handles GET /analytics/:code
func AnalyticsHandler(c *gin.Context) {
	code := c.Param("code")

	clicks, exists := analyticsStore[code]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	originalURL, exists := urlStore[code]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":         code,
		"original_url": originalURL,
		"clicks":       clicks,
		"short_url":    c.Request.Host + "/" + code,
	})
}
