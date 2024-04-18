package use_cases

import (
	"fmt"

	createInputDtos "my-app/internal/domain/input_dtos/products/create"
	createOutputDtos "my-app/internal/domain/output_dtos/products/create"
	repositories "my-app/internal/domain/repositories/products"
	entities "my-app/internal/domain/entities/products"
)

func CreateProductUseCase(dto createInputDtos.CreateProductInputDTO, repo repositories.ProductRepositoryInterface) (createOutputDtos.CreateProductOutputDTO, error) {
	products, err := repo.Insert(dto)

	if err != nil {
		fmt.Println("Error when executing CreateProductUseCase. Error ", err)
		return createOutputDtos.CreateProductOutputDTO{
			Products: []entities.Product{},
		}, err
	}

	return createOutputDtos.CreateProductOutputDTO{
		Products: products,
	}, nil
}
