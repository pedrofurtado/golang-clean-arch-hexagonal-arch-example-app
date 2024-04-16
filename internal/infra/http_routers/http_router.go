package http_routers

import (
	"net/http"
	"os"
	"fmt"

	infraUUIDGenerator "my-app/internal/infra/uuid_generators"
	infraLogger "my-app/internal/infra/loggers"
	infraLoggerInterfaces "my-app/internal/infra/loggers/interfaces"

	httpRouterChi "my-app/internal/infra/http_routers/chi"
	httpRouterGorillaMux "my-app/internal/infra/http_routers/gorilla_mux"
	httpRouterJulienschmidtHTTPRouter "my-app/internal/infra/http_routers/julienschmidt_httprouter"
)

func Init(uuidGenerator infraUUIDGenerator.GenericUUIDGenerator, logger infraLogger.GenericLogger, loggerAdditionalAttributes infraLoggerInterfaces.GenericLoggerAdditionalAttributes) http.Handler {
	adapter := os.Getenv("APP_ADAPTER_HTTP_ROUTER")

	logger.Info(fmt.Sprintf("Internal::Infra::HttpRouters::Init Http router configured | Adapter %v", adapter), loggerAdditionalAttributes)

	switch adapter {
		case "chi":
			return httpRouterChi.Init(uuidGenerator, logger, loggerAdditionalAttributes)
		case "gorilla_mux":
			return httpRouterGorillaMux.Init(uuidGenerator, logger, loggerAdditionalAttributes)
		case "julienschmidt_httprouter":
			return httpRouterJulienschmidtHTTPRouter.Init(uuidGenerator, logger, loggerAdditionalAttributes)
		default:
			msg := fmt.Sprintf("Internal::Infra::HttpRouters::Init Invalid adapter for http router | Adapter %v", adapter)
			logger.Error(msg, loggerAdditionalAttributes)
			panic(msg)
	}
}
