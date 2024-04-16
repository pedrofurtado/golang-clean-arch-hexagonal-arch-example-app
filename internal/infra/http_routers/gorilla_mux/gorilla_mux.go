package http_servers

import (
	"os"
	"fmt"
	"net/http"
	"context"

	"github.com/gorilla/mux"

	facade "my-app/internal/domain/facades/products"
	infraUUIDGenerator "my-app/internal/infra/uuid_generators"
	infraLogger "my-app/internal/infra/loggers"
	infraLoggerInterfaces "my-app/internal/infra/loggers/interfaces"
	listInputDTO "my-app/internal/domain/input_dtos/products/list"
	repository "my-app/internal/domain/repositories/products"
	infraDatabaseDriver "my-app/internal/infra/database_drivers"
	listPresenters "my-app/internal/domain/presenters/products/list"
)

func Init(uuidGenerator infraUUIDGenerator.GenericUUIDGenerator, logger infraLogger.GenericLogger, loggerAdditionalAttributes infraLoggerInterfaces.GenericLoggerAdditionalAttributes) http.Handler {
	r := mux.NewRouter()
	r.Use(generateMiddlewareRequestID(uuidGenerator))
	r.Use(contentTypeApplicationJsonMiddleware)
	r.HandleFunc("/", homeRoute).Methods("GET")
	r.HandleFunc("/products", listProductsRoute).Methods("GET")
	r.HandleFunc("/products/{productId}", findProductRoute).Methods("GET")

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

func contentTypeApplicationJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func homeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("welcome from gorilla mux | Request ID %s", r.Context().Value("my-app-request-id"))))
}

func listProductsRoute(w http.ResponseWriter, r *http.Request) {
	db := infraDatabaseDriver.Init()
	productRepository := repository.NewProductRepository(db)
	logger := infraLogger.Init(os.Getenv("APP_ADAPTER_LOGGER_LEVEL"))
	additionalAttributes := infraLoggerInterfaces.GenericLoggerAdditionalAttributes{TransactionId: r.Context().Value("my-app-request-id").(string), TraceId: ""}

	productFacade := facade.NewProductFacade(productRepository, logger, additionalAttributes)

	inputDTO := listInputDTO.ListProductInputDTO{}
	productsOutputDTO, _ := productFacade.ListProducts(inputDTO)

	presenter := listPresenters.NewListProductPresenter(logger, additionalAttributes)

	w.Write([]byte(presenter.ToJSON(productsOutputDTO)))
}

func findProductRoute(w http.ResponseWriter, r *http.Request) {
	param_value := mux.Vars(r)["productId"]
	w.Write([]byte(fmt.Sprintf("show my param from gorilla mux: %s | Request ID %s", param_value, r.Context().Value("my-app-request-id"))))
}
