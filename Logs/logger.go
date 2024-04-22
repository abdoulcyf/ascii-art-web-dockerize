package Logs

import (
	"log/slog"
	"os"
)
	
var (
	opts   = &slog.HandlerOptions{Level: slog.LevelDebug}
	logger = slog.New(slog.NewTextHandler(os.Stderr, opts))

	errMsg string
	logMsg string
)


// LoggerConfig represents the configuration options for the logger
type LoggerConfig struct {
	Level slog.Level // Log level
}

// NewLogger creates a new logger with the provided configuration
func NewLogger(config LoggerConfig) *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: config.Level,
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, opts))
	slog.SetDefault(logger)
	return logger
}
