package http_routers

import (
	"net/http"

	infra_adapters "my-app/src/config/infra_adapters"

	infra_uuid_generator "my-app/src/infra/uuid_generators"

	http_router_chi "my-app/src/infra/http_routers/chi"
	http_router_gorilla_mux "my-app/src/infra/http_routers/gorilla_mux"
	http_router_julienschmidt_httprouter "my-app/src/infra/http_routers/julienschmidt_httprouter"
)

func Init(uuid_generator infra_uuid_generator.GenericUUIDGenerator) http.Handler {
	switch infra_adapters.GetAdapters()["HTTP_ROUTER"] {
		case "chi":
			return http_router_chi.Init(uuid_generator)
		case "gorilla_mux":
			return http_router_gorilla_mux.Init(uuid_generator)
		case "julienschmidt_httprouter":
			return http_router_julienschmidt_httprouter.Init(uuid_generator)
		default:
			return nil
	}
}
