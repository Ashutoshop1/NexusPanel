# NexusPanel API 文档

## 文档信息

- **版本**: v0.1.0
- **基础 URL**: `http://localhost:8080/api/v1` (开发环境)
- **认证方式**: JWT Bearer Token
- **Content-Type**: `application/json`
- **最后更新**: 2024

---

## 目录

1. [认证授权](#认证授权)
2. [用户管理](#用户管理)
3. [服务器管理](#服务器管理)
4. [监控管理](#监控管理)
5. [告警管理](#告警管理)
6. [任务管理](#任务管理)
7. [插件管理](#插件管理)
8. [WebSocket API](#websocket-api)
9. [错误码](#错误码)

---

## 认证授权

所有 API 请求（除认证端点外）都需要在 `Authorization` 头中包含有效的 JWT 令牌：

\`\`\`
Authorization: Bearer <your-jwt-token>
\`\`\`

### 用户注册

创建新用户账号。

**端点**: \`POST /api/v1/auth/register\`

**请求体**:
\`\`\`json
{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "SecureP@ss123",
  "language": "zh-CN"
}
\`\`\`

**响应** (201 Created):
\`\`\`json
{
  "code": 201,
  "message": "用户注册成功",
  "data": {
    "user": {
      "id": 1,
      "username": "johndoe",
      "email": "john@example.com",
      "role": "user",
      "language": "zh-CN",
      "created_at": "2024-01-01T00:00:00Z"
    }
  }
}
\`\`\`

### 用户登录

认证并获取访问令牌和刷新令牌。

**端点**: \`POST /api/v1/auth/login\`

**请求体**:
\`\`\`json
{
  "username": "johndoe",
  "password": "SecureP@ss123"
}
\`\`\`

**响应** (200 OK):
\`\`\`json
{
  "code": 200,
  "message": "登录成功",
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
      "language": "zh-CN"
    }
  }
}
\`\`\`

### 刷新令牌

使用刷新令牌获取新的访问令牌。

**端点**: \`POST /api/v1/auth/refresh\`

**请求体**:
\`\`\`json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
\`\`\`

**响应** (200 OK):
\`\`\`json
{
  "code": 200,
  "message": "令牌刷新成功",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_in": 3600,
    "token_type": "Bearer"
  }
}
\`\`\`

### 用户登出

注销当前会话。

**端点**: \`POST /api/v1/auth/logout\`

**请求头**: 需要 Authorization

**响应** (200 OK):
\`\`\`json
{
  "code": 200,
  "message": "登出成功"
}
\`\`\`

---

## 用户管理

### 获取用户列表

获取分页的用户列表（仅管理员）。

**端点**: \`GET /api/v1/users\`

**请求头**: 需要 Authorization (admin)

**查询参数**:
- \`page\` (可选): 页码（默认: 1）
- \`per_page\` (可选): 每页数量（默认: 20, 最大: 100）
- \`role\` (可选): 按角色过滤
- \`status\` (可选): 按状态过滤
- \`search\` (可选): 按用户名或邮箱搜索

**响应** (200 OK):
\`\`\`json
{
  "code": 200,
  "message": "成功",
  "data": {
    "users": [
      {
        "id": 1,
        "username": "admin",
        "email": "admin@example.com",
        "role": "admin",
        "status": "active",
        "language": "zh-CN",
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
\`\`\`

### 获取用户详情

获取特定用户的详细信息。

**端点**: \`GET /api/v1/users/:id\`

**请求头**: 需要 Authorization

**响应** (200 OK):
\`\`\`json
{
  "code": 200,
  "message": "成功",
  "data": {
    "user": {
      "id": 1,
      "username": "johndoe",
      "email": "john@example.com",
      "avatar": "https://example.com/avatar.jpg",
      "role": "user",
      "status": "active",
      "language": "zh-CN",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-15T10:30:00Z",
      "last_login_at": "2024-01-15T10:30:00Z"
    }
  }
}
\`\`\`

---

## 服务器管理

### 获取服务器列表

获取所有托管服务器的列表。

**端点**: \`GET /api/v1/servers\`

**请求头**: 需要 Authorization

**查询参数**:
- \`page\` (可选): 页码
- \`per_page\` (可选): 每页数量
- \`group_id\` (可选): 按分组过滤
- \`status\` (可选): 按状态过滤 (online/offline/unknown)
- \`search\` (可选): 按名称或主机搜索

**响应** (200 OK):
\`\`\`json
{
  "code": 200,
  "message": "成功",
  "data": {
    "servers": [
      {
        "id": 1,
        "name": "Web服务器01",
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
\`\`\`

### 添加服务器

添加新服务器到管理系统。

**端点**: \`POST /api/v1/servers\`

**请求头**: 需要 Authorization

**请求体**:
\`\`\`json
{
  "name": "新服务器",
  "host": "192.168.1.200",
  "port": 22,
  "ssh_user": "root",
  "ssh_key_id": 1,
  "group_ids": [1, 2]
}
\`\`\`

**响应** (201 Created):
\`\`\`json
{
  "code": 201,
  "message": "服务器添加成功",
  "data": {
    "server": {
      "id": 10,
      "name": "新服务器",
      "host": "192.168.1.200",
      "port": 22,
      "ssh_user": "root",
      "status": "unknown",
      "created_at": "2024-01-15T15:00:00Z"
    }
  }
}
\`\`\`

### 测试连接

测试服务器的 SSH 连接。

**端点**: \`POST /api/v1/servers/:id/test-connection\`

**请求头**: 需要 Authorization

**响应** (200 OK):
\`\`\`json
{
  "code": 200,
  "message": "连接成功",
  "data": {
    "connected": true,
    "latency_ms": 45,
    "os_info": "Ubuntu 22.04 LTS"
  }
}
\`\`\`

---

## 监控管理

### 获取当前指标

获取服务器的当前系统指标。

**端点**: \`GET /api/v1/servers/:id/metrics\`

**请求头**: 需要 Authorization

**响应** (200 OK):
\`\`\`json
{
  "code": 200,
  "message": "成功",
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
\`\`\`

### 获取历史指标

获取历史监控数据。

**端点**: \`GET /api/v1/servers/:id/metrics/history\`

**请求头**: 需要 Authorization

**查询参数**:
- \`metric_type\` (必需): cpu, memory, disk, network
- \`start_time\` (必需): ISO 8601 时间戳
- \`end_time\` (必需): ISO 8601 时间戳
- \`interval\` (可选): 聚合间隔 (1m, 5m, 1h, 1d)

**响应** (200 OK):
\`\`\`json
{
  "code": 200,
  "message": "成功",
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
\`\`\`

---

## 告警管理

### 获取告警列表

获取告警列表。

**端点**: \`GET /api/v1/alerts\`

**请求头**: 需要 Authorization

**查询参数**:
- \`server_id\` (可选): 按服务器过滤
- \`severity\` (可选): critical, warning, info
- \`status\` (可选): triggered, resolved, acknowledged

**响应** (200 OK):
\`\`\`json
{
  "code": 200,
  "message": "成功",
  "data": {
    "alerts": [
      {
        "id": 1,
        "server_id": 1,
        "server_name": "Web服务器01",
        "alert_type": "cpu_high",
        "severity": "warning",
        "message": "CPU 使用率超过 80%",
        "status": "triggered",
        "triggered_at": "2024-01-15T16:00:00Z",
        "resolved_at": null
      }
    ]
  }
}
\`\`\`

---

## WebSocket API

### 实时指标流

通过 WebSocket 流式传输实时服务器指标。

**端点**: \`WS /api/v1/servers/:id/metrics/stream\`

**请求头**: 
- \`Upgrade: websocket\`
- \`Authorization: Bearer <token>\`

**消息格式**:
\`\`\`json
{
  "timestamp": "2024-01-15T18:00:00Z",
  "cpu_percent": 35.5,
  "memory_percent": 50.0,
  "disk_percent": 45.2,
  "network_in": 1024000,
  "network_out": 512000
}
\`\`\`

### Web 终端

通过 WebSocket 的交互式 SSH 终端。

**端点**: \`WS /api/v1/servers/:id/terminal\`

**请求头**: 
- \`Upgrade: websocket\`
- \`Authorization: Bearer <token>\`

**客户端 → 服务器**:
\`\`\`json
{
  "type": "input",
  "data": "ls -la\n"
}
\`\`\`

**服务器 → 客户端**:
\`\`\`json
{
  "type": "output",
  "data": "total 64\ndrwxr-xr-x  5 root root 4096 Jan 15 18:00 .\n..."
}
\`\`\`

---

## 错误码

### HTTP 状态码

- \`200 OK\` - 请求成功
- \`201 Created\` - 资源创建成功
- \`400 Bad Request\` - 请求参数无效
- \`401 Unauthorized\` - 需要认证或认证失败
- \`403 Forbidden\` - 权限不足
- \`404 Not Found\` - 资源未找到
- \`409 Conflict\` - 资源冲突（如用户名重复）
- \`422 Unprocessable Entity\` - 验证错误
- \`429 Too Many Requests\` - 请求频率超限
- \`500 Internal Server Error\` - 服务器错误
- \`501 Not Implemented\` - 功能未实现
- \`503 Service Unavailable\` - 服务暂时不可用

### 应用错误码

\`\`\`json
{
  "code": 40001,
  "message": "凭证无效",
  "details": "用户名或密码不正确"
}
\`\`\`

常见错误码:
- \`40001\` - 凭证无效
- \`40002\` - 令牌过期
- \`40003\` - 令牌无效
- \`40301\` - 权限不足
- \`40401\` - 资源未找到
- \`40901\` - 资源已存在
- \`42201\` - 验证错误
- \`50001\` - 服务器内部错误
- \`50101\` - 功能未实现
- \`50301\` - 服务不可用

---

## 频率限制

API 频率限制:
- **通用 API**: 每个用户每分钟 100 次请求
- **认证 API**: 每个 IP 每分钟 10 次请求
- **WebSocket 连接**: 每个用户最多 5 个并发连接

响应头中的频率限制信息:
\`\`\`
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 95
X-RateLimit-Reset: 1641405600
\`\`\`

---

## 分页

分页端点返回以下结构:

\`\`\`json
{
  "code": 200,
  "message": "成功",
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
\`\`\`

---

**文档结束**
