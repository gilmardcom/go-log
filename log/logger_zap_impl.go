package log

import (
	"log"
	"strings"

	"go.uber.org/zap"
)

var (
	logger Logger // Global Logger instance
)

// ZapLogger wraps zap.SugaredLogger to implement the Logger interface
type ZapLogger struct {
	sugar *zap.SugaredLogger
}

// Implementation of Logger interface methods
func (z *ZapLogger) Debug(args ...interface{}) {
	z.sugar.Debug(args...)
}

func (z *ZapLogger) Info(args ...interface{}) {
	z.sugar.Info(args...)
}

func (z *ZapLogger) Warn(args ...interface{}) {
	z.sugar.Warn(args...)
}

func (z *ZapLogger) Error(args ...interface{}) {
	z.sugar.Error(args...)
}

func (z *ZapLogger) With(fields ...interface{}) Logger {
	return &ZapLogger{sugar: z.sugar.With(fields...)}
}

// InitLogger initializes the logger based on the app mode
func InitLogger(mode string) error {
	var err error
	var zapLogger *zap.Logger

	config := zap.NewProductionConfig() // Base configuration for production
	if strings.ToLower(mode) == "development" {
		config = zap.NewDevelopmentConfig() // Base configuration for development
	}

	// Customize the stack trace level
	config.EncoderConfig.StacktraceKey = "stacktrace" // Only include stack traces for error levels
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	config.DisableStacktrace = true                  // Disable stack traces by default
	config.EncoderConfig.TimeKey = "timestamp"       // Rename the time field to "timestamp"

	zapLogger, err = config.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		log.Println("Failed to initialize logger:", err)
		return err
	}

	logger = &ZapLogger{sugar: zapLogger.Sugar()}
	// Flush logs
	defer func() {
		if syncErr := zapLogger.Sync(); syncErr != nil {
			log.Printf("Error syncing logger: %v", syncErr)
		}
	}()
	logger.Info("Logger successfully initialized")
	return nil
}

// getZapLogger returns the Logger instance
func getZapLogger() Logger {
	return logger
}
