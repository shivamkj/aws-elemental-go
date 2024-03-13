package utils

import (
	"log/slog"
	"os"

	"github.com/phsym/console-slog"
	"go.uber.org/zap"
)

type _Logger struct {
	// Debug level is for the development purpose, stripped in production. Can produce large no. of debug logs
	Debug func(msg string, keyValuePairs ...any)
	// Info level is for all the important events, available in production
	Info func(msg string, keyValuePairs ...any)
	// Warn level indicates some issue, but doesn't stops program execution, recoverable issues
	Warn func(msg string, keyValuePairs ...any)
	// Error levels are serious issues, that halts the program, things that panics after logging
	Error func(msg string, keyValuePairs ...any)
}

var Logger *_Logger

func init() {
	if os.Getenv("ENV") == "PROD" {
		Logger = _zapLogger()
	} else {
		Logger = _slogLogger()
	}
}

// Returns Zap Sugared based Logger
func _zapLogger() *_Logger {
	if Logger != nil {
		return Logger
	}

	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	sugaredLogger := zapLogger.Sugar()

	logger := _Logger{
		Debug: sugaredLogger.Debugw,
		Info:  sugaredLogger.Infow,
		Warn:  sugaredLogger.Warnw,
		Error: sugaredLogger.Errorw,
	}

	Logger = &logger
	return &logger
}

// Returns Slog based Logger
func _slogLogger() *_Logger {
	if Logger != nil {
		return Logger
	}

	slogLogger := slog.New(
		console.NewHandler(
			os.Stderr,
			&console.HandlerOptions{Level: slog.LevelDebug, AddSource: true},
		),
	)

	logger := _Logger{
		Debug: slogLogger.Debug,
		Info:  slogLogger.Info,
		Warn:  slogLogger.Warn,
		Error: slogLogger.Error,
	}

	Logger = &logger
	return &logger
}
