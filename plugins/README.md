# NexusPanel Plugins

This directory contains official plugins for NexusPanel.

## Plugin Development Guide

### Plugin Structure

A NexusPanel plugin should follow this structure:

```
plugin-name/
├── plugin.yaml          # Plugin manifest
├── main.go             # Plugin entry point
├── handlers/           # HTTP handlers (optional)
├── services/           # Business logic (optional)
├── models/             # Data models (optional)
└── README.md           # Plugin documentation
```

### Plugin Manifest (plugin.yaml)

```yaml
name: "example-plugin"
version: "1.0.0"
description: "An example plugin for NexusPanel"
author: "Your Name"
homepage: "https://github.com/yourname/example-plugin"
license: "MIT"

# Plugin capabilities
capabilities:
  - monitoring
  - management

# Required permissions
permissions:
  - server:read
  - server:write

# Configuration schema
config_schema:
  - name: "api_key"
    type: "string"
    required: true
    description: "API key for external service"
  - name: "interval"
    type: "integer"
    required: false
    default: 60
    description: "Check interval in seconds"

# Plugin entry point
entry_point: "main.go"

# Minimum NexusPanel version required
min_version: "0.2.0"
```

### Plugin Interface

```go
package main

import (
    "github.com/2670044605/NexusPanel/internal/plugins"
)

type ExamplePlugin struct {
    plugins.BasePlugin
}

func (p *ExamplePlugin) Initialize(config map[string]interface{}) error {
    // Initialize plugin with configuration
    return nil
}

func (p *ExamplePlugin) Start() error {
    // Start plugin operations
    return nil
}

func (p *ExamplePlugin) Stop() error {
    // Stop plugin operations
    return nil
}

func (p *ExamplePlugin) GetInfo() plugins.PluginInfo {
    return plugins.PluginInfo{
        Name:        "example-plugin",
        Version:     "1.0.0",
        Description: "An example plugin",
        Author:      "Your Name",
    }
}

// Export plugin
var Plugin ExamplePlugin
```

### Plugin API

Plugins can access NexusPanel APIs through the plugin context:

```go
// Access database
db := ctx.GetDatabase()

// Access logger
logger := ctx.GetLogger()

// Access configuration
config := ctx.GetConfig()

// Register HTTP handlers
ctx.RegisterHandler("GET", "/api/plugins/example", handler)

// Emit events
ctx.EmitEvent("plugin.event", data)

// Subscribe to events
ctx.SubscribeEvent("server.added", eventHandler)
```

## Official Plugins (Coming Soon)

- **Docker Plugin**: Docker container management
- **Kubernetes Plugin**: Kubernetes cluster management
- **Backup Plugin**: Automated backup and restore
- **Monitoring Plus**: Advanced monitoring features
- **Notification Plugin**: Multiple notification channels
- **Database Plugin**: Database management tools

## Community Plugins

To submit a community plugin:

1. Create your plugin following the structure above
2. Test thoroughly with NexusPanel
3. Submit a PR to the plugin registry
4. Wait for review and approval

## Plugin Security

All plugins are sandboxed and have limited access to system resources. Plugins can only:

- Access APIs with granted permissions
- Read/write to their own data directory
- Make HTTP requests (if permission granted)
- Execute commands (if permission granted)

## Support

For plugin development support:
- Documentation: https://docs.nexuspanel.com/plugins
- Discussions: https://github.com/2670044605/NexusPanel/discussions
- Issues: https://github.com/2670044605/NexusPanel/issues
