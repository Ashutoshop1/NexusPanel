package middleware

import (
	"github.com/gin-gonic/gin"
)

// I18nMiddleware handles internationalization
func I18nMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get language from various sources
		lang := c.GetHeader("Accept-Language")
		if lang == "" {
			lang = c.Query("lang")
		}
		if lang == "" {
			// Get from user preferences if authenticated
			// TODO: Get from user model
			lang = "zh-CN"
		}

		// Normalize language code
		switch lang {
		case "zh", "zh-CN", "zh-Hans":
			lang = "zh-CN"
		case "en", "en-US":
			lang = "en-US"
		default:
			lang = "zh-CN"
		}

		// Set in context
		c.Set("lang", lang)
		c.Next()
	}
}
