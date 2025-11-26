<div align="center">

# 🚀 NexusPanel

### Modular, Cloud-Native, AI-Powered Next-Generation Server Management Platform

[中文文档](./README.zh-CN.md) | [Documentation](./docs/en-US/) | [Contributing](#contributing) | [License](#license)

[![CI](https://github.com/2670044605/NexusPanel/workflows/CI/badge.svg)](https://github.com/2670044605/NexusPanel/actions)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![Vue Version](https://img.shields.io/badge/Vue-3.x-4FC08D?style=flat&logo=vue.js)](https://vuejs.org)
[![License](https://img.shields.io/badge/License-AGPL--3.0-blue.svg)](./LICENSE)

</div>

---

## 📖 Introduction

**NexusPanel** is a modern, modular server management platform designed for DevOps engineers, developers, and SMEs. Built with a cloud-native architecture and powered by AI, it provides comprehensive server monitoring, management, and automation capabilities.

### Why NexusPanel?

- 🎯 **Unified Management**: Manage all your servers from a single, intuitive interface
- 🧩 **Modular Design**: Extensible plugin system for custom functionality
- 🤖 **AI-Powered**: Intelligent diagnostics and automated problem resolution
- 🌐 **Cloud-Native**: Built for modern infrastructure with Docker and Kubernetes support
- 🔒 **Security First**: RBAC, JWT authentication, and comprehensive audit logging
- 🌍 **Multi-Language**: Full support for Chinese and English

---

## ✨ Features

### MVP Phase (v0.1.0)
- ✅ User authentication system (Login/Register/JWT/RBAC)
- ✅ Server management (Add/Edit/Delete/Group)
- ✅ Real-time monitoring (CPU/Memory/Disk/Network)
- ✅ Web-based terminal (WebSocket SSH)
- ✅ File manager with upload/download
- ✅ Basic alerting system
- ✅ Multi-language support (Chinese/English)

### Extended Phase (v0.2.0)
- 🔄 Plugin system architecture
- 🔄 Agent management module
- 🔄 Container management (Docker)
- 🔄 Scheduled tasks
- 🔄 Automatic backup

### Advanced Phase (v1.0.0)
- 🔮 AI diagnostic assistant
- 🔮 Plugin marketplace
- 🔮 Multi-tenancy support
- 🔮 Kubernetes management
- 🔮 Automated workflows

---

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher
- Node.js 18 or higher
- PostgreSQL 13+ (for production) or SQLite (for development)
- Redis (optional, for caching)

### Installation

#### 1. Clone the repository

```bash
git clone https://github.com/2670044605/NexusPanel.git
cd NexusPanel
```

#### 2. Backend setup

```bash
# Install Go dependencies
go mod download

# Copy configuration file
cp configs/config.example.yaml configs/config.yaml

# Edit configuration as needed
vim configs/config.yaml

# Build the application
make build

# Run database migrations
make migrate-up

# Start the server
make run
```

#### 3. Frontend setup

```bash
# Navigate to web directory
cd web

# Install dependencies
npm install

# Start development server
npm run dev
```

#### 4. Access the application

Open your browser and navigate to: `http://localhost:5173`

Default credentials:
- Username: `admin`
- Password: `admin123`

**⚠️ Important: Change the default password after first login!**

---

## 🛠️ Tech Stack

### Backend
- **Language**: Go 1.21+
- **Web Framework**: Gin
- **ORM**: GORM
- **Database**: PostgreSQL / SQLite
- **Cache**: Redis (optional)
- **Authentication**: JWT
- **API**: RESTful + WebSocket + gRPC
- **Logging**: Zap
- **Configuration**: Viper
- **Internationalization**: go-i18n

### Frontend
- **Framework**: Vue 3
- **Language**: TypeScript
- **Build Tool**: Vite
- **UI Library**: TailwindCSS
- **State Management**: Pinia
- **HTTP Client**: Axios
- **Internationalization**: vue-i18n
- **Icons**: Heroicons

### Infrastructure
- **Containerization**: Docker
- **Orchestration**: Docker Compose / Kubernetes
- **CI/CD**: GitHub Actions

---

## 📁 Project Structure

```
nexuspanel/
├── cmd/                          # Application entry points
│   ├── server/                   # Main server
│   └── agent/                    # Agent
├── internal/                     # Private application code
│   ├── core/                     # Core business logic
│   ├── plugins/                  # Plugin engine
│   ├── ai/                       # AI module
│   ├── database/                 # Database operations
│   ├── i18n/                     # Internationalization
│   └── api/                      # API layer
├── pkg/                          # Public reusable code
│   ├── config/                   # Configuration management
│   ├── logger/                   # Logging utilities
│   ├── utils/                    # Utility functions
│   └── crypto/                   # Cryptography utilities
├── web/                          # Frontend application
│   ├── src/                      # Source code
│   └── public/                   # Static assets
├── plugins/                      # Official plugins
├── deploy/                       # Deployment files
│   ├── docker/                   # Docker files
│   └── scripts/                  # Installation scripts
├── docs/                         # Documentation
│   ├── zh-CN/                    # Chinese docs
│   └── en-US/                    # English docs
└── configs/                      # Configuration files
```

---

## 💻 Development Guide

### Development Environment

```bash
# Install development tools
make install-tools

# Run in development mode (backend + frontend)
make dev

# Run tests
make test

# Run linters
make lint

# Format code
make fmt
```

### Building

```bash
# Build all binaries
make build

# Build server only
make build-server

# Build agent only
make build-agent

# Build frontend
make web-build
```

### Docker

```bash
# Build Docker image
make docker

# Start with docker-compose
make docker-compose-up

# Stop docker-compose
make docker-compose-down
```

### Database Migrations

```bash
# Run migrations
make migrate-up

# Rollback migrations
make migrate-down
```

---

## 📚 Documentation

- [Product Requirements Document (PRD)](./docs/en-US/PRD.md)
- [Architecture Documentation](./docs/en-US/ARCHITECTURE.md)
- [Database Design](./docs/en-US/DATABASE.md)
- [API Documentation](./docs/en-US/API.md)

---

## 🤝 Contributing

We welcome contributions from the community! Please read our [Contributing Guidelines](./CONTRIBUTING.md) before submitting PRs.

### How to contribute:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### Development Guidelines:

- Follow Go coding standards
- Write tests for new features
- Update documentation
- Ensure CI passes
- Add meaningful commit messages

---

## 📄 License

This project is licensed under the **GNU Affero General Public License v3.0 (AGPL-3.0)**.

See [LICENSE](./LICENSE) file for details.

---

## 🙏 Acknowledgments

Special thanks to all contributors and the open-source community for their invaluable support.

---

## 📬 Contact

- **Issues**: [GitHub Issues](https://github.com/2670044605/NexusPanel/issues)
- **Discussions**: [GitHub Discussions](https://github.com/2670044605/NexusPanel/discussions)
- **Email**: support@nexuspanel.com

---

<div align="center">

**Made with ❤️ by the NexusPanel Team**

[⭐ Star us on GitHub](https://github.com/2670044605/NexusPanel) | [🐛 Report Bug](https://github.com/2670044605/NexusPanel/issues) | [💡 Request Feature](https://github.com/2670044605/NexusPanel/issues)

</div>
