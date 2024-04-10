package http_servers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() http.Handler {
	fmt.Println("Http router | Init | Implementation: Gorilla Mux")

	r := mux.NewRouter()
	r.HandleFunc("/", homeRoute)
	r.HandleFunc("/hello", helloRoute)
	r.HandleFunc("/show/{my_param}", showMyParamRoute)

	return r
}

func homeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome from gorilla mux"))
}

func helloRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world from gorilla mux"))
}

func showMyParamRoute(w http.ResponseWriter, r *http.Request) {
	param_value := mux.Vars(r)["my_param"]
	w.Write([]byte(fmt.Sprintf("show my param from gorilla mux: %s", param_value)))
}
