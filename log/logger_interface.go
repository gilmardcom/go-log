package log

// Logger defines a standard logging interface
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	With(fields ...interface{}) Logger
}

// GetLogger returns the Logger instance
func GetLogger() Logger {
	return getZapLogger()
}
