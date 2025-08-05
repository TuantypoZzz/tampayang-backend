package mylogger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"tampayang-backend/config"

	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

func init() {
	logger = logrus.New()

	// Create logs directory if it doesn't exist
	logDir := logPath()
	if err := os.MkdirAll(logDir, 0755); err != nil {
		fmt.Printf("Failed to create log directory: %v\n", err)
	}

	// Set up log file with rotation
	logFile := filepath.Join(logDir, fmt.Sprintf("application-%s.log", time.Now().Format("2006-01-02")))
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open log file: %v\n", err)
		logger.SetOutput(os.Stdout)
	} else {
		logger.SetOutput(file)
	}

	// Set formatter
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
	})

	// Set log level based on environment
	if config.GO_ENV == "development" {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
}

// Info logs informational messages
func Info(action string, data interface{}) {
	entry := logger.WithFields(logrus.Fields{
		"action":    action,
		"data":      data,
		"caller":    getCaller(),
		"timestamp": time.Now().UTC(),
	})
	entry.Info("Application Info")
}

// Error logs error messages
func Error(action string, data interface{}) {
	entry := logger.WithFields(logrus.Fields{
		"action":    action,
		"data":      data,
		"caller":    getCaller(),
		"timestamp": time.Now().UTC(),
	})
	entry.Error("Application Error")
}

// Warn logs warning messages
func Warn(action string, data interface{}) {
	entry := logger.WithFields(logrus.Fields{
		"action":    action,
		"data":      data,
		"caller":    getCaller(),
		"timestamp": time.Now().UTC(),
	})
	entry.Warn("Application Warning")
}

// Debug logs debug messages (only in development)
func Debug(action string, data interface{}) {
	if config.GO_ENV == "development" {
		entry := logger.WithFields(logrus.Fields{
			"action":    action,
			"data":      data,
			"caller":    getCaller(),
			"timestamp": time.Now().UTC(),
		})
		entry.Debug("Application Debug")
	}
}

// Trace logs trace messages
func Trace(action string, data interface{}) {
	entry := logger.WithFields(logrus.Fields{
		"action":    action,
		"data":      data,
		"caller":    getCaller(),
		"timestamp": time.Now().UTC(),
	})
	entry.Trace("Application Trace")
}

// Security logs security-related events
func Security(action string, data interface{}) {
	// Create separate security log file
	securityLogFile := filepath.Join(logPath(), fmt.Sprintf("security-%s.log", time.Now().Format("2006-01-02")))
	file, err := os.OpenFile(securityLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		Error("security_log_error", map[string]interface{}{
			"error": err.Error(),
			"file":  securityLogFile,
		})
		return
	}
	defer file.Close()

	securityLogger := logrus.New()
	securityLogger.SetOutput(file)
	securityLogger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	entry := securityLogger.WithFields(logrus.Fields{
		"action":    action,
		"data":      data,
		"caller":    getCaller(),
		"timestamp": time.Now().UTC(),
		"severity":  "SECURITY",
	})
	entry.Warn("Security Event")
}

// Performance logs performance metrics
func Performance(action string, duration time.Duration, data interface{}) {
	entry := logger.WithFields(logrus.Fields{
		"action":      action,
		"duration_ms": duration.Milliseconds(),
		"data":        data,
		"caller":      getCaller(),
		"timestamp":   time.Now().UTC(),
		"metric_type": "performance",
	})
	entry.Info("Performance Metric")
}

// getCaller returns the caller function name and line number
func getCaller() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}
	return fmt.Sprintf("%s:%d", filepath.Base(file), line)
}

func logPath() string {
	rootPath := config.ProjectRootPath
	return filepath.Join(rootPath, "logs")
}
