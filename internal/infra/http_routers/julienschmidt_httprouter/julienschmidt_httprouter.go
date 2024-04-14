package http_servers

import (
	"fmt"
	"net/http"
	"context"

	"github.com/julienschmidt/httprouter"

	infra_uuid_generator "my-app/internal/infra/uuid_generators"
)

func Init(uuid_generator infra_uuid_generator.GenericUUIDGenerator) http.Handler {
	fmt.Println("Http router | Init | Implementation: Julienschmidt Httprouter")
	fmt.Println("Request ID from Julienschmidt Httprouter", uuid_generator.NewUUID())

	r := httprouter.New()
	r.GET("/", generateMiddlewareRequestID(uuid_generator)(homeRoute))
	r.GET("/hello", generateMiddlewareRequestID(uuid_generator)(helloRoute))
	r.GET("/show/:my_param", generateMiddlewareRequestID(uuid_generator)(showMyParamRoute))

	return r
}

func generateMiddlewareRequestID(uuid_generator infra_uuid_generator.GenericUUIDGenerator) func(httprouter.Handle) httprouter.Handle {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			ctx := context.WithValue(r.Context(), "my-app-request-id", uuid_generator.NewUUID())
			next(rw, r.WithContext(ctx), ps)
		}
	}
}

func homeRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("welcome from Julienschmidt Httprouter | Request ID %s", r.Context().Value("my-app-request-id"))))
}

func helloRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("hello world from Julienschmidt Httprouter | Request ID %s", r.Context().Value("my-app-request-id"))))
}

func showMyParamRoute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param_value := ps.ByName("my_param")
	w.Write([]byte(fmt.Sprintf("show my param from Julienschmidt Httprouter: %s | Request ID %s", param_value, r.Context().Value("my-app-request-id"))))
}
