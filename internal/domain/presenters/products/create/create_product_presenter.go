package presenters

import (
	"fmt"
	"encoding/json"

	createOutputDTO "my-app/internal/domain/output_dtos/products/create"
	infraLogger "my-app/internal/infra/loggers"
	infraLoggerInterfaces "my-app/internal/infra/loggers/interfaces"
)

type CreateProductPresenter struct {
	Logger infraLogger.GenericLogger
	AdditionalAttributes infraLoggerInterfaces.GenericLoggerAdditionalAttributes
}

type CreateProductResponse struct {
	SomeField string `json:"another_field_to_show_on_product_creation"`
	Data createOutputDTO.CreateProductOutputDTO `json:"created_product_data"`
}

func NewCreateProductPresenter(logger infraLogger.GenericLogger, additionalAttributes infraLoggerInterfaces.GenericLoggerAdditionalAttributes) *CreateProductPresenter {
	return &CreateProductPresenter{
		Logger: logger,
		AdditionalAttributes: additionalAttributes,
	}
}

func (p CreateProductPresenter) ToJSON(outputDTO createOutputDTO.CreateProductOutputDTO) string {
	jsonAsString, err := json.Marshal(CreateProductResponse{SomeField: "value-here", Data: outputDTO})

	if err != nil {
		msg := fmt.Sprintf("Internal::Domain::Presenters::Products::Create::CreateProductPresenter Failure on JSON Marshal | Error %v", err)
		panic(msg)
	}

	return string(jsonAsString)
}
