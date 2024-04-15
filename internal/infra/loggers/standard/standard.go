package loggers

import (
	"fmt"
	"os"
	"log/slog"

	loggerInterfaces "my-app/internal/infra/loggers/interfaces"
)

type StandardLogger struct{
	logger *slog.Logger
}

func (s StandardLogger) Debug(text string, additionalAttributes loggerInterfaces.GenericLoggerAdditionalAttributes) {
	s.logger.Debug(text)
}

func (s StandardLogger) Info(text string, additionalAttributes loggerInterfaces.GenericLoggerAdditionalAttributes) {
	s.logger.Info(text)
}

func (s StandardLogger) Warning(text string, additionalAttributes loggerInterfaces.GenericLoggerAdditionalAttributes) {
	s.logger.Warn(text)
}

func (s StandardLogger) Error(text string, additionalAttributes loggerInterfaces.GenericLoggerAdditionalAttributes) {
	s.logger.Error(text)
}

func Init(logLevel string) StandardLogger {
	slogLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:       standardLogLevel(logLevel),
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.MessageKey {
				a.Key = "message"
			}

			return a
		},
	}))

	return StandardLogger{
		logger: slogLogger,
	}
}

func standardLogLevel(logLevel string) slog.Level {
	switch logLevel {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		msg := "Must be defined a valid logger level for standard"
		fmt.Println(msg, logLevel)
		panic(msg)
	}
}
