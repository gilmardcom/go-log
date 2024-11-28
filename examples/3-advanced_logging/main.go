package main

import (
	l "log"
	"github.com/gilmardcom/go-log/log"
)

func main() {
	// Initialize logger for testing mode
	err := log.InitLogger("testing")
	if err != nil {
		l.Fatalf("Failed to initialize logger: %v", err)
	}

	// Get the logger instance
	logger := log.GetLogger()

	// Log structured data
	logger.With("module", "advanced_features", "feature", "structured_logging").
		Info("This log message contains structured data")

	// Chain multiple context additions
	chainedLogger := logger.With("sessionID", "abc123").With("userID", 42)
	chainedLogger.Debug("This debug message has chained context")
	chainedLogger.Error("An error occurred in the chained context")
}