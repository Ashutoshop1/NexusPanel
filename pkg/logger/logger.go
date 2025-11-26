package logger

import (
	"os"

	"github.com/2670044605/NexusPanel/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.SugaredLogger

// Init initializes the logger
func Init(cfg config.LoggingConfig) error {
	var zapConfig zap.Config

	// Set log level
	level := zapcore.InfoLevel
	switch cfg.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	}

	// Set log format
	if cfg.Format == "json" {
		zapConfig = zap.NewProductionConfig()
	} else {
		zapConfig = zap.NewDevelopmentConfig()
	}

	zapConfig.Level = zap.NewAtomicLevelAt(level)

	// Set output
	if cfg.Output == "file" && cfg.File.Path != "" {
		zapConfig.OutputPaths = []string{cfg.File.Path}
		zapConfig.ErrorOutputPaths = []string{cfg.File.Path}
	}

	// Build logger
	logger, err := zapConfig.Build()
	if err != nil {
		return err
	}

	log = logger.Sugar()
	return nil
}

// Sync flushes any buffered log entries
func Sync() {
	if log != nil {
		_ = log.Sync()
	}
}

// Debug logs a debug message
func Debug(args ...interface{}) {
	if log != nil {
		log.Debug(args...)
	}
}

// Debugf logs a formatted debug message
func Debugf(template string, args ...interface{}) {
	if log != nil {
		log.Debugf(template, args...)
	}
}

// Info logs an info message
func Info(args ...interface{}) {
	if log != nil {
		log.Info(args...)
	}
}

// Infof logs a formatted info message
func Infof(template string, args ...interface{}) {
	if log != nil {
		log.Infof(template, args...)
	}
}

// Warn logs a warning message
func Warn(args ...interface{}) {
	if log != nil {
		log.Warn(args...)
	}
}

// Warnf logs a formatted warning message
func Warnf(template string, args ...interface{}) {
	if log != nil {
		log.Warnf(template, args...)
	}
}

// Error logs an error message
func Error(args ...interface{}) {
	if log != nil {
		log.Error(args...)
	}
}

// Errorf logs a formatted error message
func Errorf(template string, args ...interface{}) {
	if log != nil {
		log.Errorf(template, args...)
	}
}

// Fatal logs a fatal message and exits
func Fatal(args ...interface{}) {
	if log != nil {
		log.Fatal(args...)
	} else {
		os.Exit(1)
	}
}

// Fatalf logs a formatted fatal message and exits
func Fatalf(template string, args ...interface{}) {
	if log != nil {
		log.Fatalf(template, args...)
	} else {
		os.Exit(1)
	}
}

// With adds structured context to the logger
func With(args ...interface{}) *zap.SugaredLogger {
	if log != nil {
		return log.With(args...)
	}
	return nil
}
