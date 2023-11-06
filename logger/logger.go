package logger

import (
	"log/slog"
	"os"
)

// Set the global slog logger to log all messages at or above the debug level.
func SetLoggerLevelToDebug() {
	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(textHandler)
	slog.SetDefault(logger)
}
