package facades

import (
	"fmt"

	listOutputDTO "my-app/internal/domain/output_dtos/products/list"
	listInputDTO "my-app/internal/domain/input_dtos/products/list"
	listUseCases "my-app/internal/domain/use_cases/products/list"
	repository "my-app/internal/domain/repositories/products"
	infraLogger "my-app/internal/infra/loggers"
	infraLoggerInterfaces "my-app/internal/infra/loggers/interfaces"
)

type ProductFacade struct {
	ProductRepository repository.ProductRepositoryInterface
	Logger infraLogger.GenericLogger
	AdditionalAttributes infraLoggerInterfaces.GenericLoggerAdditionalAttributes
}

func NewProductFacade(productRepository repository.ProductRepositoryInterface, logger infraLogger.GenericLogger, additionalAttributes infraLoggerInterfaces.GenericLoggerAdditionalAttributes) *ProductFacade {
	return &ProductFacade{
		ProductRepository: productRepository,
		Logger: logger,
		AdditionalAttributes: additionalAttributes,
	}
}

func (f ProductFacade) ListProducts(listProductInputDTO listInputDTO.ListProductInputDTO) (listOutputDTO.ListProductOutputDTO, error) {
	listProductsOutputDTO, err := listUseCases.ListProductsUseCase(listProductInputDTO, f.ProductRepository)

	if err != nil {
		f.Logger.Error(fmt.Sprintf("ProductFacade::ListProducts Invoked ListProductsUseCase with failure | Error %v", err), f.AdditionalAttributes)
		return listOutputDTO.ListProductOutputDTO{}, err
	}

	f.Logger.Info("ProductFacade::ListProducts Invoked ListProductsUseCase with success", f.AdditionalAttributes)

	return listProductsOutputDTO, nil
}
