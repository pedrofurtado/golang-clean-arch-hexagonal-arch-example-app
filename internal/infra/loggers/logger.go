package loggers

import (
	"os"
	"fmt"

	loggerInterfaces "my-app/internal/infra/loggers/interfaces"
	standardLogger "my-app/internal/infra/loggers/standard"
	zapLogger "my-app/internal/infra/loggers/zap"
)

type GenericLogger interface{
	Debug(text string, additionalAttributes loggerInterfaces.GenericLoggerAdditionalAttributes)
	Info(text string, additionalAttributes loggerInterfaces.GenericLoggerAdditionalAttributes)
	Warning(text string, additionalAttributes loggerInterfaces.GenericLoggerAdditionalAttributes)
	Error(text string, additionalAttributes loggerInterfaces.GenericLoggerAdditionalAttributes)
}

func Init(logLevel string) GenericLogger {
	switch os.Getenv("APP_ADAPTER_LOGGER") {
		case "standard":
			return standardLogger.Init(logLevel)
		case "zap":
			return zapLogger.Init(logLevel)
		default:
			err := "Must be defined a adapter for logger"
			fmt.Println(err)
			panic(err)
	}
}
