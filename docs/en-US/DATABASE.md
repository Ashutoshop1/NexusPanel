# NexusPanel Database Design Document

## Document Version

- **Version**: v1.0
- **Created**: 2024
- **Last Updated**: 2024
- **Database Type**: PostgreSQL 13+ / SQLite 3.x

---

## Overview

NexusPanel uses a relational database to store system data, supporting PostgreSQL (recommended for production) and SQLite (for development/lightweight deployment). The database design follows the third normal form to ensure data consistency and integrity.

### Design Principles

1. **Modular Design**: Tables organized by functional modules
2. **Data Integrity**: Foreign key constraints ensure referential integrity
3. **Performance Optimization**: Indexes added for frequently queried fields
4. **Extensibility**: JSON fields for flexible configuration storage
5. **Audit Trail**: Timestamps and operator tracking for key operations

---

## Table Structure

### 1. User Management Module

#### 1.1 users (Users Table)

Stores basic user information and authentication data.

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | BIGSERIAL | PRIMARY KEY | User ID |
| username | VARCHAR(64) | NOT NULL, UNIQUE | Username |
| email | VARCHAR(255) | NOT NULL, UNIQUE | Email address |
| password_hash | VARCHAR(255) | NOT NULL | Password hash (bcrypt) |
| avatar | VARCHAR(512) | | Avatar URL |
| role | VARCHAR(32) | NOT NULL, DEFAULT 'user' | Role (admin/operator/user/guest) |
| status | VARCHAR(32) | NOT NULL, DEFAULT 'active' | Status (active/inactive/locked/pending) |
| language | VARCHAR(16) | NOT NULL, DEFAULT 'zh-CN' | Language preference (zh-CN/en-US) |
| created_at | TIMESTAMP | NOT NULL | Creation time |
| updated_at | TIMESTAMP | NOT NULL | Update time |
| last_login_at | TIMESTAMP | | Last login time |

**Indexes**:
- `idx_users_username` (username)
- `idx_users_email` (email)
- `idx_users_status` (status)

**Constraints**:
- `chk_users_role`: role IN ('admin', 'user', 'guest', 'operator')
- `chk_users_status`: status IN ('active', 'inactive', 'locked', 'pending')

---

#### 1.2 roles (Roles Table)

Defines system roles and permissions.

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | BIGSERIAL | PRIMARY KEY | Role ID |
| name | VARCHAR(64) | NOT NULL, UNIQUE | Role name |
| description | TEXT | | Role description |
| permissions | TEXT | JSON format | Permission list (JSON array) |
| created_at | TIMESTAMP | NOT NULL | Creation time |

**Indexes**:
- `idx_roles_name` (name)

**Permissions Format Example**:
```json
["server:read", "server:write", "monitor:read", "task:execute"]
```

---

#### 1.3 user_tokens (User Tokens Table)

Stores user authentication tokens (JWT, password reset tokens, etc.).

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | BIGSERIAL | PRIMARY KEY | Token ID |
| user_id | BIGINT | NOT NULL, FK(users.id) | User ID |
| token | VARCHAR(512) | NOT NULL, UNIQUE | Token string |
| type | VARCHAR(32) | NOT NULL | Token type (access/refresh/reset/verify) |
| expires_at | TIMESTAMP | NOT NULL | Expiration time |
| created_at | TIMESTAMP | NOT NULL | Creation time |

**Indexes**:
- `idx_user_tokens_user_id` (user_id)
- `idx_user_tokens_token` (token)
- `idx_user_tokens_expires_at` (expires_at)

**Foreign Keys**:
- user_id → users(id) ON DELETE CASCADE

---

#### 1.4 user_logs (User Operation Logs)

Records user operation audit logs.

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | BIGSERIAL | PRIMARY KEY | Log ID |
| user_id | BIGINT | NOT NULL, FK(users.id) | User ID |
| action | VARCHAR(128) | NOT NULL | Action type |
| ip | VARCHAR(64) | | IP address |
| user_agent | TEXT | | User agent string |
| details | TEXT | JSON format | Detailed information (JSON object) |
| created_at | TIMESTAMP | NOT NULL | Creation time |

**Indexes**:
- `idx_user_logs_user_id` (user_id)
- `idx_user_logs_action` (action)
- `idx_user_logs_created_at` (created_at)

**Foreign Keys**:
- user_id → users(id) ON DELETE CASCADE

---

### 2. Server Management Module

#### 2.1 ssh_keys (SSH Keys Table)

Stores SSH key pairs (private keys are encrypted).

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | BIGSERIAL | PRIMARY KEY | Key ID |
| name | VARCHAR(128) | NOT NULL | Key name |
| public_key | TEXT | NOT NULL | Public key |
| private_key_encrypted | TEXT | NOT NULL | Encrypted private key |
| passphrase_encrypted | TEXT | | Encrypted passphrase |
| created_at | TIMESTAMP | NOT NULL | Creation time |
| created_by | BIGINT | NOT NULL, FK(users.id) | Creator user ID |

**Indexes**:
- `idx_ssh_keys_created_by` (created_by)

**Foreign Keys**:
- created_by → users(id) ON DELETE CASCADE

---

#### 2.2 server_groups (Server Groups Table)

Server groups supporting hierarchical structure.

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | BIGSERIAL | PRIMARY KEY | Group ID |
| name | VARCHAR(128) | NOT NULL | Group name |
| description | TEXT | | Group description |
| parent_id | BIGINT | FK(server_groups.id) | Parent group ID |
| created_at | TIMESTAMP | NOT NULL | Creation time |

**Indexes**:
- `idx_server_groups_parent_id` (parent_id)

**Foreign Keys**:
- parent_id → server_groups(id) ON DELETE SET NULL

---

#### 2.3 servers (Servers Table)

Stores server basic information and connection configuration.

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | BIGSERIAL | PRIMARY KEY | Server ID |
| name | VARCHAR(128) | NOT NULL | Server name |
| host | VARCHAR(255) | NOT NULL | Host address |
| port | INTEGER | NOT NULL, DEFAULT 22 | SSH port |
| ssh_user | VARCHAR(64) | NOT NULL | SSH username |
| ssh_key_id | BIGINT | FK(ssh_keys.id) | SSH key ID |
| status | VARCHAR(32) | NOT NULL, DEFAULT 'offline' | Status (online/offline/error/maintenance) |
| os_info | TEXT | JSON format | Operating system information |
| last_heartbeat | TIMESTAMP | | Last heartbeat time |
| created_at | TIMESTAMP | NOT NULL | Creation time |
| updated_at | TIMESTAMP | NOT NULL | Update time |
| created_by | BIGINT | NOT NULL, FK(users.id) | Creator user ID |

**Indexes**:
- `idx_servers_host` (host)
- `idx_servers_status` (status)
- `idx_servers_created_by` (created_by)

**Foreign Keys**:
- ssh_key_id → ssh_keys(id) ON DELETE SET NULL
- created_by → users(id) ON DELETE CASCADE

**Constraints**:
- `chk_servers_status`: status IN ('online', 'offline', 'error', 'maintenance')

---

#### 2.4 server_group_relations (Server-Group Relations)

Many-to-many relationship between servers and groups.

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | BIGSERIAL | PRIMARY KEY | Relation ID |
| server_id | BIGINT | NOT NULL, FK(servers.id) | Server ID |
| group_id | BIGINT | NOT NULL, FK(server_groups.id) | Group ID |

**Indexes**:
- `idx_server_group_relations_server_id` (server_id)
- `idx_server_group_relations_group_id` (group_id)

**Foreign Keys**:
- server_id → servers(id) ON DELETE CASCADE
- group_id → server_groups(id) ON DELETE CASCADE

**Constraints**:
- UNIQUE(server_id, group_id)

---

### 3. Monitoring Module

#### 3.1 monitor_metrics (Monitor Metrics Table)

Stores time-series monitoring data.

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | BIGSERIAL | PRIMARY KEY | Metric ID |
| server_id | BIGINT | NOT NULL, FK(servers.id) | Server ID |
| metric_type | VARCHAR(64) | NOT NULL | Metric type (cpu/memory/disk/network) |
| value | DOUBLE PRECISION | NOT NULL | Metric value |
| tags | TEXT | JSON format | Tags (JSON object) |
| timestamp | TIMESTAMP | NOT NULL | Timestamp |

**Indexes**:
- `idx_monitor_metrics_server_id` (server_id)
- `idx_monitor_metrics_metric_type` (metric_type)
- `idx_monitor_metrics_timestamp` (timestamp)
- `idx_monitor_metrics_composite` (server_id, metric_type, timestamp)

**Foreign Keys**:
- server_id → servers(id) ON DELETE CASCADE

---

#### 3.2 alert_rules (Alert Rules Table)

Defines alert rules.

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | BIGSERIAL | PRIMARY KEY | Rule ID |
| name | VARCHAR(128) | NOT NULL | Rule name |
| metric_type | VARCHAR(64) | NOT NULL | Monitor metric type |
| condition | VARCHAR(32) | NOT NULL | Condition (gt/lt/eq/gte/lte) |
| threshold | DOUBLE PRECISION | NOT NULL | Threshold value |
| severity | VARCHAR(32) | NOT NULL | Severity (info/warning/error/critical) |
| enabled | BOOLEAN | NOT NULL, DEFAULT TRUE | Enabled flag |
| created_at | TIMESTAMP | NOT NULL | Creation time |

**Indexes**:
- `idx_alert_rules_enabled` (enabled)

**Constraints**:
- `chk_alert_rules_condition`: condition IN ('gt', 'lt', 'eq', 'gte', 'lte')
- `chk_alert_rules_severity`: severity IN ('info', 'warning', 'error', 'critical')

---

#### 3.3 alerts (Alerts Table)

Stores triggered alerts.

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | BIGSERIAL | PRIMARY KEY | Alert ID |
| server_id | BIGINT | NOT NULL, FK(servers.id) | Server ID |
| alert_type | VARCHAR(64) | NOT NULL | Alert type |
| severity | VARCHAR(32) | NOT NULL | Severity level |
| message | TEXT | NOT NULL | Alert message |
| status | VARCHAR(32) | NOT NULL, DEFAULT 'open' | Status (open/acknowledged/resolved) |
| triggered_at | TIMESTAMP | NOT NULL | Trigger time |
| resolved_at | TIMESTAMP | | Resolution time |

**Indexes**:
- `idx_alerts_server_id` (server_id)
- `idx_alerts_status` (status)
- `idx_alerts_triggered_at` (triggered_at)

**Foreign Keys**:
- server_id → servers(id) ON DELETE CASCADE

**Constraints**:
- `chk_alerts_severity`: severity IN ('info', 'warning', 'error', 'critical')
- `chk_alerts_status`: status IN ('open', 'acknowledged', 'resolved')

---

### 4. Plugin Module

#### 4.1 plugins (Plugins Table)

Information about installed plugins.

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | BIGSERIAL | PRIMARY KEY | Plugin ID |
| name | VARCHAR(128) | NOT NULL, UNIQUE | Plugin name |
| version | VARCHAR(32) | NOT NULL | Version number |
| description | TEXT | | Plugin description |
| author | VARCHAR(128) | | Author name |
| status | VARCHAR(32) | NOT NULL, DEFAULT 'installed' | Status (installed/enabled/disabled/error) |
| config | TEXT | JSON format | Configuration (JSON object) |
| installed_at | TIMESTAMP | NOT NULL | Installation time |
| updated_at | TIMESTAMP | NOT NULL | Update time |

**Indexes**:
- `idx_plugins_status` (status)

**Constraints**:
- `chk_plugins_status`: status IN ('installed', 'enabled', 'disabled', 'error')

---

#### 4.2 plugin_settings (Plugin Settings Table)

Key-value configuration for plugins.

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | BIGSERIAL | PRIMARY KEY | Setting ID |
| plugin_id | BIGINT | NOT NULL, FK(plugins.id) | Plugin ID |
| key | VARCHAR(128) | NOT NULL | Configuration key |
| value | TEXT | | Configuration value |
| created_at | TIMESTAMP | NOT NULL | Creation time |
| updated_at | TIMESTAMP | NOT NULL | Update time |

**Indexes**:
- `idx_plugin_settings_plugin_id` (plugin_id)

**Foreign Keys**:
- plugin_id → plugins(id) ON DELETE CASCADE

**Constraints**:
- UNIQUE(plugin_id, key)

---

### 5. Task Module

#### 5.1 tasks (Tasks Table)

Scheduled and one-time tasks.

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | BIGSERIAL | PRIMARY KEY | Task ID |
| name | VARCHAR(128) | NOT NULL | Task name |
| type | VARCHAR(32) | NOT NULL | Task type (once/scheduled/recurring) |
| target_servers | TEXT | JSON array | Target server ID list |
| command | TEXT | NOT NULL | Command to execute |
| cron_expression | VARCHAR(128) | | Cron expression |
| status | VARCHAR(32) | NOT NULL, DEFAULT 'pending' | Status (pending/running/completed/failed/cancelled) |
| last_run_at | TIMESTAMP | | Last run time |
| next_run_at | TIMESTAMP | | Next run time |
| created_by | BIGINT | NOT NULL, FK(users.id) | Creator user ID |
| created_at | TIMESTAMP | NOT NULL | Creation time |

**Indexes**:
- `idx_tasks_status` (status)
- `idx_tasks_created_by` (created_by)
- `idx_tasks_next_run_at` (next_run_at)

**Foreign Keys**:
- created_by → users(id) ON DELETE CASCADE

**Constraints**:
- `chk_tasks_type`: type IN ('once', 'scheduled', 'recurring')
- `chk_tasks_status`: status IN ('pending', 'running', 'completed', 'failed', 'cancelled')

---

#### 5.2 task_logs (Task Logs Table)

Task execution logs.

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | BIGSERIAL | PRIMARY KEY | Log ID |
| task_id | BIGINT | NOT NULL, FK(tasks.id) | Task ID |
| server_id | BIGINT | NOT NULL, FK(servers.id) | Server ID |
| status | VARCHAR(32) | NOT NULL | Status (running/success/failed/timeout) |
| output | TEXT | | Output content |
| started_at | TIMESTAMP | NOT NULL | Start time |
| finished_at | TIMESTAMP | | Finish time |

**Indexes**:
- `idx_task_logs_task_id` (task_id)
- `idx_task_logs_server_id` (server_id)
- `idx_task_logs_started_at` (started_at)

**Foreign Keys**:
- task_id → tasks(id) ON DELETE CASCADE
- server_id → servers(id) ON DELETE CASCADE

**Constraints**:
- `chk_task_logs_status`: status IN ('running', 'success', 'failed', 'timeout')

---

### 6. System Configuration Module

#### 6.1 settings (System Settings Table)

Global system configuration.

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | BIGSERIAL | PRIMARY KEY | Setting ID |
| key | VARCHAR(128) | NOT NULL, UNIQUE | Configuration key |
| value | TEXT | | Configuration value |
| description | TEXT | | Description |
| updated_at | TIMESTAMP | NOT NULL | Update time |

**Indexes**:
- `idx_settings_key` (key)

---

## Entity Relationship Diagram

```
users ─┬─→ user_tokens
       ├─→ user_logs
       ├─→ ssh_keys
       ├─→ servers
       └─→ tasks

servers ─┬─→ server_group_relations ─→ server_groups
         ├─→ monitor_metrics
         ├─→ alerts
         └─→ task_logs

plugins ─→ plugin_settings

tasks ─→ task_logs
```

---

## Index Strategy

### Primary Indexes

1. **Unique Indexes**: username, email, token and other unique fields
2. **Foreign Key Indexes**: Automatically created for all foreign keys
3. **Status Indexes**: status fields for quick filtering
4. **Time Indexes**: created_at, timestamp and other time fields
5. **Composite Indexes**: (server_id, metric_type, timestamp) for monitoring queries

---

## Data Retention Policy

### Regular Cleanup

1. **Monitor Metrics**: Retain 30 days (configurable)
2. **User Logs**: Retain 90 days
3. **Task Logs**: Retain 30 days
4. **Resolved Alerts**: Retain 30 days
5. **Expired Tokens**: Daily cleanup

---

## Data Backup

### Backup Strategy

1. **Full Backup**: Daily at 2:00 AM
2. **Incremental Backup**: Every 6 hours
3. **Retention Period**: 7 days
4. **Backup Verification**: Weekly recovery tests

---

## Migration Instructions

### Initialize Database

```bash
# PostgreSQL
psql -U nexuspanel -d nexuspanel -f internal/database/migrations/001_initial.sql

# SQLite
sqlite3 nexuspanel.db < internal/database/migrations/001_initial.sql
```

### Default Data

The system automatically creates on first startup:
- Admin account: admin / admin123
- Default roles: admin, operator, user, guest
- System settings: default language, version, etc.

---

## Performance Optimization

1. **Regular Cleanup**: Use scheduled tasks to clean historical data
2. **Table Partitioning**: Consider time-based partitioning for high-volume monitoring data
3. **Read/Write Separation**: Use master-slave replication in production
4. **Connection Pool**: Configure appropriate maximum connections
5. **Slow Query Monitoring**: Regularly check slow query logs

---

## Security Considerations

1. **Password Storage**: Use bcrypt encryption
2. **SSH Keys**: Store private keys encrypted with AES-256
3. **Sensitive Fields**: Do not log passwords and other sensitive information
4. **SQL Injection**: Use parameterized queries
5. **Access Control**: Minimize database user permissions

---

## Appendix

### Field Naming Convention

- Use lowercase letters and underscores
- ID fields uniformly named as `id`
- Foreign key fields named as `table_id`
- Time fields use `_at` suffix

### Data Type Selection

- Primary Key: BIGSERIAL (auto-incrementing integer)
- Strings: VARCHAR(appropriate length)
- Long Text: TEXT
- Time: TIMESTAMP
- JSON: TEXT (using JSON format strings)

---

**End of Document**
