package presenters

import (
	"fmt"
	"encoding/json"

	listOutputDTO "my-app/internal/domain/output_dtos/products/list"
	infraLogger "my-app/internal/infra/loggers"
	infraLoggerInterfaces "my-app/internal/infra/loggers/interfaces"
)

type ListProductPresenter struct {
	Logger infraLogger.GenericLogger
	AdditionalAttributes infraLoggerInterfaces.GenericLoggerAdditionalAttributes
}

type ListProductResponse struct {
	SomeField string `json:"some_field"`
	Data listOutputDTO.ListProductOutputDTO `json:"data"`
}

func NewListProductPresenter(logger infraLogger.GenericLogger, additionalAttributes infraLoggerInterfaces.GenericLoggerAdditionalAttributes) *ListProductPresenter {
	return &ListProductPresenter{
		Logger: logger,
		AdditionalAttributes: additionalAttributes,
	}
}

func (p ListProductPresenter) ToJSON(outputDTO listOutputDTO.ListProductOutputDTO) string {
	jsonAsString, err := json.Marshal(ListProductResponse{SomeField: "value-here", Data: outputDTO})

	if err != nil {
		msg := fmt.Sprintf("Internal::Domain::Presenters::Products::List::ListProductPresenter Failure on JSON Marshal | Error %v", err)
		panic(msg)
	}

	return string(jsonAsString)
}
