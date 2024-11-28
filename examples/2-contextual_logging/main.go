package main

import ( 
	l "log"
	"github.com/gilmardcom/go-log/log"
)

func main() {
	// Initialize logger for production mode
	err := log.InitLogger("production")
	if err != nil {
		l.Fatalf("Failed to initialize logger: %v", err)
	}

	// Get the logger instance
	logger := log.GetLogger()

	// Add context to the logger
	contextualLogger := logger.With("userID", 12345, "operation", "signup")

	// Log messages with context
	contextualLogger.Info("Starting user operation")
	contextualLogger.Error("User operation failed due to an error")
}