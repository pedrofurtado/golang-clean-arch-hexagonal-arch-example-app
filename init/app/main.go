package main

import (
	"time"
	"fmt"
	"runtime"
	"net/http"

	// Infra - HTTP Routers
	infra_http_router "my-app/src/infra/http_routers"

	// Infra - UUID Generator
	infra_uuid_generator "my-app/src/infra/uuid_generators"
)

const APP_PORT = 80

func main() {
	// Init - Log
	fmt.Printf("App listening on port %v | Golang version %v | Now is %s\n", APP_PORT, runtime.Version(), time.Now().String())

	// Init - Infra dependencies
	infra_uuid_generator := infra_uuid_generator.Init()
	router := infra_http_router.Init(infra_uuid_generator)

	// Init app
	http.ListenAndServe(fmt.Sprintf(":%d", APP_PORT), router)
}
