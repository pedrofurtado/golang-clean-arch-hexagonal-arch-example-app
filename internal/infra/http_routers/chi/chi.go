package http_servers

import (
	"os"
	"fmt"
	"net/http"
	"context"

	"github.com/go-chi/chi/v5"

	facade "my-app/internal/domain/facades/products"
	infraUUIDGenerator "my-app/internal/infra/uuid_generators"
	infraLogger "my-app/internal/infra/loggers"
	infraLoggerInterfaces "my-app/internal/infra/loggers/interfaces"
	listInputDTO "my-app/internal/domain/input_dtos/products/list"
	repository "my-app/internal/domain/repositories/products"
	infraDatabaseDriver "my-app/internal/infra/database_drivers"
)

func Init(uuidGenerator infraUUIDGenerator.GenericUUIDGenerator, logger infraLogger.GenericLogger, loggerAdditionalAttributes infraLoggerInterfaces.GenericLoggerAdditionalAttributes) http.Handler {
	r := chi.NewRouter()
	r.Use(generateMiddlewareRequestID(uuidGenerator))
	r.Get("/", homeRoute)
	r.Get("/products", listProductsRoute)
	r.Get("/products/{productId}", findProductRoute)

	return r
}

func generateMiddlewareRequestID(uuidGenerator infraUUIDGenerator.GenericUUIDGenerator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "my-app-request-id", uuidGenerator.NewUUID())
			next.ServeHTTP(rw, r.WithContext(ctx))
		})
	}
}

func homeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("welcome from chi | Request ID %s", r.Context().Value("my-app-request-id"))))
}

func listProductsRoute(w http.ResponseWriter, r *http.Request) {
	db := infraDatabaseDriver.Init()
	productRepository := repository.NewProductRepository(db)
	logger := infraLogger.Init(os.Getenv("APP_ADAPTER_LOGGER_LEVEL"))
	additionalAttributes := infraLoggerInterfaces.GenericLoggerAdditionalAttributes{TransactionId: r.Context().Value("my-app-request-id").(string), TraceId: ""}

	productFacade := facade.NewProductFacade(productRepository, logger, additionalAttributes)

	inputDTO := listInputDTO.ListProductInputDTO{}
	productsOutputDTO, _ := productFacade.ListProducts(inputDTO)

	w.Write([]byte(fmt.Sprintf("hello world from chi | Request ID %s | Products Output DTO %v", r.Context().Value("my-app-request-id"), productsOutputDTO)))
}

func findProductRoute(w http.ResponseWriter, r *http.Request) {
	param_value := chi.URLParam(r, "productId")
	w.Write([]byte(fmt.Sprintf("show my param from chi: %s | Request ID %s", param_value, r.Context().Value("my-app-request-id"))))
}
