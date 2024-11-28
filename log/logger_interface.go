package log

// Logger defines a standard logging interface
type Logger interface {
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)
	With(fields ...any) Logger
}

// GetLogger returns the Logger instance
func GetLogger() Logger {
	return getZapLogger()
}
