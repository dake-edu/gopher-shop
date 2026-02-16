package logger

import (
	"log/slog"
	"os"
)

// Setup initializes the global logger with JSON handler.
// Pedagogical Anchor: "The Black Box Recorder".
func Setup() {
	// JSONHandler outputs logs as structured JSON objects.
	// Source: true adds source file/line info (great for debugging).
	opts := &slog.HandlerOptions{
		Level:     slog.LevelDebug, // Log everything for learning purposes
		AddSource: true,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)
}
