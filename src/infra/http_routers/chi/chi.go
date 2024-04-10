package http_servers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Init() http.Handler {
	fmt.Println("Http router | Init | Implementation: Chi")

	r := chi.NewRouter()
	r.Get("/", homeRoute)
	r.Get("/hello", helloRoute)
	r.Get("/show/{my_param}", showMyParamRoute)

	return r
}

func homeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome from chi"))
}

func helloRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world from chi"))
}

func showMyParamRoute(w http.ResponseWriter, r *http.Request) {
	param_value := chi.URLParam(r, "my_param")
	w.Write([]byte(fmt.Sprintf("show my param from chi: %s", param_value)))
}
