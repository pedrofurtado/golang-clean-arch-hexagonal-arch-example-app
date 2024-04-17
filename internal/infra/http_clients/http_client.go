package http_clients

import (
	"os"
	"net/url"
	"fmt"

	infraLogger "my-app/internal/infra/loggers"
	infraLoggerInterfaces "my-app/internal/infra/loggers/interfaces"
	standardHttpClient "my-app/internal/infra/http_clients/standard"
	restyHttpClient "my-app/internal/infra/http_clients/resty"
)

type GenericHTTPClient interface {
	Get(requestURL string, bodyParams map[string]string, queryParams url.Values, headerParams map[string]string, timeoutParam int) (map[string]interface{}, int, error)
}

func Init(logger infraLogger.GenericLogger, loggerAdditionalAttributes infraLoggerInterfaces.GenericLoggerAdditionalAttributes) GenericHTTPClient {
	adapter := os.Getenv("APP_ADAPTER_HTTP_CLIENT")

	logger.Info(fmt.Sprintf("Internal::Infra::HttpClients::Init HTTP Client configured | Adapter %v", adapter), loggerAdditionalAttributes)

	switch adapter {
		case "standard":
			return standardHttpClient.Init(logger, loggerAdditionalAttributes)
		case "resty":
			return restyHttpClient.Init(logger, loggerAdditionalAttributes)
		default:
			msg := fmt.Sprintf("Internal::Infra::HttpClients::Init Invalid adapter for HTTP client | Adapter %v", adapter)
			logger.Error(msg, loggerAdditionalAttributes)
			panic(msg)
	}
}
