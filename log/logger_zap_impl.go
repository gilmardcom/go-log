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

// Implementation of Logger interface
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

	switch strings.ToLower(mode) {
	default:
	case "development", "testing":
		zapLogger, err = zap.NewDevelopment()
		if err != nil {
			log.Println("cannot create 'development' logger")
			return err
		}
	case "production":
		zapLogger, err = zap.NewProduction()
		if err != nil {
			log.Println("cannot create 'production' logger")
			return err
		}
	}

	// Initialize the logger instance with the wrapped ZapLogger
	logger = &ZapLogger{sugar: zapLogger.Sugar()}

	defer zapLogger.Sync() // Flushes buffer, if any

	logger.Info("Logger successfully initialized")

	return nil
}

// getZapLogger returns the Logger instance
func getZapLogger() Logger {
	return logger
}
