-- NexusPanel Initial Database Schema
-- Supports both PostgreSQL and SQLite
-- Version: 001

-- ============================================================================
-- User Management Tables
-- ============================================================================

-- Users table
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(64) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    avatar VARCHAR(512),
    role VARCHAR(32) NOT NULL DEFAULT 'user',
    status VARCHAR(32) NOT NULL DEFAULT 'active',
    language VARCHAR(16) NOT NULL DEFAULT 'zh-CN',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_login_at TIMESTAMP,
    CONSTRAINT chk_users_role CHECK (role IN ('admin', 'user', 'guest', 'operator')),
    CONSTRAINT chk_users_status CHECK (status IN ('active', 'inactive', 'locked', 'pending'))
);

CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_status ON users(status);

-- Roles table
CREATE TABLE IF NOT EXISTS roles (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(64) NOT NULL UNIQUE,
    description TEXT,
    permissions TEXT, -- JSON format
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_roles_name ON roles(name);

-- User tokens table
CREATE TABLE IF NOT EXISTS user_tokens (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    token VARCHAR(512) NOT NULL UNIQUE,
    type VARCHAR(32) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT chk_user_tokens_type CHECK (type IN ('access', 'refresh', 'reset', 'verify'))
);

CREATE INDEX idx_user_tokens_user_id ON user_tokens(user_id);
CREATE INDEX idx_user_tokens_token ON user_tokens(token);
CREATE INDEX idx_user_tokens_expires_at ON user_tokens(expires_at);

-- User operation logs table
CREATE TABLE IF NOT EXISTS user_logs (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    action VARCHAR(128) NOT NULL,
    ip VARCHAR(64),
    user_agent TEXT,
    details TEXT, -- JSON format
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_user_logs_user_id ON user_logs(user_id);
CREATE INDEX idx_user_logs_action ON user_logs(action);
CREATE INDEX idx_user_logs_created_at ON user_logs(created_at);

-- ============================================================================
-- Server Management Tables
-- ============================================================================

-- SSH keys table
CREATE TABLE IF NOT EXISTS ssh_keys (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    public_key TEXT NOT NULL,
    private_key_encrypted TEXT NOT NULL,
    passphrase_encrypted TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by BIGINT NOT NULL,
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_ssh_keys_created_by ON ssh_keys(created_by);

-- Server groups table
CREATE TABLE IF NOT EXISTS server_groups (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    description TEXT,
    parent_id BIGINT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (parent_id) REFERENCES server_groups(id) ON DELETE SET NULL
);

CREATE INDEX idx_server_groups_parent_id ON server_groups(parent_id);

-- Servers table
CREATE TABLE IF NOT EXISTS servers (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    host VARCHAR(255) NOT NULL,
    port INTEGER NOT NULL DEFAULT 22,
    ssh_user VARCHAR(64) NOT NULL,
    ssh_key_id BIGINT,
    status VARCHAR(32) NOT NULL DEFAULT 'offline',
    os_info TEXT, -- JSON format
    last_heartbeat TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by BIGINT NOT NULL,
    FOREIGN KEY (ssh_key_id) REFERENCES ssh_keys(id) ON DELETE SET NULL,
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT chk_servers_status CHECK (status IN ('online', 'offline', 'error', 'maintenance'))
);

CREATE INDEX idx_servers_host ON servers(host);
CREATE INDEX idx_servers_status ON servers(status);
CREATE INDEX idx_servers_created_by ON servers(created_by);

-- Server group relations table
CREATE TABLE IF NOT EXISTS server_group_relations (
    id BIGSERIAL PRIMARY KEY,
    server_id BIGINT NOT NULL,
    group_id BIGINT NOT NULL,
    FOREIGN KEY (server_id) REFERENCES servers(id) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES server_groups(id) ON DELETE CASCADE,
    UNIQUE(server_id, group_id)
);

CREATE INDEX idx_server_group_relations_server_id ON server_group_relations(server_id);
CREATE INDEX idx_server_group_relations_group_id ON server_group_relations(group_id);

-- ============================================================================
-- Monitoring Tables
-- ============================================================================

-- Monitor metrics table (time series data)
CREATE TABLE IF NOT EXISTS monitor_metrics (
    id BIGSERIAL PRIMARY KEY,
    server_id BIGINT NOT NULL,
    metric_type VARCHAR(64) NOT NULL,
    value DOUBLE PRECISION NOT NULL,
    tags TEXT, -- JSON format
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (server_id) REFERENCES servers(id) ON DELETE CASCADE
);

CREATE INDEX idx_monitor_metrics_server_id ON monitor_metrics(server_id);
CREATE INDEX idx_monitor_metrics_metric_type ON monitor_metrics(metric_type);
CREATE INDEX idx_monitor_metrics_timestamp ON monitor_metrics(timestamp);
CREATE INDEX idx_monitor_metrics_composite ON monitor_metrics(server_id, metric_type, timestamp);

-- Alert rules table
CREATE TABLE IF NOT EXISTS alert_rules (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    metric_type VARCHAR(64) NOT NULL,
    condition VARCHAR(32) NOT NULL,
    threshold DOUBLE PRECISION NOT NULL,
    severity VARCHAR(32) NOT NULL,
    enabled BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT chk_alert_rules_condition CHECK (condition IN ('gt', 'lt', 'eq', 'gte', 'lte')),
    CONSTRAINT chk_alert_rules_severity CHECK (severity IN ('info', 'warning', 'error', 'critical'))
);

CREATE INDEX idx_alert_rules_enabled ON alert_rules(enabled);

-- Alerts table
CREATE TABLE IF NOT EXISTS alerts (
    id BIGSERIAL PRIMARY KEY,
    server_id BIGINT NOT NULL,
    alert_type VARCHAR(64) NOT NULL,
    severity VARCHAR(32) NOT NULL,
    message TEXT NOT NULL,
    status VARCHAR(32) NOT NULL DEFAULT 'open',
    triggered_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    resolved_at TIMESTAMP,
    FOREIGN KEY (server_id) REFERENCES servers(id) ON DELETE CASCADE,
    CONSTRAINT chk_alerts_severity CHECK (severity IN ('info', 'warning', 'error', 'critical')),
    CONSTRAINT chk_alerts_status CHECK (status IN ('open', 'acknowledged', 'resolved'))
);

CREATE INDEX idx_alerts_server_id ON alerts(server_id);
CREATE INDEX idx_alerts_status ON alerts(status);
CREATE INDEX idx_alerts_triggered_at ON alerts(triggered_at);

-- ============================================================================
-- Plugin Tables
-- ============================================================================

-- Plugins table
CREATE TABLE IF NOT EXISTS plugins (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL UNIQUE,
    version VARCHAR(32) NOT NULL,
    description TEXT,
    author VARCHAR(128),
    status VARCHAR(32) NOT NULL DEFAULT 'installed',
    config TEXT, -- JSON format
    installed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT chk_plugins_status CHECK (status IN ('installed', 'enabled', 'disabled', 'error'))
);

CREATE INDEX idx_plugins_status ON plugins(status);

-- Plugin settings table
CREATE TABLE IF NOT EXISTS plugin_settings (
    id BIGSERIAL PRIMARY KEY,
    plugin_id BIGINT NOT NULL,
    key VARCHAR(128) NOT NULL,
    value TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (plugin_id) REFERENCES plugins(id) ON DELETE CASCADE,
    UNIQUE(plugin_id, key)
);

CREATE INDEX idx_plugin_settings_plugin_id ON plugin_settings(plugin_id);

-- ============================================================================
-- Task Tables
-- ============================================================================

-- Tasks table
CREATE TABLE IF NOT EXISTS tasks (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    type VARCHAR(32) NOT NULL,
    target_servers TEXT, -- JSON array
    command TEXT NOT NULL,
    cron_expression VARCHAR(128),
    status VARCHAR(32) NOT NULL DEFAULT 'pending',
    last_run_at TIMESTAMP,
    next_run_at TIMESTAMP,
    created_by BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT chk_tasks_type CHECK (type IN ('once', 'scheduled', 'recurring')),
    CONSTRAINT chk_tasks_status CHECK (status IN ('pending', 'running', 'completed', 'failed', 'cancelled'))
);

CREATE INDEX idx_tasks_status ON tasks(status);
CREATE INDEX idx_tasks_created_by ON tasks(created_by);
CREATE INDEX idx_tasks_next_run_at ON tasks(next_run_at);

-- Task logs table
CREATE TABLE IF NOT EXISTS task_logs (
    id BIGSERIAL PRIMARY KEY,
    task_id BIGINT NOT NULL,
    server_id BIGINT NOT NULL,
    status VARCHAR(32) NOT NULL,
    output TEXT,
    started_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    finished_at TIMESTAMP,
    FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE,
    FOREIGN KEY (server_id) REFERENCES servers(id) ON DELETE CASCADE,
    CONSTRAINT chk_task_logs_status CHECK (status IN ('running', 'success', 'failed', 'timeout'))
);

CREATE INDEX idx_task_logs_task_id ON task_logs(task_id);
CREATE INDEX idx_task_logs_server_id ON task_logs(server_id);
CREATE INDEX idx_task_logs_started_at ON task_logs(started_at);

-- ============================================================================
-- System Configuration Tables
-- ============================================================================

-- Settings table
CREATE TABLE IF NOT EXISTS settings (
    id BIGSERIAL PRIMARY KEY,
    key VARCHAR(128) NOT NULL UNIQUE,
    value TEXT,
    description TEXT,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_settings_key ON settings(key);

-- ============================================================================
-- Insert Default Data
-- ============================================================================

-- Insert default admin user (password: admin123)
-- Note: This is a bcrypt hash of "admin123"
INSERT INTO users (username, email, password_hash, role, status, language) 
VALUES ('admin', 'admin@nexuspanel.com', '$2a$10$N9qo8uLOickgx2ZMRZoMye1J8EJ.5YyXS0qPVYQ2vZlZyKxbLQKUC', 'admin', 'active', 'zh-CN')
ON CONFLICT (username) DO NOTHING;

-- Insert default roles
INSERT INTO roles (name, description, permissions) 
VALUES 
    ('admin', 'Administrator with full access', '["*"]'),
    ('operator', 'Operator with server management access', '["server:*", "monitor:read", "task:*"]'),
    ('user', 'Regular user with read access', '["server:read", "monitor:read"]'),
    ('guest', 'Guest with minimal access', '["server:read"]')
ON CONFLICT (name) DO NOTHING;

-- Insert default system settings
INSERT INTO settings (key, value, description) 
VALUES 
    ('system.name', 'NexusPanel', 'System name'),
    ('system.version', '0.1.0', 'System version'),
    ('system.language', 'zh-CN', 'Default system language'),
    ('maintenance.enabled', 'false', 'Maintenance mode enabled'),
    ('registration.enabled', 'true', 'User registration enabled')
ON CONFLICT (key) DO NOTHING;
