package main

import (
	"time"
	"fmt"
	"runtime"
	"net/http"
	"os"
	"strconv"

	useCases "my-app/src/domain/use_cases/products/list"
	productRepository "my-app/src/domain/repositories/products"
	createInputDTO "my-app/src/domain/input_dtos/products/create"
	listInputDTO "my-app/src/domain/input_dtos/products/list"

	infraDatabaseDriver "my-app/src/infra/database_drivers"
	infraHTTPRouter "my-app/src/infra/http_routers"
	infraUUIDGenerator "my-app/src/infra/uuid_generators"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))

	if err != nil {
		fmt.Println("Error when convert the env APP_PORT into integer. Error %v", err)
		panic(err)
	}

	fmt.Printf("App listening on port %v | Golang version %v | Now is %s\n", port, runtime.Version(), time.Now().String())

	////////////
	db := infraDatabaseDriver.Init()
	productRepository := productRepository.NewProductRepository(db)

	createProductInputDTO := createInputDTO.CreateProductInputDTO{Identifier: 1, FullName: "John smith", StateName: "ready",}
	productRepository.Insert(createProductInputDTO)

	listProductInputDTO := listInputDTO.ListProductInputDTO{Identifier: 1, FullName: "Headset", StateName: "ready"}
	products, err := useCases.ListProductsUseCase(listProductInputDTO, productRepository)
	if err != nil {
		fmt.Println("Error when invoke ListProductsUseCase. Error ", err)
	}
	fmt.Println("Products from case use", products)
	////////////

	uuidGenerator := infraUUIDGenerator.Init()
	router := infraHTTPRouter.Init(uuidGenerator)

	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
