package middleware

import (
	"github.com/2670044605/NexusPanel/pkg/config"
	"github.com/gin-gonic/gin"
)

// CORS middleware for handling Cross-Origin Resource Sharing
func CORS(cfg config.CORSConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !cfg.Enabled {
			c.Next()
			return
		}

		origin := c.Request.Header.Get("Origin")

		// Set CORS headers
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Accept-Language")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Auth middleware for JWT authentication
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement JWT authentication
		// 1. Extract token from Authorization header
		// 2. Validate token
		// 3. Set user info in context
		// 4. Call c.Next() if valid, otherwise c.AbortWithStatusJSON()

		// For now, just pass through
		c.Next()
	}
}

// RequireAdmin middleware ensures user is an admin
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Check if user has admin role
		// Get user from context and verify role

		// For now, just pass through
		c.Next()
	}
}
