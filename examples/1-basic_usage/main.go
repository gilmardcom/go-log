package main

import (
	l "log"
	"gilmard.com/go-log/log"
)

func main() {
	// Initialize logger for development mode
	err := log.InitLogger("development")
	if err != nil {
		l.Fatalf("Failed to initialize logger: %v", err)
	}

	// Get the logger instance
	logger := log.GetLogger()

	// Log at different levels
	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
}