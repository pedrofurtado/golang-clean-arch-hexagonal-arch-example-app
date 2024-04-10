package main

import (
	"time"
	"fmt"
	"runtime"
	"net/http"

	// Infra - HTTP Routers
	http_router_chi "my-app/src/infra/http_routers/chi"
	http_router_gorilla_mux "my-app/src/infra/http_routers/gorilla_mux"
	http_router_julienschmidt_httprouter "my-app/src/infra/http_routers/julienschmidt_httprouter"
)

const APP_PORT = 80

var INFRA_ADAPTERS = map[string]string{
	// Possible values: chi | gorilla_mux | julienschmidt_httprouter
	"HTTP_ROUTER": "gorilla_mux",
}

func main() {
	fmt.Printf("App listening on port %v | Golang version %v | Now is %s\n", APP_PORT, runtime.Version(), time.Now().String())

	// Init infra dependencies
	router := initHTTPRouter()

	// Init app
	http.ListenAndServe(fmt.Sprintf(":%d", APP_PORT), router)
}

func initHTTPRouter() http.Handler {
	switch INFRA_ADAPTERS["HTTP_ROUTER"] {
		case "chi":
			return http_router_chi.Init()
		case "gorilla_mux":
			return http_router_gorilla_mux.Init()
		case "julienschmidt_httprouter":
			return http_router_julienschmidt_httprouter.Init()
		default:
			return nil
	}
}
