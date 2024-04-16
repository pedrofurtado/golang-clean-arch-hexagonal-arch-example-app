package main

import (
	"time"
	"fmt"
	"runtime"
	"net/http"
	"os"
	"strconv"

	infraLogger "my-app/internal/infra/loggers"
	infraLoggerInterfaces "my-app/internal/infra/loggers/interfaces"
	infraHTTPRouter "my-app/internal/infra/http_routers"
	infraUUIDGenerator "my-app/internal/infra/uuid_generators"
)

func main() {
	uuidGenerator := infraUUIDGenerator.Init()

	loggerAdditionalAttributes := infraLoggerInterfaces.GenericLoggerAdditionalAttributes{TransactionId: uuidGenerator.NewUUID(), TraceId: uuidGenerator.NewUUID()}
	loggerAdapter := os.Getenv("APP_ADAPTER_LOGGER")
	loggerLevel := os.Getenv("APP_ADAPTER_LOGGER_LEVEL")
	logger := infraLogger.Init(loggerLevel)
	logger.Info(fmt.Sprintf("Cmd::Web::Main Logger configured | Adapter %v | Level %v", loggerAdapter, loggerLevel), loggerAdditionalAttributes)

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))

	if err != nil {
		msg := fmt.Sprintf("Cmd::Web::Main APP_PORT invalid | Port %v | Error %v", port, err)
		logger.Error(msg, loggerAdditionalAttributes)
		panic(msg)
	}

	logger.Info(fmt.Sprintf("Cmd::Web::Main App Started | Listening on port %v | Golang version %v | Current time is %s\n", port, runtime.Version(), time.Now().String()), loggerAdditionalAttributes)

	router := infraHTTPRouter.Init(uuidGenerator, logger, loggerAdditionalAttributes)

	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
