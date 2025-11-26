package main

import (
	"fmt"
	"log"

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

	logger.Info("Starting NexusPanel Agent...")
	logger.Info("Agent functionality will be implemented in v0.2.0")

	// TODO: Implement agent functionality
	// - gRPC server for receiving commands from main server
	// - System metrics collection
	// - Command execution
	// - File operations
	// - Heartbeat mechanism

	logger.Info("Agent started successfully")
	
	// Keep agent running
	select {}
}

func printBanner() {
	banner := `
╔═══════════════════════════════════════════════════════════╗
║   NexusPanel Agent                                        ║
║   Version: 0.1.0                                          ║
╚═══════════════════════════════════════════════════════════╝
`
	fmt.Println(banner)
}
