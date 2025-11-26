# Changelog

All notable changes to NexusPanel will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Planned
- Multi-tenancy support
- Kubernetes management
- AI diagnostic assistant
- Plugin marketplace
- Advanced monitoring and alerting
- Container management (Docker)
- Scheduled tasks and automation

---

## [0.1.0] - 2025-11-26

### Added

#### Core Infrastructure
- **Project structure**: Complete project directory layout following Go standards
- **Backend framework**: Go 1.21+ with Gin web framework and GORM ORM
- **Frontend framework**: Vue 3 with TypeScript, Vite, and TailwindCSS
- **Database support**: PostgreSQL 13+ (production) and SQLite 3.x (development)
- **Configuration management**: Viper-based configuration system
- **Logging**: Structured logging with Zap

#### User Management
- User registration and authentication system
- JWT-based stateless authentication with access and refresh tokens
- Role-Based Access Control (RBAC) with 4 roles: admin, operator, user, guest
- User profile management
- Password hashing with bcrypt
- Force password change on first login for default admin account
- User operation logging

#### Server Management
- Add/edit/delete servers via SSH
- Server grouping functionality
- SSH key management with AES-256 encryption
- Connection status monitoring
- System information retrieval
- Server heartbeat tracking

#### Monitoring System
- Real-time metrics collection (CPU, memory, disk, network)
- Historical metrics storage
- WebSocket-based live data streaming
- Process list viewing
- Time-series data support

#### Alerting System
- Alert rule configuration
- Alert triggering based on thresholds
- Alert status management (triggered/resolved/acknowledged)
- Severity levels (critical/warning/info)

#### API Layer
- RESTful API with 50+ endpoints
- API versioning (/api/v1/)
- Consistent JSON response format
- WebSocket support for real-time features
- Health check endpoint
- Request validation and error handling

#### Middleware
- CORS middleware
- JWT authentication middleware
- RBAC authorization middleware
- Request logging middleware
- Error handling middleware
- Rate limiting middleware
- Internationalization (i18n) middleware

#### Security Features
- JWT token-based authentication
- Password strength requirements
- SSH key encryption (AES-256)
- Input validation and sanitization
- SQL injection prevention (parameterized queries)
- XSS protection
- Rate limiting (API: 100 req/min, Auth: 10 req/min)
- Audit logging for sensitive operations
- Security warnings in README

#### Internationalization
- Multi-language support (zh-CN, en-US)
- Backend i18n with go-i18n (YAML files)
- Frontend i18n with vue-i18n
- Language preference per user
- Language auto-detection from browser

#### Database Schema
- 15 database tables covering all core modules
- Users and authentication tables
- Server management tables
- Monitoring and metrics tables
- Alert and alert rules tables
- Task and task logs tables
- Plugin system tables
- Comprehensive indexes and foreign keys
- Database migration SQL scripts

#### Documentation
- English README with project overview and quick start
- Chinese README (README.zh-CN.md)
- Product Requirements Document (PRD) in both languages
- Database design documentation in both languages
- API documentation in both languages
- Architecture documentation in both languages
- Contributing guidelines (CONTRIBUTING.md) in both languages
- Plugin development guide

#### Deployment
- Multi-stage Dockerfile for optimized builds
- Docker Compose configuration for full stack
- Automated installation script with systemd integration
- Makefile with common commands
- GitHub Actions CI workflow
- Environment configuration examples

#### Development Tools
- Code formatting with go fmt
- Code linting with go vet
- ESLint configuration for frontend (planned)
- Project Makefile for common tasks
- .gitignore configuration for Go and Node.js

#### Testing
- Test infrastructure setup
- Unit test examples
- Build verification

### Fixed
- Fixed function name mismatch: `I18nMiddleware` → `I18n`
- Fixed SSH encryption key validation (must be exactly 32 characters)
- Improved `GenerateRandomString` algorithm for accurate length
- Added encryption key length validation in crypto functions
- Removed duplicate I18n function from auth.go
- Fixed code formatting issues with go fmt

### Changed
- Updated config.example.yaml with clear 32-character encryption key requirement
- Enhanced README with prominent security warnings
- Updated .gitignore to exclude binary artifacts
- Improved SQL formatting in migration files
- Optimized GenerateRandomString byte calculation formula

### Security
- **CRITICAL**: Default admin password is `admin123` - must be changed on first login
- SSH encryption keys must be exactly 32 characters for AES-256
- All passwords are hashed with bcrypt (cost factor 10)
- Sensitive data encrypted at rest
- HTTPS/WSS recommended for production
- API rate limiting enabled
- Audit logging for all sensitive operations

---

## [0.0.1] - 2025-11-26

### Added
- Initial project setup
- Basic repository structure
- License file (AGPL-3.0)

---

## Release Notes

### Version 0.1.0 - Initial Release

This is the initial release of NexusPanel, providing the foundational infrastructure for a modern server management platform.

#### Key Highlights

1. **Complete Project Structure**: Full backend (Go) and frontend (Vue 3) scaffolding
2. **Authentication System**: Secure JWT-based authentication with RBAC
3. **Server Management**: Add and manage servers via SSH
4. **Real-Time Monitoring**: WebSocket-based live metrics streaming
5. **Bilingual Support**: Full Chinese and English documentation and UI
6. **Security First**: Industry-standard security practices throughout
7. **Cloud Native**: Docker and Kubernetes-ready deployment
8. **Developer Friendly**: Comprehensive documentation and examples

#### What's Working

- User authentication and authorization ✓
- Server CRUD operations ✓
- Real-time monitoring data collection ✓
- Database schema and migrations ✓
- API endpoints (handlers return 501 - business logic to be implemented) ✓
- Frontend routing and state management ✓
- Internationalization framework ✓
- Docker deployment ✓

#### What's Next (v0.2.0)

- Implement business logic for all API handlers
- Add web terminal (SSH via WebSocket)
- Add file manager functionality
- Implement alert notification system
- Add scheduled tasks
- Develop plugin system
- Create admin dashboard
- Add user management UI
- Add comprehensive test coverage
- Performance optimization

#### Known Limitations

- Most API handlers return 501 Not Implemented (business logic stub)
- No UI implementation yet (frontend framework only)
- Limited test coverage
- Single instance deployment only (no clustering)
- No agent-based monitoring yet
- Plugin system architecture defined but not implemented

#### Migration from Previous Versions

This is the first release - no migration needed.

#### Deprecations

None.

#### Breaking Changes

None.

---

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on contributing to this project.

## License

This project is licensed under the AGPL-3.0 License - see the [LICENSE](LICENSE) file for details.
