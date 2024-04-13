package use_cases

import (
	"fmt"

	inputDtos "my-app/src/domain/input_dtos/products/list"
	outputDtos "my-app/src/domain/output_dtos/products/list"
	repositories "my-app/src/domain/repositories/products"
	entities "my-app/src/domain/entities/products"
)

func ListProductsUseCase(dto inputDtos.ListProductInputDTO, repo repositories.ProductRepositoryInterface) (outputDtos.ListProductOutputDTO, error) {
	products, err := repo.ListBy(dto)

	if err != nil {
		fmt.Println("Error when executing ListProductsUseCase. Error ", err)
		return outputDtos.ListProductOutputDTO{
			Products: []entities.Product{},
		}, err
	}

	return outputDtos.ListProductOutputDTO{
		Products: products,
	}, nil
}
