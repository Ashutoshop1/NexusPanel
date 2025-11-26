# NexusPanel Architecture Document

## Document Information

- **Version**: v1.0
- **Status**: Active
- **Last Updated**: 2024
- **Authors**: NexusPanel Team

---

## Table of Contents

1. [System Overview](#system-overview)
2. [Architecture Principles](#architecture-principles)
3. [System Architecture](#system-architecture)
4. [Component Design](#component-design)
5. [Data Flow](#data-flow)
6. [Security Architecture](#security-architecture)
7. [Deployment Architecture](#deployment-architecture)
8. [Scalability & Performance](#scalability--performance)

---

## System Overview

### Purpose

NexusPanel is a modern, cloud-native server management platform designed to provide unified management capabilities for servers, containers, and cloud resources through an intuitive web interface.

### Key Characteristics

- **Modular**: Plugin-based extensibility
- **Cloud-Native**: Containerized deployment, Kubernetes-ready
- **Multi-Tenant**: Support for multiple organizations (future)
- **Real-Time**: WebSocket-based live monitoring
- **Internationalized**: Multi-language support (zh-CN, en-US)
- **Secure**: Industry-standard security practices

### Technology Stack

#### Backend
- **Language**: Go 1.21+
- **Web Framework**: Gin
- **ORM**: GORM
- **Database**: PostgreSQL 13+ / SQLite 3.x
- **Cache**: Redis (optional)
- **Configuration**: Viper
- **Logging**: Zap
- **Authentication**: JWT

#### Frontend
- **Framework**: Vue 3 (Composition API)
- **Language**: TypeScript
- **Build Tool**: Vite
- **Styling**: TailwindCSS
- **State Management**: Pinia
- **Router**: Vue Router
- **i18n**: vue-i18n

---

## Architecture Principles

### 1. Separation of Concerns

The system is divided into distinct layers with clear responsibilities:
- **Presentation Layer**: UI components and user interaction
- **API Layer**: RESTful endpoints and WebSocket handlers
- **Business Logic Layer**: Core application logic
- **Data Access Layer**: Database operations and models
- **Infrastructure Layer**: External services and utilities

### 2. Modularity

- Core functionality in `internal/` packages
- Reusable utilities in `pkg/` packages
- Plugin system for extensibility
- Clear module boundaries

### 3. Statelessness

- JWT-based authentication (stateless)
- Horizontal scalability
- Session data in database/cache
- No server-side session storage

### 4. API-First Design

- Well-defined RESTful APIs
- OpenAPI/Swagger documentation
- API versioning (`/api/v1/`)
- Consistent response format

### 5. Security by Default

- HTTPS in production
- Input validation and sanitization
- SQL injection prevention
- XSS protection
- CSRF protection
- Rate limiting
- Audit logging

---

## System Architecture

### High-Level Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                        Client Layer                          │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐     │
│  │ Web Browser  │  │  Mobile App  │  │   CLI Tool   │     │
│  │  (Vue 3)     │  │   (Future)   │  │   (Future)   │     │
│  └──────┬───────┘  └──────┬───────┘  └──────┬───────┘     │
└─────────┼──────────────────┼──────────────────┼─────────────┘
          │                  │                  │
          │ HTTPS/WSS        │ HTTPS            │ HTTPS
          ▼                  ▼                  ▼
┌─────────────────────────────────────────────────────────────┐
│                    Reverse Proxy (Nginx)                     │
│              ┌────────────────────────────────┐              │
│              │    SSL/TLS Termination         │              │
│              │    Load Balancing               │              │
│              │    Rate Limiting                │              │
│              └────────────────────────────────┘              │
└────────────────────────────┬────────────────────────────────┘
                             │
                             ▼
┌─────────────────────────────────────────────────────────────┐
│                      API Gateway Layer                       │
│              (Gin Framework + Middleware)                    │
│  ┌──────────────────────────────────────────────────────┐  │
│  │ Middleware Chain:                                     │  │
│  │  • CORS                                               │  │
│  │  • Authentication (JWT)                               │  │
│  │  • Authorization (RBAC)                               │  │
│  │  • Rate Limiting                                      │  │
│  │  • Request Logging                                    │  │
│  │  • Error Handling                                     │  │
│  │  • i18n                                               │  │
│  └──────────────────────────────────────────────────────┘  │
└────────────────────────────┬────────────────────────────────┘
                             │
          ┌──────────────────┼──────────────────┐
          │                  │                  │
          ▼                  ▼                  ▼
┌──────────────────┐ ┌──────────────────┐ ┌──────────────────┐
│   REST Handler   │ │ WebSocket Handler│ │  gRPC Handler    │
│                  │ │                  │ │   (Future)       │
│  • User API      │ │  • Metrics Stream│ │                  │
│  • Server API    │ │  • Terminal      │ │  • Agent Comm    │
│  • Monitor API   │ │  • Events        │ │                  │
└────────┬─────────┘ └────────┬─────────┘ └────────┬─────────┘
         │                    │                     │
         └────────────────────┼─────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                   Business Logic Layer                       │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐          │
│  │    Auth     │ │   Server    │ │   Monitor   │          │
│  │   Module    │ │   Module    │ │   Module    │          │
│  └─────────────┘ └─────────────┘ └─────────────┘          │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐          │
│  │    User     │ │    Task     │ │   Plugin    │          │
│  │   Module    │ │   Module    │ │   Engine    │          │
│  └─────────────┘ └─────────────┘ └─────────────┘          │
└────────────────────────────┬────────────────────────────────┘
                             │
                             ▼
┌─────────────────────────────────────────────────────────────┐
│                    Data Access Layer                         │
│                     (GORM + Models)                          │
│  ┌─────────────────────────────────────────────────────┐   │
│  │ Database Models:                                     │   │
│  │  • User, Role, Token                                │   │
│  │  • Server, ServerGroup, SSHKey                      │   │
│  │  • Metric, Alert, AlertRule                         │   │
│  │  • Task, TaskLog                                    │   │
│  │  • Plugin, PluginSetting                            │   │
│  └─────────────────────────────────────────────────────┘   │
└────────────────────────────┬────────────────────────────────┘
                             │
          ┌──────────────────┼──────────────────┐
          │                  │                  │
          ▼                  ▼                  ▼
┌──────────────────┐ ┌──────────────────┐ ┌──────────────────┐
│   PostgreSQL     │ │      Redis       │ │     File         │
│   (Primary DB)   │ │     (Cache)      │ │     Storage      │
│                  │ │                  │ │                  │
│  • User data     │ │  • Sessions      │ │  • Uploads       │
│  • Server info   │ │  • Temp data     │ │  • Logs          │
│  • Metrics       │ │  • Rate limits   │ │  • Backups       │
└──────────────────┘ └──────────────────┘ └──────────────────┘

┌─────────────────────────────────────────────────────────────┐
│                    External Systems                          │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐          │
│  │   Target    │ │    Email    │ │    Cloud    │          │
│  │   Servers   │ │   Service   │ │   Storage   │          │
│  │   (SSH)     │ │   (SMTP)    │ │  (S3/OSS)   │          │
│  └─────────────┘ └─────────────┘ └─────────────┘          │
└─────────────────────────────────────────────────────────────┘
```

---

## Component Design

### 1. Frontend Architecture (Vue 3)

```
web/
├── src/
│   ├── components/          # Reusable UI components
│   │   ├── common/         # Common components (Button, Input)
│   │   ├── charts/         # Chart components
│   │   └── forms/          # Form components
│   │
│   ├── views/              # Page components
│   │   ├── LoginView.vue
│   │   ├── DashboardView.vue
│   │   ├── ServersView.vue
│   │   └── MonitoringView.vue
│   │
│   ├── stores/             # Pinia stores (state management)
│   │   ├── auth.ts         # Authentication state
│   │   ├── servers.ts      # Server management state
│   │   └── monitoring.ts   # Monitoring state
│   │
│   ├── router/             # Vue Router configuration
│   │   └── index.ts        # Route definitions
│   │
│   ├── api/                # API client
│   │   ├── client.ts       # Axios instance
│   │   ├── auth.ts         # Auth API calls
│   │   ├── servers.ts      # Server API calls
│   │   └── monitoring.ts   # Monitoring API calls
│   │
│   ├── i18n/               # Internationalization
│   │   ├── index.ts        # i18n configuration
│   │   ├── zh-CN.ts        # Chinese translations
│   │   └── en-US.ts        # English translations
│   │
│   ├── utils/              # Utility functions
│   │   ├── format.ts       # Data formatting
│   │   ├── validation.ts   # Input validation
│   │   └── websocket.ts    # WebSocket helper
│   │
│   ├── composables/        # Vue composables
│   │   ├── useAuth.ts      # Authentication logic
│   │   ├── useWebSocket.ts # WebSocket connection
│   │   └── useMonitoring.ts# Monitoring logic
│   │
│   ├── App.vue             # Root component
│   └── main.ts             # Application entry point
```

**Key Patterns**:
- **Composition API**: Modern Vue 3 approach
- **TypeScript**: Type safety throughout
- **Pinia**: Centralized state management
- **Composables**: Reusable logic extraction
- **Route Guards**: Authentication checks

### 2. Backend Architecture (Go)

```
cmd/
├── server/                 # Main server application
│   └── main.go            # Entry point
└── agent/                 # Agent application (future)
    └── main.go

internal/                   # Private application code
├── api/                   # API layer
│   ├── router.go          # Route definitions
│   ├── middleware/        # HTTP middleware
│   │   ├── auth.go        # JWT authentication
│   │   ├── cors.go        # CORS handling
│   │   ├── logger.go      # Request logging
│   │   └── i18n.go        # Internationalization
│   └── handlers/          # HTTP handlers
│       ├── auth.go        # Authentication endpoints
│       ├── user.go        # User management endpoints
│       ├── server.go      # Server management endpoints
│       └── monitor.go     # Monitoring endpoints
│
├── core/                  # Business logic
│   ├── auth/             # Authentication logic
│   ├── user/             # User management logic
│   ├── server/           # Server management logic
│   └── monitor/          # Monitoring logic
│
├── database/             # Database layer
│   ├── database.go       # Database connection
│   ├── models/           # GORM models
│   │   └── models.go
│   └── migrations/       # SQL migrations
│       └── 001_initial.sql
│
├── plugins/              # Plugin system
│   ├── engine.go        # Plugin loader
│   └── registry.go      # Plugin registry
│
└── i18n/                # Backend i18n
    ├── i18n.go          # i18n engine
    ├── zh-CN.yaml       # Chinese messages
    └── en-US.yaml       # English messages

pkg/                      # Public libraries
├── config/              # Configuration management
│   └── config.go
├── logger/              # Logging utilities
│   └── logger.go
├── crypto/              # Encryption utilities
│   └── crypto.go
└── utils/               # General utilities
    └── utils.go
```

**Key Patterns**:
- **Clean Architecture**: Separation of concerns
- **Dependency Injection**: Loose coupling
- **Interface-based Design**: Testability
- **Error Handling**: Consistent error types
- **Middleware Chain**: Request processing pipeline

### 3. Database Design

#### Entity Relationship Diagram

```
┌──────────────┐
│    users     │
│──────────────│
│ id (PK)      │
│ username     │
│ email        │
│ password_hash│
│ role         │
└──────┬───────┘
       │
       │ 1:N
       ▼
┌──────────────┐         ┌──────────────┐
│ user_tokens  │         │  user_logs   │
│──────────────│         │──────────────│
│ id (PK)      │         │ id (PK)      │
│ user_id (FK) │         │ user_id (FK) │
│ token        │         │ action       │
└──────────────┘         └──────────────┘

┌──────────────┐
│   servers    │
│──────────────│
│ id (PK)      │
│ name         │
│ host         │
│ ssh_key_id(FK)│
│ created_by(FK)│
└──────┬───────┘
       │
       │ 1:N
       ▼
┌──────────────┐         ┌──────────────┐
│monitor_metrics│        │   alerts     │
│──────────────│         │──────────────│
│ id (PK)      │         │ id (PK)      │
│ server_id(FK)│         │ server_id(FK)│
│ metric_type  │         │ alert_type   │
│ value        │         │ severity     │
└──────────────┘         └──────────────┘

┌──────────────┐
│  ssh_keys    │
│──────────────│
│ id (PK)      │
│ name         │
│ public_key   │
│ private_key  │
│ created_by(FK)│
└──────────────┘

┌──────────────┐
│   plugins    │
│──────────────│
│ id (PK)      │
│ name         │
│ version      │
│ status       │
└──────┬───────┘
       │
       │ 1:N
       ▼
┌──────────────┐
│plugin_settings│
│──────────────│
│ id (PK)      │
│ plugin_id(FK)│
│ key          │
│ value        │
└──────────────┘
```

### 4. Authentication & Authorization Flow

```
User Login Flow:
┌──────────┐
│  Client  │
└────┬─────┘
     │ 1. POST /api/v1/auth/login
     │    {username, password}
     ▼
┌────────────┐
│   Router   │
└────┬───────┘
     │ 2. Route to handler
     ▼
┌────────────┐
│  Handler   │
└────┬───────┘
     │ 3. Verify credentials
     ▼
┌────────────┐
│    Auth    │
│   Module   │
└────┬───────┘
     │ 4. Query user by username
     ▼
┌────────────┐
│  Database  │
└────┬───────┘
     │ 5. Return user record
     ▼
┌────────────┐
│    Auth    │
│   Module   │
└────┬───────┘
     │ 6. Compare password hash
     │ 7. Generate JWT tokens
     │    (access + refresh)
     ▼
┌────────────┐
│  Handler   │
└────┬───────┘
     │ 8. Return tokens + user info
     ▼
┌────────────┐
│   Client   │
│ Stores     │
│ tokens in  │
│ localStorage│
└────────────┘

Authenticated Request Flow:
┌──────────┐
│  Client  │
└────┬─────┘
     │ 1. GET /api/v1/servers
     │    Authorization: Bearer <token>
     ▼
┌────────────┐
│   Router   │
└────┬───────┘
     │ 2. Auth middleware
     ▼
┌────────────┐
│    Auth    │
│ Middleware │
└────┬───────┘
     │ 3. Extract token from header
     │ 4. Validate JWT signature
     │ 5. Check expiration
     │ 6. Extract user ID & role
     │ 7. Set user context
     ▼
┌────────────┐
│  Handler   │
└────┬───────┘
     │ 8. Check permissions (RBAC)
     │ 9. Process request
     ▼
┌────────────┐
│   Client   │
│  Receives  │
│  Response  │
└────────────┘
```

### 5. Real-Time Monitoring Flow

```
WebSocket Connection:
┌──────────┐
│  Client  │
└────┬─────┘
     │ 1. WS /api/v1/servers/:id/metrics/stream
     │    Upgrade: websocket
     │    Authorization: Bearer <token>
     ▼
┌────────────┐
│ WS Handler │
└────┬───────┘
     │ 2. Authenticate token
     │ 3. Establish WS connection
     │ 4. Start metrics collection goroutine
     ▼
┌────────────┐
│  Monitor   │
│   Module   │
└────┬───────┘
     │ 5. Connect to target server (SSH)
     │ 6. Collect metrics (CPU, Memory, etc.)
     │ 7. Send metrics via WebSocket
     │    (every 1 second)
     ▼
┌────────────┐
│   Client   │
│  Updates   │
│   Charts   │
└────────────┘

Data Collection on Server:
┌────────────┐
│ NexusPanel │
│   Server   │
└────┬───────┘
     │ SSH Connection
     ▼
┌────────────┐
│   Target   │
│   Server   │
└────┬───────┘
     │ Execute commands:
     │  • top -bn1 (CPU)
     │  • free -m (Memory)
     │  • df -h (Disk)
     │  • ifconfig (Network)
     ▼
┌────────────┐
│   Parse    │
│   Output   │
└────┬───────┘
     │ Format as JSON
     ▼
┌────────────┐
│   Store    │
│  Metrics   │
│ (Optional) │
└────┬───────┘
     │ Save to DB for history
     ▼
┌────────────┐
│   Stream   │
│  to Client │
│via WebSocket│
└────────────┘
```

---

## Data Flow

### 1. Server Addition Flow

```
1. User fills form on frontend
2. Frontend validates input
3. POST /api/v1/servers with server details
4. Auth middleware validates JWT
5. Handler validates request body
6. Server module:
   - Tests SSH connection
   - Retrieves system info
   - Creates server record in DB
7. Success response returned
8. Frontend updates server list
9. Background job starts monitoring
```

### 2. Alert Processing Flow

```
1. Monitor module collects metrics
2. Compare metrics against alert rules
3. If threshold exceeded:
   - Create alert record
   - Trigger notification service
4. Alert displayed in UI (WebSocket)
5. User can acknowledge/resolve alert
6. Alert status updated in DB
```

---

## Security Architecture

### 1. Authentication

- **JWT Tokens**: Stateless authentication
- **Token Types**: Access (short-lived) + Refresh (long-lived)
- **Token Storage**: HttpOnly cookies (production) / localStorage (dev)
- **Password Requirements**: Min 8 chars, uppercase, lowercase, number
- **Password Hashing**: bcrypt with cost 10

### 2. Authorization

- **RBAC**: Role-Based Access Control
- **Roles**: admin, operator, user, guest
- **Permissions**: Granular permissions per endpoint
- **Middleware**: Authorization checks before handlers

### 3. Data Protection

- **Encryption at Rest**: 
  - SSH keys encrypted with AES-256
  - Database passwords encrypted
- **Encryption in Transit**: HTTPS/WSS in production
- **Sensitive Data**: Never logged or exposed in errors

### 4. Input Validation

- **Request Validation**: All inputs validated
- **SQL Injection Prevention**: Parameterized queries (GORM)
- **XSS Prevention**: Output escaping
- **CSRF Protection**: SameSite cookies + CSRF tokens

### 5. Rate Limiting

- **Authentication Endpoints**: 10 req/min per IP
- **API Endpoints**: 100 req/min per user
- **WebSocket**: Max 5 concurrent connections per user

### 6. Audit Logging

- All sensitive operations logged:
  - User login/logout
  - User creation/deletion
  - Server addition/removal
  - Configuration changes
- Logs include: timestamp, user, action, IP, result

---

## Deployment Architecture

### 1. Single Server Deployment

```
┌──────────────────────────────────────┐
│         Single Server                 │
│                                       │
│  ┌─────────────────────────────────┐│
│  │         Nginx                    ││
│  │  (Reverse Proxy + Static Files) ││
│  └─────────────┬───────────────────┘│
│                │                     │
│  ┌─────────────▼───────────────────┐│
│  │      NexusPanel Server          ││
│  │       (Go Binary)                ││
│  └─────────────┬───────────────────┘│
│                │                     │
│  ┌─────────────▼───────────────────┐│
│  │       PostgreSQL                ││
│  └─────────────────────────────────┘│
└──────────────────────────────────────┘
```

### 2. Docker Compose Deployment

```yaml
services:
  web:
    image: nginx:alpine
    ports: ["80:80", "443:443"]
    volumes:
      - ./web/dist:/usr/share/nginx/html
  
  api:
    image: nexuspanel/server:latest
    environment:
      - DB_HOST=db
      - REDIS_HOST=redis
  
  db:
    image: postgres:13
    volumes:
      - postgres_data:/var/lib/postgresql/data
  
  redis:
    image: redis:7-alpine
```

### 3. Kubernetes Deployment (Future)

```
┌────────────────────────────────────────┐
│         Kubernetes Cluster             │
│                                        │
│  ┌──────────────────────────────────┐ │
│  │        Ingress Controller        │ │
│  └─────────┬────────────────────────┘ │
│            │                           │
│  ┌─────────▼──────────┐  ┌──────────┐ │
│  │   API Pods (3x)    │  │  Redis   │ │
│  │  NexusPanel Server │  │   Pod    │ │
│  └────────────────────┘  └──────────┘ │
│            │                           │
│  ┌─────────▼──────────┐                │
│  │  PostgreSQL Pod    │                │
│  │  (StatefulSet)     │                │
│  └────────────────────┘                │
└────────────────────────────────────────┘
```

---

## Scalability & Performance

### 1. Horizontal Scaling

- **Stateless API**: Can run multiple instances
- **Load Balancer**: Nginx/HAProxy for distribution
- **Session Storage**: Redis for shared state
- **Database**: Read replicas for scaling reads

### 2. Performance Optimization

- **Database Indexes**: On frequently queried columns
- **Connection Pooling**: Reuse DB connections
- **Caching Strategy**: 
  - Redis for frequently accessed data
  - In-memory cache for configuration
- **Query Optimization**: 
  - Avoid N+1 queries
  - Use eager loading where appropriate
- **Asset Optimization**:
  - Minified JS/CSS
  - Gzip compression
  - CDN for static assets (future)

### 3. Monitoring & Observability

- **Metrics**: Prometheus-compatible metrics
- **Logging**: Structured logging (JSON format)
- **Tracing**: Distributed tracing (future)
- **Health Checks**: `/health` endpoint for readiness

---

**Document End**
