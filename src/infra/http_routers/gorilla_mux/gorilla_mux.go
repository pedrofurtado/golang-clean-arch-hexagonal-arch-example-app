package http_servers

import (
	"fmt"
	"net/http"
	"context"

	"github.com/gorilla/mux"

	infra_uuid_generator "my-app/src/infra/uuid_generators"
)

func Init(uuid_generator infra_uuid_generator.GenericUUIDGenerator) http.Handler {
	fmt.Println("Http router | Init | Implementation: Gorilla Mux")
	fmt.Println("Request ID from Gorilla Mux", uuid_generator.NewUUID())

	r := mux.NewRouter()
	r.Use(generateMiddlewareRequestID(uuid_generator))
	r.HandleFunc("/", homeRoute)
	r.HandleFunc("/hello", helloRoute)
	r.HandleFunc("/show/{my_param}", showMyParamRoute)

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
	w.Write([]byte(fmt.Sprintf("welcome from gorilla mux | Request ID %s", r.Context().Value("my-app-request-id"))))
}

func helloRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("hello world from gorilla mux | Request ID %s", r.Context().Value("my-app-request-id"))))
}

func showMyParamRoute(w http.ResponseWriter, r *http.Request) {
	param_value := mux.Vars(r)["my_param"]
	w.Write([]byte(fmt.Sprintf("show my param from gorilla mux: %s | Request ID %s", param_value, r.Context().Value("my-app-request-id"))))
}
