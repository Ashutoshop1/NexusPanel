package main

import (
	"fmt"
	"log"

	"github.com/2670044605/NexusPanel/internal/api"
	"github.com/2670044605/NexusPanel/internal/database"
	"github.com/2670044605/NexusPanel/pkg/config"
	"github.com/2670044605/NexusPanel/pkg/logger"
)

func main() {
	// Print banner
	printBanner()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	if err := logger.Init(cfg.Logging); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	logger.Info("Starting NexusPanel Server...")
	logger.Infof("Server mode: %s", cfg.Server.Mode)

	// Initialize database
	db, err := database.Init(cfg.Database)
	if err != nil {
		logger.Fatalf("Failed to initialize database: %v", err)
	}
	logger.Info("Database connected successfully")

	// Run auto migration if enabled
	if cfg.Database.AutoMigrate {
		if err := database.AutoMigrate(db); err != nil {
			logger.Fatalf("Failed to run database migrations: %v", err)
		}
		logger.Info("Database migrations completed")
	}

	// Initialize API router
	router := api.NewRouter(cfg, db)

	// Start server
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	logger.Infof("Server listening on %s", addr)
	
	if err := router.Run(addr); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}

func printBanner() {
	banner := `
╔═══════════════════════════════════════════════════════════╗
║                                                           ║
║   ███╗   ██╗███████╗██╗  ██╗██╗   ██╗███████╗           ║
║   ████╗  ██║██╔════╝╚██╗██╔╝██║   ██║██╔════╝           ║
║   ██╔██╗ ██║█████╗   ╚███╔╝ ██║   ██║███████╗           ║
║   ██║╚██╗██║██╔══╝   ██╔██╗ ██║   ██║╚════██║           ║
║   ██║ ╚████║███████╗██╔╝ ██╗╚██████╔╝███████║           ║
║   ╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚══════╝           ║
║                                                           ║
║   ██████╗  █████╗ ███╗   ██╗███████╗██╗                ║
║   ██╔══██╗██╔══██╗████╗  ██║██╔════╝██║                ║
║   ██████╔╝███████║██╔██╗ ██║█████╗  ██║                ║
║   ██╔═══╝ ██╔══██║██║╚██╗██║██╔══╝  ██║                ║
║   ██║     ██║  ██║██║ ╚████║███████╗███████╗           ║
║   ╚═╝     ╚═╝  ╚═╝╚═╝  ╚═══╝╚══════╝╚══════╝           ║
║                                                           ║
║   Next-Gen Server Management Platform                     ║
║   Version: 0.1.0                                          ║
║   License: AGPL-3.0                                       ║
║                                                           ║
╚═══════════════════════════════════════════════════════════╝
`
	fmt.Println(banner)
}
