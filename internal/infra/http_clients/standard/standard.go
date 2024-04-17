package http_clients

import (
	"fmt"
	"net/http"
	"net/url"
	"encoding/json"
	"bytes"
	"time"

	infraLogger "my-app/internal/infra/loggers"
	infraLoggerInterfaces "my-app/internal/infra/loggers/interfaces"
)

type HTTPClientStandard struct {
	Logger infraLogger.GenericLogger
	AdditionalAttributes infraLoggerInterfaces.GenericLoggerAdditionalAttributes
}

func(h HTTPClientStandard) Get(requestURL string, bodyParams map[string]string, queryParams url.Values, headerParams map[string]string, timeoutParam int) (map[string]interface{}, int, error) {
	// Prepare the request body params
	body, err := json.Marshal(bodyParams)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error on serialize req body params: %v", err), h.AdditionalAttributes)
		return nil, -1, err
	}

	// Create HTTP request
	req, err := http.NewRequest("GET", requestURL+"?"+queryParams.Encode(), bytes.NewBuffer(body))
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error on creation of HTTP request: %v", err), h.AdditionalAttributes)
		return nil, -1, err
	}

	// Add headers
	req.Header.Set("Content-Type", "application/json")
	for headerName, headerValue := range headerParams {
		req.Header.Set(headerName, headerValue)
	}

	// Do the request
	client := &http.Client{
		Timeout: time.Duration(timeoutParam) * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error during http request execution: %v", err), h.AdditionalAttributes)
		return nil, -1, err
	}
	defer resp.Body.Close()

	// Read body from response
	var responseJSON map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseJSON)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error when reading response body:", err), h.AdditionalAttributes)
		return nil, -1, err
	}

	responseJSON["_executed_by"] = "standard_adapter"
	responseJSON["_status_code_from_req"] = resp.StatusCode

	return responseJSON, resp.StatusCode, nil
}

func Init(logger infraLogger.GenericLogger, additionalAttributes infraLoggerInterfaces.GenericLoggerAdditionalAttributes) HTTPClientStandard {
	return HTTPClientStandard{
		Logger: logger,
		AdditionalAttributes: additionalAttributes,
	}
}
