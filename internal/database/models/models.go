package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// JSONMap is a custom type for JSON fields
type JSONMap map[string]interface{}

// Value implements the driver.Valuer interface
func (j JSONMap) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

// Scan implements the sql.Scanner interface
func (j *JSONMap) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, j)
}

// JSONArray is a custom type for JSON arrays
type JSONArray []interface{}

// Value implements the driver.Valuer interface
func (j JSONArray) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

// Scan implements the sql.Scanner interface
func (j *JSONArray) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, j)
}

// ============================================================================
// User Models
// ============================================================================

// User represents a user in the system
type User struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	Username     string     `gorm:"uniqueIndex;size:64;not null" json:"username"`
	Email        string     `gorm:"uniqueIndex;size:255;not null" json:"email"`
	PasswordHash string     `gorm:"size:255;not null" json:"-"`
	Avatar       string     `gorm:"size:512" json:"avatar,omitempty"`
	Role         string     `gorm:"size:32;not null;default:user" json:"role"`
	Status       string     `gorm:"size:32;not null;default:active" json:"status"`
	Language     string     `gorm:"size:16;not null;default:zh-CN" json:"language"`
	CreatedAt    time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"not null" json:"updated_at"`
	LastLoginAt  *time.Time `json:"last_login_at,omitempty"`
}

// TableName specifies the table name for User
func (User) TableName() string {
	return "users"
}

// Role represents a role in the system
type Role struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"uniqueIndex;size:64;not null" json:"name"`
	Description string    `json:"description,omitempty"`
	Permissions JSONArray `gorm:"type:text" json:"permissions"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
}

// TableName specifies the table name for Role
func (Role) TableName() string {
	return "roles"
}

// UserToken represents a user token
type UserToken struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	Token     string    `gorm:"uniqueIndex;size:512;not null" json:"token"`
	Type      string    `gorm:"size:32;not null" json:"type"` // access, refresh, reset, verify
	ExpiresAt time.Time `gorm:"not null;index" json:"expires_at"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	User      User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// TableName specifies the table name for UserToken
func (UserToken) TableName() string {
	return "user_tokens"
}

// UserLog represents a user operation log
type UserLog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	Action    string    `gorm:"size:128;not null;index" json:"action"`
	IP        string    `gorm:"size:64" json:"ip,omitempty"`
	UserAgent string    `json:"user_agent,omitempty"`
	Details   JSONMap   `gorm:"type:text" json:"details,omitempty"`
	CreatedAt time.Time `gorm:"not null;index" json:"created_at"`
	User      User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// TableName specifies the table name for UserLog
func (UserLog) TableName() string {
	return "user_logs"
}

// ============================================================================
// Server Models
// ============================================================================

// SSHKey represents an SSH key
type SSHKey struct {
	ID                   uint      `gorm:"primaryKey" json:"id"`
	Name                 string    `gorm:"size:128;not null" json:"name"`
	PublicKey            string    `gorm:"type:text;not null" json:"public_key"`
	PrivateKeyEncrypted  string    `gorm:"type:text;not null" json:"-"`
	PassphraseEncrypted  string    `gorm:"type:text" json:"-"`
	CreatedAt            time.Time `gorm:"not null" json:"created_at"`
	CreatedBy            uint      `gorm:"not null;index" json:"created_by"`
	Creator              User      `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
}

// TableName specifies the table name for SSHKey
func (SSHKey) TableName() string {
	return "ssh_keys"
}

// ServerGroup represents a server group
type ServerGroup struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	Name        string       `gorm:"size:128;not null" json:"name"`
	Description string       `json:"description,omitempty"`
	ParentID    *uint        `gorm:"index" json:"parent_id,omitempty"`
	CreatedAt   time.Time    `gorm:"not null" json:"created_at"`
	Parent      *ServerGroup `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
}

// TableName specifies the table name for ServerGroup
func (ServerGroup) TableName() string {
	return "server_groups"
}

// Server represents a server
type Server struct {
	ID            uint       `gorm:"primaryKey" json:"id"`
	Name          string     `gorm:"size:128;not null" json:"name"`
	Host          string     `gorm:"size:255;not null;index" json:"host"`
	Port          int        `gorm:"not null;default:22" json:"port"`
	SSHUser       string     `gorm:"size:64;not null" json:"ssh_user"`
	SSHKeyID      *uint      `json:"ssh_key_id,omitempty"`
	Status        string     `gorm:"size:32;not null;default:offline;index" json:"status"`
	OSInfo        JSONMap    `gorm:"type:text" json:"os_info,omitempty"`
	LastHeartbeat *time.Time `json:"last_heartbeat,omitempty"`
	CreatedAt     time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"not null" json:"updated_at"`
	CreatedBy     uint       `gorm:"not null;index" json:"created_by"`
	SSHKey        *SSHKey    `gorm:"foreignKey:SSHKeyID" json:"ssh_key,omitempty"`
	Creator       User       `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
}

// TableName specifies the table name for Server
func (Server) TableName() string {
	return "servers"
}

// ServerGroupRelation represents the relationship between servers and groups
type ServerGroupRelation struct {
	ID       uint        `gorm:"primaryKey" json:"id"`
	ServerID uint        `gorm:"not null;index" json:"server_id"`
	GroupID  uint        `gorm:"not null;index" json:"group_id"`
	Server   Server      `gorm:"foreignKey:ServerID" json:"server,omitempty"`
	Group    ServerGroup `gorm:"foreignKey:GroupID" json:"group,omitempty"`
}

// TableName specifies the table name for ServerGroupRelation
func (ServerGroupRelation) TableName() string {
	return "server_group_relations"
}

// ============================================================================
// Monitoring Models
// ============================================================================

// MonitorMetric represents a monitoring metric
type MonitorMetric struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	ServerID   uint      `gorm:"not null;index:idx_metrics_composite" json:"server_id"`
	MetricType string    `gorm:"size:64;not null;index:idx_metrics_composite" json:"metric_type"`
	Value      float64   `gorm:"not null" json:"value"`
	Tags       JSONMap   `gorm:"type:text" json:"tags,omitempty"`
	Timestamp  time.Time `gorm:"not null;index:idx_metrics_composite" json:"timestamp"`
	Server     Server    `gorm:"foreignKey:ServerID" json:"server,omitempty"`
}

// TableName specifies the table name for MonitorMetric
func (MonitorMetric) TableName() string {
	return "monitor_metrics"
}

// AlertRule represents an alert rule
type AlertRule struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"size:128;not null" json:"name"`
	MetricType string    `gorm:"size:64;not null" json:"metric_type"`
	Condition  string    `gorm:"size:32;not null" json:"condition"` // gt, lt, eq, gte, lte
	Threshold  float64   `gorm:"not null" json:"threshold"`
	Severity   string    `gorm:"size:32;not null" json:"severity"` // info, warning, error, critical
	Enabled    bool      `gorm:"not null;default:true;index" json:"enabled"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at"`
}

// TableName specifies the table name for AlertRule
func (AlertRule) TableName() string {
	return "alert_rules"
}

// Alert represents an alert
type Alert struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	ServerID    uint       `gorm:"not null;index" json:"server_id"`
	AlertType   string     `gorm:"size:64;not null" json:"alert_type"`
	Severity    string     `gorm:"size:32;not null" json:"severity"` // info, warning, error, critical
	Message     string     `gorm:"type:text;not null" json:"message"`
	Status      string     `gorm:"size:32;not null;default:open;index" json:"status"` // open, acknowledged, resolved
	TriggeredAt time.Time  `gorm:"not null;index" json:"triggered_at"`
	ResolvedAt  *time.Time `json:"resolved_at,omitempty"`
	Server      Server     `gorm:"foreignKey:ServerID" json:"server,omitempty"`
}

// TableName specifies the table name for Alert
func (Alert) TableName() string {
	return "alerts"
}

// ============================================================================
// Plugin Models
// ============================================================================

// Plugin represents a plugin
type Plugin struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"uniqueIndex;size:128;not null" json:"name"`
	Version     string    `gorm:"size:32;not null" json:"version"`
	Description string    `json:"description,omitempty"`
	Author      string    `gorm:"size:128" json:"author,omitempty"`
	Status      string    `gorm:"size:32;not null;default:installed;index" json:"status"` // installed, enabled, disabled, error
	Config      JSONMap   `gorm:"type:text" json:"config,omitempty"`
	InstalledAt time.Time `gorm:"not null" json:"installed_at"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at"`
}

// TableName specifies the table name for Plugin
func (Plugin) TableName() string {
	return "plugins"
}

// PluginSetting represents a plugin setting
type PluginSetting struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PluginID  uint      `gorm:"not null;index" json:"plugin_id"`
	Key       string    `gorm:"size:128;not null" json:"key"`
	Value     string    `gorm:"type:text" json:"value,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
	Plugin    Plugin    `gorm:"foreignKey:PluginID" json:"plugin,omitempty"`
}

// TableName specifies the table name for PluginSetting
func (PluginSetting) TableName() string {
	return "plugin_settings"
}

// ============================================================================
// Task Models
// ============================================================================

// Task represents a task
type Task struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	Name           string     `gorm:"size:128;not null" json:"name"`
	Type           string     `gorm:"size:32;not null" json:"type"` // once, scheduled, recurring
	TargetServers  JSONArray  `gorm:"type:text" json:"target_servers"`
	Command        string     `gorm:"type:text;not null" json:"command"`
	CronExpression string     `gorm:"size:128" json:"cron_expression,omitempty"`
	Status         string     `gorm:"size:32;not null;default:pending;index" json:"status"` // pending, running, completed, failed, cancelled
	LastRunAt      *time.Time `json:"last_run_at,omitempty"`
	NextRunAt      *time.Time `gorm:"index" json:"next_run_at,omitempty"`
	CreatedBy      uint       `gorm:"not null;index" json:"created_by"`
	CreatedAt      time.Time  `gorm:"not null" json:"created_at"`
	Creator        User       `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
}

// TableName specifies the table name for Task
func (Task) TableName() string {
	return "tasks"
}

// TaskLog represents a task execution log
type TaskLog struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	TaskID     uint       `gorm:"not null;index" json:"task_id"`
	ServerID   uint       `gorm:"not null;index" json:"server_id"`
	Status     string     `gorm:"size:32;not null" json:"status"` // running, success, failed, timeout
	Output     string     `gorm:"type:text" json:"output,omitempty"`
	StartedAt  time.Time  `gorm:"not null;index" json:"started_at"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	Task       Task       `gorm:"foreignKey:TaskID" json:"task,omitempty"`
	Server     Server     `gorm:"foreignKey:ServerID" json:"server,omitempty"`
}

// TableName specifies the table name for TaskLog
func (TaskLog) TableName() string {
	return "task_logs"
}

// ============================================================================
// System Models
// ============================================================================

// Setting represents a system setting
type Setting struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Key         string    `gorm:"uniqueIndex;size:128;not null" json:"key"`
	Value       string    `gorm:"type:text" json:"value,omitempty"`
	Description string    `json:"description,omitempty"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at"`
}

// TableName specifies the table name for Setting
func (Setting) TableName() string {
	return "settings"
}
