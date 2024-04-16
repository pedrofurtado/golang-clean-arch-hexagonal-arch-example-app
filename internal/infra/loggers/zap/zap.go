package loggers

import (
	"fmt"

	_ "github.com/jsternberg/zap-logfmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	loggerInterfaces "my-app/internal/infra/loggers/interfaces"
)

type ZapLogger struct {
	logger *zap.Logger
}

func (s ZapLogger) Debug(text string, additionalAttributes loggerInterfaces.GenericLoggerAdditionalAttributes) {
	defer s.logger.Sync()
	s.logger.Debug(text, zap.String("transactionId", additionalAttributes.TransactionId), zap.String("traceId", additionalAttributes.TraceId))
}

func (s ZapLogger) Info(text string, additionalAttributes loggerInterfaces.GenericLoggerAdditionalAttributes) {
	defer s.logger.Sync()
	s.logger.Info(text, zap.String("transactionId", additionalAttributes.TransactionId), zap.String("traceId", additionalAttributes.TraceId))
}

func (s ZapLogger) Warning(text string, additionalAttributes loggerInterfaces.GenericLoggerAdditionalAttributes) {
	defer s.logger.Sync()
	s.logger.Warn(text, zap.String("transactionId", additionalAttributes.TransactionId), zap.String("traceId", additionalAttributes.TraceId))
}

func (s ZapLogger) Error(text string, additionalAttributes loggerInterfaces.GenericLoggerAdditionalAttributes) {
	defer s.logger.Sync()
	s.logger.Error(text, zap.String("transactionId", additionalAttributes.TransactionId), zap.String("traceId", additionalAttributes.TraceId))
}

func Init(logLevel string) ZapLogger {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zapLogLevel(logLevel)),
		Encoding:    "logfmt",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	zapLogger, _ := logConfig.Build()

	return ZapLogger{
		logger: zapLogger,
	}
}

func zapLogLevel(logLevel string) zapcore.Level {
	switch logLevel {
		case "debug":
			return zapcore.DebugLevel
		case "info":
			return zapcore.InfoLevel
		case "warn":
			return zapcore.WarnLevel
		case "error":
			return zapcore.ErrorLevel
		default:
			msg := "Must be defined a valid logger level for zap"
			fmt.Println(msg, logLevel)
			panic(msg)
	}
}
