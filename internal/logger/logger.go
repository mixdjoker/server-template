package logger

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

// Log level constants define the standard logging levels used throughout the application.
const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warning"
)

var logger = New()

// SetupLoggerLevel configures the logger with the specified log level and output format.
func SetupLoggerLevel(level string, format string) {
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: format,
	}

	logger.Logger = logger.Output(consoleWriter)

	switch level {
	case DebugLevel:
		logger.Logger = logger.Level(zerolog.DebugLevel)
		logger.Info().Msg("Log Level is switched to DEBUG level")
	case InfoLevel:
		logger.Logger = logger.Level(zerolog.InfoLevel)
		logger.Info().Msg("Log Level is switched to INFO level")
	case WarnLevel:
		logger.Logger = logger.Level(zerolog.WarnLevel)
		logger.Info().Msg("Log Level is switched to WARNING level")
	default:
		logger.Logger = logger.Level(zerolog.InfoLevel)
		logger.Info().Msg("Log Level is switched to INFO level")
	}
}

// Info starts new message in info level with ctx
func Info(_ context.Context) *zerolog.Event {
	ev := logger.Info()
	// discoverContext(ctx, ev)
	return ev
}

// Debug creates a new debug event with the given context.
func Debug(_ context.Context) *zerolog.Event {
	ev := logger.Debug()
	// discoverContext(ctx, ev)
	return ev
}

// Warn creates a new warning event with the given context.
func Warn(_ context.Context) *zerolog.Event {
	ev := logger.Warn()
	// discoverContext(ctx, ev)
	return ev
}

// Error creates a new error event with the given context.
func Error(_ context.Context) *zerolog.Event {
	ev := logger.Error()
	// discoverContext(ctx, ev)
	return ev
}

// Err creates a new error event with a specific error and the given context.
func Err(_ context.Context, err error) *zerolog.Event {
	ev := logger.Err(err)
	// discoverContext(ctx, ev)
	return ev
}

// Fatalf logging error and TERMINATE application. Use it VERY careful!!!
func Fatalf(_ context.Context, err error, msg string, v ...any) {
	ev := logger.Err(err)
	// discoverContext(ctx, ev)

	ev.Msgf(msg, v...)
	os.Exit(1)
}

// func discoverContext(ctx context.Context, ev *zerolog.Event) {
// if ip, ok := ctx.Value(TxIP).(string); ok {
// 	ev = ev.Str(ipLogPref, ip)
// }
// }

// GetLogger returns a pointer to the package-level logger instance.
func GetLogger() *CustomLogger {
	return &logger
}

// CustomLogger is a wrapper around zerolog.Logger that allows for
// extending the functionality of the base logger.
type CustomLogger struct {
	zerolog.Logger
}

// New creates and returns a new instance of CustomLogger.
// This function initializes a zerolog.Logger that writes to standard output (os.Stdout)
// with timestamped log entries and returns it wrapped in a CustomLogger.
func New() CustomLogger {
	l := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return CustomLogger{Logger: l}
}
