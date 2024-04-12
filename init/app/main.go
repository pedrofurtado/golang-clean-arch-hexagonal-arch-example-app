package main

import (
	"time"
	"fmt"
	"runtime"
	"net/http"
	"os"
	"strconv"

	productRepository "my-app/src/domain/repositories/products"
	productInputDTO "my-app/src/domain/input_dtos/products"

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
	productInputDTO := productInputDTO.ProductInputDTO{Identifier: 1, FullName: "John smith", StateName: "ready",}
	productRepository := productRepository.NewProductRepository(db)
	productRepository.Insert(productInputDTO)
	////////////

	uuidGenerator := infraUUIDGenerator.Init()
	router := infraHTTPRouter.Init(uuidGenerator)

	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
