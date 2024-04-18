package facades

import (
	"fmt"

	listOutputDTO "my-app/internal/domain/output_dtos/products/list"
	createOutputDTO "my-app/internal/domain/output_dtos/products/create"
	listInputDTO "my-app/internal/domain/input_dtos/products/list"
	createInputDTO "my-app/internal/domain/input_dtos/products/create"
	listUseCases "my-app/internal/domain/use_cases/products/list"
	createUseCases "my-app/internal/domain/use_cases/products/create"
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

func (f ProductFacade) CreateProduct(createProductInputDTO createInputDTO.CreateProductInputDTO) (createOutputDTO.CreateProductOutputDTO, error) {
	createProductOutputDTO, err := createUseCases.CreateProductUseCase(createProductInputDTO, f.ProductRepository)

	if err != nil {
		f.Logger.Error(fmt.Sprintf("ProductFacade::CreateProduct Invoked CreateProductUseCase with failure | Error %v", err), f.AdditionalAttributes)
		return createOutputDTO.CreateProductOutputDTO{}, err
	}

	f.Logger.Info("ProductFacade::CreateProduct Invoked CreateProductUseCase with success", f.AdditionalAttributes)

	return createProductOutputDTO, nil
}
