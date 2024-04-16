package http_servers

import (
	"os"
	"fmt"
	"net/http"
	"context"

	"github.com/julienschmidt/httprouter"

	facade "my-app/internal/domain/facades/products"
	infraUUIDGenerator "my-app/internal/infra/uuid_generators"
	infraLogger "my-app/internal/infra/loggers"
	infraLoggerInterfaces "my-app/internal/infra/loggers/interfaces"
	listInputDTO "my-app/internal/domain/input_dtos/products/list"
	repository "my-app/internal/domain/repositories/products"
	infraDatabaseDriver "my-app/internal/infra/database_drivers"
)

func Init(uuidGenerator infraUUIDGenerator.GenericUUIDGenerator, logger infraLogger.GenericLogger, loggerAdditionalAttributes infraLoggerInterfaces.GenericLoggerAdditionalAttributes) http.Handler {
	r := httprouter.New()
	r.GET("/", generateMiddlewareRequestID(uuidGenerator)(homeRoute))
	r.GET("/products", generateMiddlewareRequestID(uuidGenerator)(listProductsRoute))
	r.GET("/products/:productId", generateMiddlewareRequestID(uuidGenerator)(findProductRoute))

	return r
}

func generateMiddlewareRequestID(uuidGenerator infraUUIDGenerator.GenericUUIDGenerator) func(httprouter.Handle) httprouter.Handle {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			ctx := context.WithValue(r.Context(), "my-app-request-id", uuidGenerator.NewUUID())
			next(rw, r.WithContext(ctx), ps)
		}
	}
}

func homeRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("welcome from Julienschmidt Httprouter | Request ID %s", r.Context().Value("my-app-request-id"))))
}

func listProductsRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := infraDatabaseDriver.Init()
	productRepository := repository.NewProductRepository(db)
	logger := infraLogger.Init(os.Getenv("APP_ADAPTER_LOGGER_LEVEL"))
	additionalAttributes := infraLoggerInterfaces.GenericLoggerAdditionalAttributes{TransactionId: r.Context().Value("my-app-request-id").(string), TraceId: ""}

	productFacade := facade.NewProductFacade(productRepository, logger, additionalAttributes)

	inputDTO := listInputDTO.ListProductInputDTO{}
	productsOutputDTO, _ := productFacade.ListProducts(inputDTO)

	w.Write([]byte(fmt.Sprintf("hello world from Julienschmidt Httprouter | Request ID %s | Products Output DTO %v", r.Context().Value("my-app-request-id"), productsOutputDTO)))
}

func findProductRoute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param_value := ps.ByName("productId")
	w.Write([]byte(fmt.Sprintf("show my param from Julienschmidt Httprouter: %s | Request ID %s", param_value, r.Context().Value("my-app-request-id"))))
}
