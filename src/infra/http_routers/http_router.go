package http_routers

import (
	"net/http"
	"os"
	"fmt"

	infraUUIDGenerator "my-app/src/infra/uuid_generators"

	httpRouterChi "my-app/src/infra/http_routers/chi"
	httpRouterGorillaMux "my-app/src/infra/http_routers/gorilla_mux"
	httpRouterJulienschmidtHTTPRouter "my-app/src/infra/http_routers/julienschmidt_httprouter"
)

func Init(uuid_generator infraUUIDGenerator.GenericUUIDGenerator) http.Handler {
	switch os.Getenv("APP_ADAPTER_HTTP_ROUTER") {
		case "chi":
			return httpRouterChi.Init(uuid_generator)
		case "gorilla_mux":
			return httpRouterGorillaMux.Init(uuid_generator)
		case "julienschmidt_httprouter":
			return httpRouterJulienschmidtHTTPRouter.Init(uuid_generator)
		default:
			err := "Must be defined a adapter for http router"
			fmt.Println(err)
			panic(err)
	}
}
