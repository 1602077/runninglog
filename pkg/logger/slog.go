package logger

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

// New instantiates a new slog logger. This should be called in the main
// entrypoint of your application.
func NewSlogLogger(config *Config) {
	var opts *slog.HandlerOptions
	var handler slog.Handler
	var level slog.Level

	switch config.Level {
	case LevelDebug:
		level = slog.LevelDebug
	case LevelWarn:
		level = slog.LevelWarn
	case LevelError:
		level = slog.LevelError
	case LevelInfo:
		level = slog.LevelInfo
	}

	opts = &slog.HandlerOptions{Level: level}

	if config.StructuredLogging {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	logger = slog.New(handler)
}

func Info(msg string, args ...any) {
	logger.Info(msg, args...)
}

func Debug(msg string, args ...any) {
	logger.Debug(msg, args...)
}

func Error(msg string, args ...any) {
	logger.Error(msg, args...)
}
