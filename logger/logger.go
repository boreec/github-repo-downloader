package logger

import (
	"log/slog"
	"os"
)

func SetLoggerLevelToDebug() {
	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(textHandler)
	slog.SetDefault(logger)
}
