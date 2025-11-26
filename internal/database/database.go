package database

import (
	"fmt"
	"time"

	"github.com/2670044605/NexusPanel/internal/database/models"
	"github.com/2670044605/NexusPanel/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Init initializes the database connection
func Init(cfg config.DatabaseConfig) (*gorm.DB, error) {
	var dialector gorm.Dialector

	switch cfg.Type {
	case "postgres":
		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			cfg.Postgres.Host,
			cfg.Postgres.Port,
			cfg.Postgres.User,
			cfg.Postgres.Password,
			cfg.Postgres.Database,
			cfg.Postgres.SSLMode,
		)
		dialector = postgres.Open(dsn)

	case "sqlite":
		dialector = sqlite.Open(cfg.SQLite.Path)

	default:
		return nil, fmt.Errorf("unsupported database type: %s", cfg.Type)
	}

	// Configure GORM
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	// Open database connection
	db, err := gorm.Open(dialector, gormConfig)
	if err != nil {
		return nil, err
	}

	// Get underlying *sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Set connection pool settings for PostgreSQL
	if cfg.Type == "postgres" {
		sqlDB.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
		sqlDB.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
		sqlDB.SetConnMaxLifetime(time.Duration(cfg.Postgres.ConnMaxLifetime) * time.Second)
	}

	return db, nil
}

// AutoMigrate runs automatic database migrations
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		// User management models
		&models.User{},
		&models.Role{},
		&models.UserToken{},
		&models.UserLog{},

		// Server management models
		&models.SSHKey{},
		&models.ServerGroup{},
		&models.Server{},
		&models.ServerGroupRelation{},

		// Monitoring models
		&models.MonitorMetric{},
		&models.AlertRule{},
		&models.Alert{},

		// Plugin models
		&models.Plugin{},
		&models.PluginSetting{},

		// Task models
		&models.Task{},
		&models.TaskLog{},

		// System models
		&models.Setting{},
	)
}
