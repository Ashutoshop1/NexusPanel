# NexusPanel 数据库设计文档

## 文档版本

- **版本**: v1.0
- **创建日期**: 2024
- **最后更新**: 2024
- **数据库类型**: PostgreSQL 13+ / SQLite 3.x

---

## 概述

NexusPanel 使用关系型数据库存储系统数据，支持 PostgreSQL（生产环境推荐）和 SQLite（开发环境/轻量部署）。数据库设计遵循第三范式，确保数据一致性和完整性。

### 设计原则

1. **模块化设计**: 按功能模块划分表结构
2. **数据完整性**: 使用外键约束保证引用完整性
3. **性能优化**: 为常用查询字段添加索引
4. **可扩展性**: 使用 JSON 字段存储灵活配置
5. **审计追踪**: 记录关键操作的时间戳和操作人

---

## 表结构设计

### 1. 用户管理模块

#### 1.1 users (用户表)

存储系统用户的基本信息和认证数据。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | BIGSERIAL | PRIMARY KEY | 用户 ID |
| username | VARCHAR(64) | NOT NULL, UNIQUE | 用户名 |
| email | VARCHAR(255) | NOT NULL, UNIQUE | 邮箱地址 |
| password_hash | VARCHAR(255) | NOT NULL | 密码哈希值（bcrypt） |
| avatar | VARCHAR(512) | | 头像 URL |
| role | VARCHAR(32) | NOT NULL, DEFAULT 'user' | 角色（admin/operator/user/guest） |
| status | VARCHAR(32) | NOT NULL, DEFAULT 'active' | 状态（active/inactive/locked/pending） |
| language | VARCHAR(16) | NOT NULL, DEFAULT 'zh-CN' | 语言偏好（zh-CN/en-US） |
| created_at | TIMESTAMP | NOT NULL | 创建时间 |
| updated_at | TIMESTAMP | NOT NULL | 更新时间 |
| last_login_at | TIMESTAMP | | 最后登录时间 |

**索引**:
- `idx_users_username` (username)
- `idx_users_email` (email)
- `idx_users_status` (status)

**约束**:
- `chk_users_role`: role IN ('admin', 'user', 'guest', 'operator')
- `chk_users_status`: status IN ('active', 'inactive', 'locked', 'pending')

---

#### 1.2 roles (角色表)

定义系统角色和权限。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | BIGSERIAL | PRIMARY KEY | 角色 ID |
| name | VARCHAR(64) | NOT NULL, UNIQUE | 角色名称 |
| description | TEXT | | 角色描述 |
| permissions | TEXT | JSON 格式 | 权限列表（JSON 数组） |
| created_at | TIMESTAMP | NOT NULL | 创建时间 |

**索引**:
- `idx_roles_name` (name)

**权限格式示例**:
```json
["server:read", "server:write", "monitor:read", "task:execute"]
```

---

#### 1.3 user_tokens (用户令牌表)

存储用户认证令牌（JWT、重置密码令牌等）。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | BIGSERIAL | PRIMARY KEY | 令牌 ID |
| user_id | BIGINT | NOT NULL, FK(users.id) | 用户 ID |
| token | VARCHAR(512) | NOT NULL, UNIQUE | 令牌字符串 |
| type | VARCHAR(32) | NOT NULL | 令牌类型（access/refresh/reset/verify） |
| expires_at | TIMESTAMP | NOT NULL | 过期时间 |
| created_at | TIMESTAMP | NOT NULL | 创建时间 |

**索引**:
- `idx_user_tokens_user_id` (user_id)
- `idx_user_tokens_token` (token)
- `idx_user_tokens_expires_at` (expires_at)

**外键**:
- user_id → users(id) ON DELETE CASCADE

---

#### 1.4 user_logs (用户操作日志)

记录用户操作审计日志。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | BIGSERIAL | PRIMARY KEY | 日志 ID |
| user_id | BIGINT | NOT NULL, FK(users.id) | 用户 ID |
| action | VARCHAR(128) | NOT NULL | 操作类型 |
| ip | VARCHAR(64) | | IP 地址 |
| user_agent | TEXT | | 用户代理字符串 |
| details | TEXT | JSON 格式 | 详细信息（JSON 对象） |
| created_at | TIMESTAMP | NOT NULL | 创建时间 |

**索引**:
- `idx_user_logs_user_id` (user_id)
- `idx_user_logs_action` (action)
- `idx_user_logs_created_at` (created_at)

**外键**:
- user_id → users(id) ON DELETE CASCADE

---

### 2. 服务器管理模块

#### 2.1 ssh_keys (SSH 密钥表)

存储 SSH 密钥对（私钥已加密）。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | BIGSERIAL | PRIMARY KEY | 密钥 ID |
| name | VARCHAR(128) | NOT NULL | 密钥名称 |
| public_key | TEXT | NOT NULL | 公钥 |
| private_key_encrypted | TEXT | NOT NULL | 加密的私钥 |
| passphrase_encrypted | TEXT | | 加密的密码短语 |
| created_at | TIMESTAMP | NOT NULL | 创建时间 |
| created_by | BIGINT | NOT NULL, FK(users.id) | 创建人 |

**索引**:
- `idx_ssh_keys_created_by` (created_by)

**外键**:
- created_by → users(id) ON DELETE CASCADE

---

#### 2.2 server_groups (服务器分组表)

服务器分组，支持层级结构。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | BIGSERIAL | PRIMARY KEY | 分组 ID |
| name | VARCHAR(128) | NOT NULL | 分组名称 |
| description | TEXT | | 分组描述 |
| parent_id | BIGINT | FK(server_groups.id) | 父分组 ID |
| created_at | TIMESTAMP | NOT NULL | 创建时间 |

**索引**:
- `idx_server_groups_parent_id` (parent_id)

**外键**:
- parent_id → server_groups(id) ON DELETE SET NULL

---

#### 2.3 servers (服务器表)

存储服务器基本信息和连接配置。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | BIGSERIAL | PRIMARY KEY | 服务器 ID |
| name | VARCHAR(128) | NOT NULL | 服务器名称 |
| host | VARCHAR(255) | NOT NULL | 主机地址 |
| port | INTEGER | NOT NULL, DEFAULT 22 | SSH 端口 |
| ssh_user | VARCHAR(64) | NOT NULL | SSH 用户名 |
| ssh_key_id | BIGINT | FK(ssh_keys.id) | SSH 密钥 ID |
| status | VARCHAR(32) | NOT NULL, DEFAULT 'offline' | 状态（online/offline/error/maintenance） |
| os_info | TEXT | JSON 格式 | 操作系统信息 |
| last_heartbeat | TIMESTAMP | | 最后心跳时间 |
| created_at | TIMESTAMP | NOT NULL | 创建时间 |
| updated_at | TIMESTAMP | NOT NULL | 更新时间 |
| created_by | BIGINT | NOT NULL, FK(users.id) | 创建人 |

**索引**:
- `idx_servers_host` (host)
- `idx_servers_status` (status)
- `idx_servers_created_by` (created_by)

**外键**:
- ssh_key_id → ssh_keys(id) ON DELETE SET NULL
- created_by → users(id) ON DELETE CASCADE

**约束**:
- `chk_servers_status`: status IN ('online', 'offline', 'error', 'maintenance')

---

#### 2.4 server_group_relations (服务器分组关系表)

服务器与分组的多对多关系。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | BIGSERIAL | PRIMARY KEY | 关系 ID |
| server_id | BIGINT | NOT NULL, FK(servers.id) | 服务器 ID |
| group_id | BIGINT | NOT NULL, FK(server_groups.id) | 分组 ID |

**索引**:
- `idx_server_group_relations_server_id` (server_id)
- `idx_server_group_relations_group_id` (group_id)

**外键**:
- server_id → servers(id) ON DELETE CASCADE
- group_id → server_groups(id) ON DELETE CASCADE

**约束**:
- UNIQUE(server_id, group_id)

---

### 3. 监控模块

#### 3.1 monitor_metrics (监控指标表)

存储时序监控数据。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | BIGSERIAL | PRIMARY KEY | 指标 ID |
| server_id | BIGINT | NOT NULL, FK(servers.id) | 服务器 ID |
| metric_type | VARCHAR(64) | NOT NULL | 指标类型（cpu/memory/disk/network） |
| value | DOUBLE PRECISION | NOT NULL | 指标值 |
| tags | TEXT | JSON 格式 | 标签（JSON 对象） |
| timestamp | TIMESTAMP | NOT NULL | 时间戳 |

**索引**:
- `idx_monitor_metrics_server_id` (server_id)
- `idx_monitor_metrics_metric_type` (metric_type)
- `idx_monitor_metrics_timestamp` (timestamp)
- `idx_monitor_metrics_composite` (server_id, metric_type, timestamp)

**外键**:
- server_id → servers(id) ON DELETE CASCADE

---

#### 3.2 alert_rules (告警规则表)

定义告警规则。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | BIGSERIAL | PRIMARY KEY | 规则 ID |
| name | VARCHAR(128) | NOT NULL | 规则名称 |
| metric_type | VARCHAR(64) | NOT NULL | 监控指标类型 |
| condition | VARCHAR(32) | NOT NULL | 条件（gt/lt/eq/gte/lte） |
| threshold | DOUBLE PRECISION | NOT NULL | 阈值 |
| severity | VARCHAR(32) | NOT NULL | 严重程度（info/warning/error/critical） |
| enabled | BOOLEAN | NOT NULL, DEFAULT TRUE | 是否启用 |
| created_at | TIMESTAMP | NOT NULL | 创建时间 |

**索引**:
- `idx_alert_rules_enabled` (enabled)

**约束**:
- `chk_alert_rules_condition`: condition IN ('gt', 'lt', 'eq', 'gte', 'lte')
- `chk_alert_rules_severity`: severity IN ('info', 'warning', 'error', 'critical')

---

#### 3.3 alerts (告警表)

存储触发的告警。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | BIGSERIAL | PRIMARY KEY | 告警 ID |
| server_id | BIGINT | NOT NULL, FK(servers.id) | 服务器 ID |
| alert_type | VARCHAR(64) | NOT NULL | 告警类型 |
| severity | VARCHAR(32) | NOT NULL | 严重程度 |
| message | TEXT | NOT NULL | 告警消息 |
| status | VARCHAR(32) | NOT NULL, DEFAULT 'open' | 状态（open/acknowledged/resolved） |
| triggered_at | TIMESTAMP | NOT NULL | 触发时间 |
| resolved_at | TIMESTAMP | | 解决时间 |

**索引**:
- `idx_alerts_server_id` (server_id)
- `idx_alerts_status` (status)
- `idx_alerts_triggered_at` (triggered_at)

**外键**:
- server_id → servers(id) ON DELETE CASCADE

**约束**:
- `chk_alerts_severity`: severity IN ('info', 'warning', 'error', 'critical')
- `chk_alerts_status`: status IN ('open', 'acknowledged', 'resolved')

---

### 4. 插件模块

#### 4.1 plugins (插件表)

已安装的插件信息。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | BIGSERIAL | PRIMARY KEY | 插件 ID |
| name | VARCHAR(128) | NOT NULL, UNIQUE | 插件名称 |
| version | VARCHAR(32) | NOT NULL | 版本号 |
| description | TEXT | | 插件描述 |
| author | VARCHAR(128) | | 作者 |
| status | VARCHAR(32) | NOT NULL, DEFAULT 'installed' | 状态（installed/enabled/disabled/error） |
| config | TEXT | JSON 格式 | 配置（JSON 对象） |
| installed_at | TIMESTAMP | NOT NULL | 安装时间 |
| updated_at | TIMESTAMP | NOT NULL | 更新时间 |

**索引**:
- `idx_plugins_status` (status)

**约束**:
- `chk_plugins_status`: status IN ('installed', 'enabled', 'disabled', 'error')

---

#### 4.2 plugin_settings (插件配置表)

插件的键值配置。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | BIGSERIAL | PRIMARY KEY | 配置 ID |
| plugin_id | BIGINT | NOT NULL, FK(plugins.id) | 插件 ID |
| key | VARCHAR(128) | NOT NULL | 配置键 |
| value | TEXT | | 配置值 |
| created_at | TIMESTAMP | NOT NULL | 创建时间 |
| updated_at | TIMESTAMP | NOT NULL | 更新时间 |

**索引**:
- `idx_plugin_settings_plugin_id` (plugin_id)

**外键**:
- plugin_id → plugins(id) ON DELETE CASCADE

**约束**:
- UNIQUE(plugin_id, key)

---

### 5. 任务模块

#### 5.1 tasks (任务表)

计划任务和一次性任务。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | BIGSERIAL | PRIMARY KEY | 任务 ID |
| name | VARCHAR(128) | NOT NULL | 任务名称 |
| type | VARCHAR(32) | NOT NULL | 任务类型（once/scheduled/recurring） |
| target_servers | TEXT | JSON 数组 | 目标服务器 ID 列表 |
| command | TEXT | NOT NULL | 执行命令 |
| cron_expression | VARCHAR(128) | | Cron 表达式 |
| status | VARCHAR(32) | NOT NULL, DEFAULT 'pending' | 状态（pending/running/completed/failed/cancelled） |
| last_run_at | TIMESTAMP | | 上次运行时间 |
| next_run_at | TIMESTAMP | | 下次运行时间 |
| created_by | BIGINT | NOT NULL, FK(users.id) | 创建人 |
| created_at | TIMESTAMP | NOT NULL | 创建时间 |

**索引**:
- `idx_tasks_status` (status)
- `idx_tasks_created_by` (created_by)
- `idx_tasks_next_run_at` (next_run_at)

**外键**:
- created_by → users(id) ON DELETE CASCADE

**约束**:
- `chk_tasks_type`: type IN ('once', 'scheduled', 'recurring')
- `chk_tasks_status`: status IN ('pending', 'running', 'completed', 'failed', 'cancelled')

---

#### 5.2 task_logs (任务日志表)

任务执行日志。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | BIGSERIAL | PRIMARY KEY | 日志 ID |
| task_id | BIGINT | NOT NULL, FK(tasks.id) | 任务 ID |
| server_id | BIGINT | NOT NULL, FK(servers.id) | 服务器 ID |
| status | VARCHAR(32) | NOT NULL | 状态（running/success/failed/timeout） |
| output | TEXT | | 输出内容 |
| started_at | TIMESTAMP | NOT NULL | 开始时间 |
| finished_at | TIMESTAMP | | 完成时间 |

**索引**:
- `idx_task_logs_task_id` (task_id)
- `idx_task_logs_server_id` (server_id)
- `idx_task_logs_started_at` (started_at)

**外键**:
- task_id → tasks(id) ON DELETE CASCADE
- server_id → servers(id) ON DELETE CASCADE

**约束**:
- `chk_task_logs_status`: status IN ('running', 'success', 'failed', 'timeout')

---

### 6. 系统配置模块

#### 6.1 settings (系统设置表)

系统全局配置。

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | BIGSERIAL | PRIMARY KEY | 设置 ID |
| key | VARCHAR(128) | NOT NULL, UNIQUE | 配置键 |
| value | TEXT | | 配置值 |
| description | TEXT | | 描述 |
| updated_at | TIMESTAMP | NOT NULL | 更新时间 |

**索引**:
- `idx_settings_key` (key)

---

## 数据关系图

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

## 索引策略

### 主要索引

1. **唯一索引**: username, email, token 等唯一字段
2. **外键索引**: 所有外键字段自动建立索引
3. **状态索引**: status 字段用于快速过滤
4. **时间索引**: created_at, timestamp 等时间字段
5. **组合索引**: (server_id, metric_type, timestamp) 用于监控数据查询

---

## 数据保留策略

### 定期清理

1. **监控指标**: 保留 30 天（可配置）
2. **用户日志**: 保留 90 天
3. **任务日志**: 保留 30 天
4. **已解决的告警**: 保留 30 天
5. **过期令牌**: 每日清理

---

## 数据备份

### 备份策略

1. **全量备份**: 每日凌晨 2:00
2. **增量备份**: 每 6 小时一次
3. **保留周期**: 7 天
4. **备份验证**: 每周执行恢复测试

---

## 迁移说明

### 初始化数据库

```bash
# PostgreSQL
psql -U nexuspanel -d nexuspanel -f internal/database/migrations/001_initial.sql

# SQLite
sqlite3 nexuspanel.db < internal/database/migrations/001_initial.sql
```

### 默认数据

系统会在首次启动时自动创建：
- 管理员账户: admin / admin123
- 默认角色: admin, operator, user, guest
- 系统配置: 默认语言、版本等

---

## 性能优化建议

1. **定期清理过期数据**: 使用定时任务清理历史数据
2. **分表策略**: 监控数据量大时考虑按时间分表
3. **读写分离**: 生产环境使用主从复制
4. **连接池配置**: 合理设置最大连接数
5. **慢查询监控**: 定期检查慢查询日志

---

## 安全考虑

1. **密码存储**: 使用 bcrypt 加密
2. **SSH 密钥**: 使用 AES-256 加密存储私钥
3. **敏感字段**: 不在日志中记录密码等敏感信息
4. **SQL 注入防护**: 使用参数化查询
5. **访问控制**: 数据库用户权限最小化

---

## 附录

### 字段命名规范

- 使用小写字母和下划线
- ID 字段统一命名为 `id`
- 外键字段命名为 `关联表_id`
- 时间字段使用 `_at` 后缀

### 数据类型选择

- 主键: BIGSERIAL (自增长整型)
- 字符串: VARCHAR(适当长度)
- 长文本: TEXT
- 时间: TIMESTAMP
- JSON: TEXT (使用 JSON 格式字符串)

---

**文档结束**
