package http_clients

import (
	"fmt"
	"net/url"
	"encoding/json"
	"github.com/go-resty/resty/v2"

	infraLogger "my-app/internal/infra/loggers"
	infraLoggerInterfaces "my-app/internal/infra/loggers/interfaces"
)

type HTTPClientResty struct {
	Logger infraLogger.GenericLogger
	AdditionalAttributes infraLoggerInterfaces.GenericLoggerAdditionalAttributes
}

func(h HTTPClientResty) Get(requestURL string, bodyParams map[string]string, queryParams url.Values, headerParams map[string]string, timeoutParam int) (map[string]interface{}, int, error) {
	client := resty.New()

	resp, err := client.R().SetBody(bodyParams).
													SetHeader("Content-Type", "application/json").
													SetHeaders(headerParams).
													Get(requestURL+"?"+queryParams.Encode())

	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error during request execution: %v", err), h.AdditionalAttributes)
		return nil, -1, err
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &jsonResponse); err != nil {
		h.Logger.Error(fmt.Sprintf("Error in decode json response: %v", err), h.AdditionalAttributes)
		return nil, -1, err
	}

	jsonResponse["_executed_by"] = "resty_adapter"
	jsonResponse["_status_code_from_req"] = resp.StatusCode()

	return jsonResponse, resp.StatusCode(), nil
}

func Init(logger infraLogger.GenericLogger, additionalAttributes infraLoggerInterfaces.GenericLoggerAdditionalAttributes) HTTPClientResty {
	return HTTPClientResty{
		Logger: logger,
		AdditionalAttributes: additionalAttributes,
	}
}
