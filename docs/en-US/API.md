# NexusPanel API Documentation

## Document Information

- **Version**: v0.1.0
- **Base URL**: `http://localhost:8080/api/v1` (development)
- **Authentication**: JWT Bearer Token
- **Content-Type**: `application/json`
- **Last Updated**: 2024

---

## Table of Contents

1. [Authentication](#authentication)
2. [User Management](#user-management)
3. [Server Management](#server-management)
4. [Monitoring](#monitoring)
5. [Alerting](#alerting)
6. [Task Management](#task-management)
7. [Plugin Management](#plugin-management)
8. [WebSocket APIs](#websocket-apis)
9. [Error Codes](#error-codes)

---

## Authentication

All API requests (except authentication endpoints) require a valid JWT token in the `Authorization` header:

```
Authorization: Bearer <your-jwt-token>
```

### Register

Create a new user account.

**Endpoint**: `POST /api/v1/auth/register`

**Request Body**:
```json
{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "SecureP@ss123",
  "language": "en-US"
}
```

**Response** (201 Created):
```json
{
  "code": 201,
  "message": "User registered successfully",
  "data": {
    "user": {
      "id": 1,
      "username": "johndoe",
      "email": "john@example.com",
      "role": "user",
      "language": "en-US",
      "created_at": "2024-01-01T00:00:00Z"
    }
  }
}
```

### Login

Authenticate and receive access and refresh tokens.

**Endpoint**: `POST /api/v1/auth/login`

**Request Body**:
```json
{
  "username": "johndoe",
  "password": "SecureP@ss123"
}
```

**Response** (200 OK):
```json
{
  "code": 200,
  "message": "Login successful",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_in": 3600,
    "token_type": "Bearer",
    "user": {
      "id": 1,
      "username": "johndoe",
      "email": "john@example.com",
      "role": "user",
      "language": "en-US"
    }
  }
}
```

### Refresh Token

Get a new access token using refresh token.

**Endpoint**: `POST /api/v1/auth/refresh`

**Request Body**:
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Response** (200 OK):
```json
{
  "code": 200,
  "message": "Token refreshed successfully",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_in": 3600,
    "token_type": "Bearer"
  }
}
```

### Logout

Invalidate current session.

**Endpoint**: `POST /api/v1/auth/logout`

**Headers**: Authorization required

**Response** (200 OK):
```json
{
  "code": 200,
  "message": "Logout successful"
}
```

### Reset Password

Request password reset.

**Endpoint**: `POST /api/v1/auth/reset-password`

**Request Body**:
```json
{
  "email": "john@example.com"
}
```

**Response** (200 OK):
```json
{
  "code": 200,
  "message": "Password reset email sent"
}
```

---

## User Management

### List Users

Get a paginated list of users (admin only).

**Endpoint**: `GET /api/v1/users`

**Headers**: Authorization required (admin)

**Query Parameters**:
- `page` (optional): Page number (default: 1)
- `per_page` (optional): Items per page (default: 20, max: 100)
- `role` (optional): Filter by role
- `status` (optional): Filter by status
- `search` (optional): Search by username or email

**Response** (200 OK):
```json
{
  "code": 200,
  "message": "Success",
  "data": {
    "users": [
      {
        "id": 1,
        "username": "admin",
        "email": "admin@example.com",
        "role": "admin",
        "status": "active",
        "language": "en-US",
        "created_at": "2024-01-01T00:00:00Z",
        "last_login_at": "2024-01-15T10:30:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "per_page": 20,
      "total_pages": 5,
      "total_count": 95
    }
  }
}
```

### Get User

Get details of a specific user.

**Endpoint**: `GET /api/v1/users/:id`

**Headers**: Authorization required

**Response** (200 OK):
```json
{
  "code": 200,
  "message": "Success",
  "data": {
    "user": {
      "id": 1,
      "username": "johndoe",
      "email": "john@example.com",
      "avatar": "https://example.com/avatar.jpg",
      "role": "user",
      "status": "active",
      "language": "en-US",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-15T10:30:00Z",
      "last_login_at": "2024-01-15T10:30:00Z"
    }
  }
}
```

### Create User

Create a new user (admin only).

**Endpoint**: `POST /api/v1/users`

**Headers**: Authorization required (admin)

**Request Body**:
```json
{
  "username": "newuser",
  "email": "newuser@example.com",
  "password": "SecureP@ss123",
  "role": "user",
  "language": "en-US"
}
```

**Response** (201 Created):
```json
{
  "code": 201,
  "message": "User created successfully",
  "data": {
    "user": {
      "id": 10,
      "username": "newuser",
      "email": "newuser@example.com",
      "role": "user",
      "status": "active",
      "language": "en-US",
      "created_at": "2024-01-15T12:00:00Z"
    }
  }
}
```

### Update User

Update user information.

**Endpoint**: `PUT /api/v1/users/:id`

**Headers**: Authorization required

**Request Body**:
```json
{
  "email": "newemail@example.com",
  "avatar": "https://example.com/new-avatar.jpg",
  "language": "zh-CN"
}
```

**Response** (200 OK):
```json
{
  "code": 200,
  "message": "User updated successfully",
  "data": {
    "user": {
      "id": 1,
      "username": "johndoe",
      "email": "newemail@example.com",
      "avatar": "https://example.com/new-avatar.jpg",
      "role": "user",
      "status": "active",
      "language": "zh-CN",
      "updated_at": "2024-01-15T13:00:00Z"
    }
  }
}
```

### Delete User

Delete a user (admin only).

**Endpoint**: `DELETE /api/v1/users/:id`

**Headers**: Authorization required (admin)

**Response** (200 OK):
```json
{
  "code": 200,
  "message": "User deleted successfully"
}
```

---

## Server Management

### List Servers

Get a list of all managed servers.

**Endpoint**: `GET /api/v1/servers`

**Headers**: Authorization required

**Query Parameters**:
- `page` (optional): Page number
- `per_page` (optional): Items per page
- `group_id` (optional): Filter by group
- `status` (optional): Filter by status (online/offline/unknown)
- `search` (optional): Search by name or host

**Response** (200 OK):
```json
{
  "code": 200,
  "message": "Success",
  "data": {
    "servers": [
      {
        "id": 1,
        "name": "Web Server 01",
        "host": "192.168.1.100",
        "port": 22,
        "ssh_user": "root",
        "status": "online",
        "os_info": "Ubuntu 22.04 LTS",
        "last_heartbeat": "2024-01-15T14:00:00Z",
        "created_at": "2024-01-01T00:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "per_page": 20,
      "total_pages": 3,
      "total_count": 45
    }
  }
}
```

### Get Server

Get details of a specific server.

**Endpoint**: `GET /api/v1/servers/:id`

**Headers**: Authorization required

**Response** (200 OK):
```json
{
  "code": 200,
  "message": "Success",
  "data": {
    "server": {
      "id": 1,
      "name": "Web Server 01",
      "host": "192.168.1.100",
      "port": 22,
      "ssh_user": "root",
      "ssh_key_id": 1,
      "status": "online",
      "os_info": "Ubuntu 22.04 LTS",
      "last_heartbeat": "2024-01-15T14:00:00Z",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-15T14:00:00Z",
      "groups": [
        {
          "id": 1,
          "name": "Production"
        }
      ]
    }
  }
}
```

### Add Server

Add a new server to manage.

**Endpoint**: `POST /api/v1/servers`

**Headers**: Authorization required

**Request Body**:
```json
{
  "name": "New Server",
  "host": "192.168.1.200",
  "port": 22,
  "ssh_user": "root",
  "ssh_key_id": 1,
  "group_ids": [1, 2]
}
```

**Response** (201 Created):
```json
{
  "code": 201,
  "message": "Server added successfully",
  "data": {
    "server": {
      "id": 10,
      "name": "New Server",
      "host": "192.168.1.200",
      "port": 22,
      "ssh_user": "root",
      "status": "unknown",
      "created_at": "2024-01-15T15:00:00Z"
    }
  }
}
```

### Update Server

Update server information.

**Endpoint**: `PUT /api/v1/servers/:id`

**Headers**: Authorization required

**Request Body**:
```json
{
  "name": "Updated Server Name",
  "port": 2222,
  "group_ids": [1]
}
```

**Response** (200 OK):
```json
{
  "code": 200,
  "message": "Server updated successfully",
  "data": {
    "server": {
      "id": 1,
      "name": "Updated Server Name",
      "host": "192.168.1.100",
      "port": 2222,
      "ssh_user": "root",
      "status": "online",
      "updated_at": "2024-01-15T15:30:00Z"
    }
  }
}
```

### Delete Server

Remove a server from management.

**Endpoint**: `DELETE /api/v1/servers/:id`

**Headers**: Authorization required

**Response** (200 OK):
```json
{
  "code": 200,
  "message": "Server deleted successfully"
}
```

### Test Connection

Test SSH connection to a server.

**Endpoint**: `POST /api/v1/servers/:id/test-connection`

**Headers**: Authorization required

**Response** (200 OK):
```json
{
  "code": 200,
  "message": "Connection successful",
  "data": {
    "connected": true,
    "latency_ms": 45,
    "os_info": "Ubuntu 22.04 LTS"
  }
}
```

### Get System Info

Get detailed system information.

**Endpoint**: `GET /api/v1/servers/:id/info`

**Headers**: Authorization required

**Response** (200 OK):
```json
{
  "code": 200,
  "message": "Success",
  "data": {
    "hostname": "web-server-01",
    "os": "Linux",
    "platform": "ubuntu",
    "platform_version": "22.04",
    "kernel_version": "5.15.0-91-generic",
    "architecture": "x86_64",
    "cpu_count": 4,
    "total_memory": 8589934592,
    "uptime": 864000
  }
}
```

---

## Monitoring

### Get Current Metrics

Get current system metrics for a server.

**Endpoint**: `GET /api/v1/servers/:id/metrics`

**Headers**: Authorization required

**Response** (200 OK):
```json
{
  "code": 200,
  "message": "Success",
  "data": {
    "timestamp": "2024-01-15T16:00:00Z",
    "cpu": {
      "usage_percent": 35.5,
      "cores": [
        {"core": 0, "usage_percent": 30.2},
        {"core": 1, "usage_percent": 40.8}
      ]
    },
    "memory": {
      "total": 8589934592,
      "used": 4294967296,
      "free": 4294967296,
      "usage_percent": 50.0
    },
    "disk": [
      {
        "device": "/dev/sda1",
        "mount_point": "/",
        "total": 107374182400,
        "used": 53687091200,
        "free": 53687091200,
        "usage_percent": 50.0
      }
    ],
    "network": {
      "bytes_sent": 1073741824,
      "bytes_recv": 2147483648,
      "packets_sent": 1000000,
      "packets_recv": 2000000
    }
  }
}
```

### Get Historical Metrics

Get historical metrics data.

**Endpoint**: `GET /api/v1/servers/:id/metrics/history`

**Headers**: Authorization required

**Query Parameters**:
- `metric_type` (required): cpu, memory, disk, network
- `start_time` (required): ISO 8601 timestamp
- `end_time` (required): ISO 8601 timestamp
- `interval` (optional): Aggregation interval (1m, 5m, 1h, 1d)

**Response** (200 OK):
```json
{
  "code": 200,
  "message": "Success",
  "data": {
    "metric_type": "cpu",
    "interval": "5m",
    "data_points": [
      {
        "timestamp": "2024-01-15T15:00:00Z",
        "value": 32.5
      },
      {
        "timestamp": "2024-01-15T15:05:00Z",
        "value": 35.2
      }
    ]
  }
}
```

### Get Process List

Get list of running processes.

**Endpoint**: `GET /api/v1/servers/:id/processes`

**Headers**: Authorization required

**Query Parameters**:
- `sort` (optional): Sort by (cpu, memory, name)
- `limit` (optional): Limit number of results

**Response** (200 OK):
```json
{
  "code": 200,
  "message": "Success",
  "data": {
    "processes": [
      {
        "pid": 1234,
        "name": "nginx",
        "user": "www-data",
        "cpu_percent": 5.2,
        "memory_percent": 2.1,
        "status": "running",
        "command": "nginx: worker process"
      }
    ]
  }
}
```

---

## Alerting

### List Alerts

Get list of alerts.

**Endpoint**: `GET /api/v1/alerts`

**Headers**: Authorization required

**Query Parameters**:
- `server_id` (optional): Filter by server
- `severity` (optional): critical, warning, info
- `status` (optional): triggered, resolved, acknowledged

**Response** (200 OK):
```json
{
  "code": 200,
  "message": "Success",
  "data": {
    "alerts": [
      {
        "id": 1,
        "server_id": 1,
        "server_name": "Web Server 01",
        "alert_type": "cpu_high",
        "severity": "warning",
        "message": "CPU usage exceeded 80%",
        "status": "triggered",
        "triggered_at": "2024-01-15T16:00:00Z",
        "resolved_at": null
      }
    ]
  }
}
```

### List Alert Rules

Get configured alert rules.

**Endpoint**: `GET /api/v1/alerts/rules`

**Headers**: Authorization required

**Response** (200 OK):
```json
{
  "code": 200,
  "message": "Success",
  "data": {
    "rules": [
      {
        "id": 1,
        "name": "High CPU Usage",
        "metric_type": "cpu",
        "condition": "greater_than",
        "threshold": 80,
        "severity": "warning",
        "enabled": true,
        "created_at": "2024-01-01T00:00:00Z"
      }
    ]
  }
}
```

### Create Alert Rule

Create a new alert rule.

**Endpoint**: `POST /api/v1/alerts/rules`

**Headers**: Authorization required

**Request Body**:
```json
{
  "name": "High Memory Usage",
  "metric_type": "memory",
  "condition": "greater_than",
  "threshold": 90,
  "severity": "critical",
  "enabled": true
}
```

**Response** (201 Created):
```json
{
  "code": 201,
  "message": "Alert rule created successfully",
  "data": {
    "rule": {
      "id": 10,
      "name": "High Memory Usage",
      "metric_type": "memory",
      "condition": "greater_than",
      "threshold": 90,
      "severity": "critical",
      "enabled": true,
      "created_at": "2024-01-15T17:00:00Z"
    }
  }
}
```

---

## Task Management

### List Tasks

Get list of tasks.

**Endpoint**: `GET /api/v1/tasks`

**Headers**: Authorization required

**Response** (200 OK):
```json
{
  "code": 200,
  "message": "Success",
  "data": {
    "tasks": [
      {
        "id": 1,
        "name": "Daily Backup",
        "type": "scheduled",
        "cron_expression": "0 2 * * *",
        "status": "active",
        "last_run_at": "2024-01-15T02:00:00Z",
        "next_run_at": "2024-01-16T02:00:00Z",
        "created_at": "2024-01-01T00:00:00Z"
      }
    ]
  }
}
```

---

## Plugin Management

### List Plugins

Get list of installed plugins.

**Endpoint**: `GET /api/v1/plugins`

**Headers**: Authorization required

**Response** (200 OK):
```json
{
  "code": 200,
  "message": "Success",
  "data": {
    "plugins": [
      {
        "id": 1,
        "name": "docker-manager",
        "version": "1.0.0",
        "description": "Docker container management plugin",
        "author": "NexusPanel Team",
        "status": "active",
        "installed_at": "2024-01-01T00:00:00Z"
      }
    ]
  }
}
```

---

## WebSocket APIs

### Real-Time Metrics Stream

Stream real-time server metrics via WebSocket.

**Endpoint**: `WS /api/v1/servers/:id/metrics/stream`

**Headers**: 
- `Upgrade: websocket`
- `Authorization: Bearer <token>`

**Message Format**:
```json
{
  "timestamp": "2024-01-15T18:00:00Z",
  "cpu_percent": 35.5,
  "memory_percent": 50.0,
  "disk_percent": 45.2,
  "network_in": 1024000,
  "network_out": 512000
}
```

### Web Terminal

Interactive SSH terminal via WebSocket.

**Endpoint**: `WS /api/v1/servers/:id/terminal`

**Headers**: 
- `Upgrade: websocket`
- `Authorization: Bearer <token>`

**Client → Server**:
```json
{
  "type": "input",
  "data": "ls -la\n"
}
```

**Server → Client**:
```json
{
  "type": "output",
  "data": "total 64\ndrwxr-xr-x  5 root root 4096 Jan 15 18:00 .\n..."
}
```

---

## Error Codes

### HTTP Status Codes

- `200 OK` - Request succeeded
- `201 Created` - Resource created successfully
- `400 Bad Request` - Invalid request parameters
- `401 Unauthorized` - Authentication required or failed
- `403 Forbidden` - Insufficient permissions
- `404 Not Found` - Resource not found
- `409 Conflict` - Resource conflict (e.g., duplicate username)
- `422 Unprocessable Entity` - Validation error
- `429 Too Many Requests` - Rate limit exceeded
- `500 Internal Server Error` - Server error
- `501 Not Implemented` - Feature not yet implemented
- `503 Service Unavailable` - Service temporarily unavailable

### Application Error Codes

```json
{
  "code": 40001,
  "message": "Invalid credentials",
  "details": "Username or password is incorrect"
}
```

Common error codes:
- `40001` - Invalid credentials
- `40002` - Token expired
- `40003` - Token invalid
- `40301` - Insufficient permissions
- `40401` - Resource not found
- `40901` - Resource already exists
- `42201` - Validation error
- `50001` - Internal server error
- `50101` - Feature not implemented
- `50301` - Service unavailable

---

## Rate Limiting

API rate limits:
- **General APIs**: 100 requests per minute per user
- **Authentication APIs**: 10 requests per minute per IP
- **WebSocket connections**: 5 concurrent connections per user

Rate limit headers in response:
```
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 95
X-RateLimit-Reset: 1641405600
```

---

## Pagination

Paginated endpoints return the following structure:

```json
{
  "code": 200,
  "message": "Success",
  "data": {
    "items": [...],
    "pagination": {
      "current_page": 1,
      "per_page": 20,
      "total_pages": 5,
      "total_count": 100,
      "has_next": true,
      "has_prev": false
    }
  }
}
```

---

**Document End**
