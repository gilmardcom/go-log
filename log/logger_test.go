package log

import (
	"bytes"
	"strings"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// MockWriter to capture log output
type MockWriter struct {
	buf bytes.Buffer
}

func (w *MockWriter) Write(p []byte) (n int, err error) {
	return w.buf.Write(p)
}

func (w *MockWriter) String() string {
	return w.buf.String()
}

func (w *MockWriter) Reset() {
	w.buf.Reset()
}

// Helper function to create a test logger
func setupTestLogger() (*MockWriter, Logger) {
	mockWriter := &MockWriter{}

	// Configure a zap logger to write to MockWriter
	encoderCfg := zapcore.EncoderConfig{
		MessageKey: "message",
		LevelKey:   "level",
		EncodeLevel: func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(strings.ToUpper(l.String()))
		},
		TimeKey:    "",
		CallerKey:  "",
		EncodeTime: nil,
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(mockWriter),
		zap.DebugLevel,
	)
	zapLogger := zap.New(core).Sugar()

	// Wrap the zap logger
	return mockWriter, &ZapLogger{sugar: zapLogger}
}

func TestLoggerDebug(t *testing.T) {
	writer, logger := setupTestLogger()
	logger.Debug("debug message")

	output := writer.String()
	if !strings.Contains(output, "debug message") {
		t.Errorf("Expected 'debug message' in log output, got: %s", output)
	}
}

func TestLoggerInfo(t *testing.T) {
	writer, logger := setupTestLogger()
	logger.Info("info message")

	output := writer.String()
	if !strings.Contains(output, "info message") {
		t.Errorf("Expected 'info message' in log output, got: %s", output)
	}
}

func TestLoggerWarn(t *testing.T) {
	writer, logger := setupTestLogger()
	logger.Warn("warn message")

	output := writer.String()
	if !strings.Contains(output, "warn message") {
		t.Errorf("Expected 'warn message' in log output, got: %s", output)
	}
}

func TestLoggerError(t *testing.T) {
	writer, logger := setupTestLogger()
	logger.Error("error message")

	output := writer.String()
	if !strings.Contains(output, "error message") {
		t.Errorf("Expected 'error message' in log output, got: %s", output)
	}
}

func TestLoggerFatal(t *testing.T) {
	writer, logger := setupTestLogger()
	logger.Fatal("fatal error message")

	output := writer.String()
	if !strings.Contains(output, "fatal error message") {
		t.Errorf("Expected 'fatal error message' in log output, got: %s", output)
	}
}

func TestLoggerWith(t *testing.T) {
	writer, logger := setupTestLogger()
	contextualLogger := logger.With("key", "value")
	contextualLogger.Info("contextual message")

	output := writer.String()
	if !strings.Contains(output, "contextual message") || !strings.Contains(output, `"key":"value"`) {
		t.Errorf("Expected 'contextual message' with key-value pair in log output, got: %s", output)
	}
}
