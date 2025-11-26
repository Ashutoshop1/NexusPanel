package config

import (
	"github.com/spf13/viper"
)

// Config represents the application configuration
type Config struct {
	Server      ServerConfig      `mapstructure:"server"`
	Database    DatabaseConfig    `mapstructure:"database"`
	Redis       RedisConfig       `mapstructure:"redis"`
	JWT         JWTConfig         `mapstructure:"jwt"`
	I18n        I18nConfig        `mapstructure:"i18n"`
	Logging     LoggingConfig     `mapstructure:"logging"`
	Security    SecurityConfig    `mapstructure:"security"`
	SSH         SSHConfig         `mapstructure:"ssh"`
	Agent       AgentConfig       `mapstructure:"agent"`
	Monitoring  MonitoringConfig  `mapstructure:"monitoring"`
	Plugins     PluginsConfig     `mapstructure:"plugins"`
	AI          AIConfig          `mapstructure:"ai"`
	Backup      BackupConfig      `mapstructure:"backup"`
	Email       EmailConfig       `mapstructure:"email"`
	WebSocket   WebSocketConfig   `mapstructure:"websocket"`
	FileManager FileManagerConfig `mapstructure:"filemanager"`
	System      SystemConfig      `mapstructure:"system"`
}

// ServerConfig represents server configuration
type ServerConfig struct {
	Host    string        `mapstructure:"host"`
	Port    int           `mapstructure:"port"`
	Mode    string        `mapstructure:"mode"`
	TLS     TLSConfig     `mapstructure:"tls"`
	Session SessionConfig `mapstructure:"session"`
}

// TLSConfig represents TLS configuration
type TLSConfig struct {
	Enabled  bool   `mapstructure:"enabled"`
	CertFile string `mapstructure:"cert_file"`
	KeyFile  string `mapstructure:"key_file"`
}

// SessionConfig represents session configuration
type SessionConfig struct {
	Secret string `mapstructure:"secret"`
	MaxAge int    `mapstructure:"max_age"`
}

// DatabaseConfig represents database configuration
type DatabaseConfig struct {
	Type        string         `mapstructure:"type"`
	SQLite      SQLiteConfig   `mapstructure:"sqlite"`
	Postgres    PostgresConfig `mapstructure:"postgres"`
	AutoMigrate bool           `mapstructure:"auto_migrate"`
}

// SQLiteConfig represents SQLite configuration
type SQLiteConfig struct {
	Path string `mapstructure:"path"`
}

// PostgresConfig represents PostgreSQL configuration
type PostgresConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	Database        string `mapstructure:"database"`
	SSLMode         string `mapstructure:"sslmode"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

// RedisConfig represents Redis configuration
type RedisConfig struct {
	Enabled  bool   `mapstructure:"enabled"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

// JWTConfig represents JWT configuration
type JWTConfig struct {
	Secret             string `mapstructure:"secret"`
	ExpireHours        int    `mapstructure:"expire_hours"`
	RefreshExpireHours int    `mapstructure:"refresh_expire_hours"`
}

// I18nConfig represents internationalization configuration
type I18nConfig struct {
	DefaultLanguage string   `mapstructure:"default_language"`
	Languages       []string `mapstructure:"languages"`
	Path            string   `mapstructure:"path"`
}

// LoggingConfig represents logging configuration
type LoggingConfig struct {
	Level  string        `mapstructure:"level"`
	Format string        `mapstructure:"format"`
	Output string        `mapstructure:"output"`
	File   LogFileConfig `mapstructure:"file"`
}

// LogFileConfig represents log file configuration
type LogFileConfig struct {
	Path       string `mapstructure:"path"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

// SecurityConfig represents security configuration
type SecurityConfig struct {
	CORS      CORSConfig      `mapstructure:"cors"`
	RateLimit RateLimitConfig `mapstructure:"rate_limit"`
	Password  PasswordConfig  `mapstructure:"password"`
}

// CORSConfig represents CORS configuration
type CORSConfig struct {
	Enabled          bool     `mapstructure:"enabled"`
	AllowedOrigins   []string `mapstructure:"allowed_origins"`
	AllowedMethods   []string `mapstructure:"allowed_methods"`
	AllowedHeaders   []string `mapstructure:"allowed_headers"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
	MaxAge           int      `mapstructure:"max_age"`
}

// RateLimitConfig represents rate limit configuration
type RateLimitConfig struct {
	Enabled           bool `mapstructure:"enabled"`
	RequestsPerMinute int  `mapstructure:"requests_per_minute"`
}

// PasswordConfig represents password policy configuration
type PasswordConfig struct {
	MinLength        int  `mapstructure:"min_length"`
	RequireUppercase bool `mapstructure:"require_uppercase"`
	RequireLowercase bool `mapstructure:"require_lowercase"`
	RequireNumber    bool `mapstructure:"require_number"`
	RequireSpecial   bool `mapstructure:"require_special"`
}

// SSHConfig represents SSH configuration
type SSHConfig struct {
	EncryptionKey string `mapstructure:"encryption_key"`
	Timeout       int    `mapstructure:"timeout"`
	KeepAlive     int    `mapstructure:"keep_alive"`
}

// AgentConfig represents agent configuration
type AgentConfig struct {
	Protocol          string    `mapstructure:"protocol"`
	Port              int       `mapstructure:"port"`
	HeartbeatInterval int       `mapstructure:"heartbeat_interval"`
	Timeout           int       `mapstructure:"timeout"`
	TLS               TLSConfig `mapstructure:"tls"`
}

// MonitoringConfig represents monitoring configuration
type MonitoringConfig struct {
	CollectionInterval int `mapstructure:"collection_interval"`
	RetentionDays      int `mapstructure:"retention_days"`
	AlertCheckInterval int `mapstructure:"alert_check_interval"`
}

// PluginsConfig represents plugins configuration
type PluginsConfig struct {
	Directory string `mapstructure:"directory"`
	HotReload bool   `mapstructure:"hot_reload"`
	Timeout   int    `mapstructure:"timeout"`
}

// AIConfig represents AI configuration
type AIConfig struct {
	Enabled     bool    `mapstructure:"enabled"`
	Provider    string  `mapstructure:"provider"`
	APIKey      string  `mapstructure:"api_key"`
	APIEndpoint string  `mapstructure:"api_endpoint"`
	Model       string  `mapstructure:"model"`
	MaxTokens   int     `mapstructure:"max_tokens"`
	Temperature float64 `mapstructure:"temperature"`
}

// BackupConfig represents backup configuration
type BackupConfig struct {
	Enabled       bool   `mapstructure:"enabled"`
	Schedule      string `mapstructure:"schedule"`
	RetentionDays int    `mapstructure:"retention_days"`
	StoragePath   string `mapstructure:"storage_path"`
}

// EmailConfig represents email configuration
type EmailConfig struct {
	Enabled  bool   `mapstructure:"enabled"`
	SMTPHost string `mapstructure:"smtp_host"`
	SMTPPort int    `mapstructure:"smtp_port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	From     string `mapstructure:"from"`
	UseTLS   bool   `mapstructure:"use_tls"`
}

// WebSocketConfig represents WebSocket configuration
type WebSocketConfig struct {
	Path              string `mapstructure:"path"`
	ReadBufferSize    int    `mapstructure:"read_buffer_size"`
	WriteBufferSize   int    `mapstructure:"write_buffer_size"`
	HeartbeatInterval int    `mapstructure:"heartbeat_interval"`
}

// FileManagerConfig represents file manager configuration
type FileManagerConfig struct {
	MaxUploadSize       int64    `mapstructure:"max_upload_size"`
	AllowedExtensions   []string `mapstructure:"allowed_extensions"`
	ForbiddenExtensions []string `mapstructure:"forbidden_extensions"`
	UploadDir           string   `mapstructure:"upload_dir"`
}

// SystemConfig represents system configuration
type SystemConfig struct {
	Admin      AdminConfig      `mapstructure:"admin"`
	Pagination PaginationConfig `mapstructure:"pagination"`
	Timezone   string           `mapstructure:"timezone"`
}

// AdminConfig represents admin user configuration
type AdminConfig struct {
	Username string `mapstructure:"username"`
	Email    string `mapstructure:"email"`
	Password string `mapstructure:"password"`
}

// PaginationConfig represents pagination configuration
type PaginationConfig struct {
	DefaultPageSize int `mapstructure:"default_page_size"`
	MaxPageSize     int `mapstructure:"max_page_size"`
}

// Load loads configuration from file
func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")

	// Set defaults
	setDefaults()

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		// If config file not found, use defaults
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	// Allow environment variables
	viper.AutomaticEnv()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// setDefaults sets default configuration values
func setDefaults() {
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.mode", "debug")

	viper.SetDefault("database.type", "sqlite")
	viper.SetDefault("database.sqlite.path", "./nexuspanel.db")
	viper.SetDefault("database.auto_migrate", true)

	viper.SetDefault("i18n.default_language", "zh-CN")
	viper.SetDefault("i18n.languages", []string{"zh-CN", "en-US"})

	viper.SetDefault("logging.level", "info")
	viper.SetDefault("logging.format", "console")
	viper.SetDefault("logging.output", "stdout")
}
