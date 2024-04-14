package http_servers

import (
	"fmt"
	"net/http"
	"context"

	"github.com/go-chi/chi/v5"

	infra_uuid_generator "my-app/internal/infra/uuid_generators"
)

func Init(uuid_generator infra_uuid_generator.GenericUUIDGenerator) http.Handler {
	fmt.Println("Http router | Init | Implementation: Chi")
	fmt.Println("Request ID from chi", uuid_generator.NewUUID())

	r := chi.NewRouter()
	r.Use(generateMiddlewareRequestID(uuid_generator))
	r.Get("/", homeRoute)
	r.Get("/hello", helloRoute)
	r.Get("/show/{my_param}", showMyParamRoute)

	return r
}

func generateMiddlewareRequestID(uuid_generator infra_uuid_generator.GenericUUIDGenerator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "my-app-request-id", uuid_generator.NewUUID())
			next.ServeHTTP(rw, r.WithContext(ctx))
		})
	}
}

func homeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("welcome from chi | Request ID %s", r.Context().Value("my-app-request-id"))))
}

func helloRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("hello world from chi | Request ID %s", r.Context().Value("my-app-request-id"))))
}

func showMyParamRoute(w http.ResponseWriter, r *http.Request) {
	param_value := chi.URLParam(r, "my_param")
	w.Write([]byte(fmt.Sprintf("show my param from chi: %s | Request ID %s", param_value, r.Context().Value("my-app-request-id"))))
}
