package api

import (
	"github.com/2670044605/NexusPanel/internal/api/handlers"
	"github.com/2670044605/NexusPanel/internal/api/middleware"
	"github.com/2670044605/NexusPanel/pkg/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NewRouter creates and configures the API router
func NewRouter(cfg *config.Config, db *gorm.DB) *gin.Engine {
	// Set Gin mode
	gin.SetMode(cfg.Server.Mode)

	// Create router
	router := gin.New()

	// Global middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORS(cfg.Security.CORS))
	router.Use(middleware.I18n())

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"version": "0.1.0",
		})
	})

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Public routes (no authentication required)
		public := v1.Group("")
		{
			// Authentication
			auth := public.Group("/auth")
			{
				auth.POST("/register", handlers.Register)
				auth.POST("/login", handlers.Login)
				auth.POST("/refresh", handlers.RefreshToken)
			}
		}

		// Protected routes (authentication required)
		protected := v1.Group("")
		protected.Use(middleware.Auth())
		{
			// User profile
			protected.GET("/profile", handlers.GetProfile)
			protected.PUT("/profile", handlers.UpdateProfile)
			protected.POST("/logout", handlers.Logout)

			// Servers
			servers := protected.Group("/servers")
			{
				servers.GET("", handlers.ListServers)
				servers.POST("", handlers.CreateServer)
				servers.GET("/:id", handlers.GetServer)
				servers.PUT("/:id", handlers.UpdateServer)
				servers.DELETE("/:id", handlers.DeleteServer)
				servers.GET("/:id/status", handlers.GetServerStatus)
				servers.GET("/:id/metrics", handlers.GetServerMetrics)
			}

			// Server groups
			groups := protected.Group("/groups")
			{
				groups.GET("", handlers.ListServerGroups)
				groups.POST("", handlers.CreateServerGroup)
				groups.GET("/:id", handlers.GetServerGroup)
				groups.PUT("/:id", handlers.UpdateServerGroup)
				groups.DELETE("/:id", handlers.DeleteServerGroup)
			}

			// Monitoring
			monitoring := protected.Group("/monitoring")
			{
				monitoring.GET("/metrics", handlers.GetMetrics)
				monitoring.GET("/alerts", handlers.ListAlerts)
				monitoring.POST("/alerts/:id/acknowledge", handlers.AcknowledgeAlert)
				monitoring.POST("/alerts/:id/resolve", handlers.ResolveAlert)
			}

			// Alert rules
			alertRules := protected.Group("/alert-rules")
			{
				alertRules.GET("", handlers.ListAlertRules)
				alertRules.POST("", handlers.CreateAlertRule)
				alertRules.GET("/:id", handlers.GetAlertRule)
				alertRules.PUT("/:id", handlers.UpdateAlertRule)
				alertRules.DELETE("/:id", handlers.DeleteAlertRule)
			}

			// Tasks
			tasks := protected.Group("/tasks")
			{
				tasks.GET("", handlers.ListTasks)
				tasks.POST("", handlers.CreateTask)
				tasks.GET("/:id", handlers.GetTask)
				tasks.PUT("/:id", handlers.UpdateTask)
				tasks.DELETE("/:id", handlers.DeleteTask)
				tasks.POST("/:id/execute", handlers.ExecuteTask)
				tasks.GET("/:id/logs", handlers.GetTaskLogs)
			}

			// Plugins (v0.2.0)
			plugins := protected.Group("/plugins")
			{
				plugins.GET("", handlers.ListPlugins)
				plugins.POST("", handlers.InstallPlugin)
				plugins.GET("/:id", handlers.GetPlugin)
				plugins.PUT("/:id", handlers.UpdatePlugin)
				plugins.DELETE("/:id", handlers.UninstallPlugin)
				plugins.POST("/:id/enable", handlers.EnablePlugin)
				plugins.POST("/:id/disable", handlers.DisablePlugin)
			}

			// System settings (admin only)
			admin := protected.Group("/admin")
			admin.Use(middleware.RequireAdmin())
			{
				// Users management
				users := admin.Group("/users")
				{
					users.GET("", handlers.ListUsers)
					users.POST("", handlers.CreateUser)
					users.GET("/:id", handlers.GetUser)
					users.PUT("/:id", handlers.UpdateUser)
					users.DELETE("/:id", handlers.DeleteUser)
				}

				// System settings
				settings := admin.Group("/settings")
				{
					settings.GET("", handlers.GetSettings)
					settings.PUT("", handlers.UpdateSettings)
				}

				// Logs
				logs := admin.Group("/logs")
				{
					logs.GET("/users", handlers.GetUserLogs)
					logs.GET("/system", handlers.GetSystemLogs)
				}
			}
		}
	}

	// WebSocket routes
	ws := router.Group("/ws")
	ws.Use(middleware.Auth())
	{
		ws.GET("/terminal/:id", handlers.WebSocketTerminal)
		ws.GET("/metrics/:id", handlers.WebSocketMetrics)
	}

	return router
}
