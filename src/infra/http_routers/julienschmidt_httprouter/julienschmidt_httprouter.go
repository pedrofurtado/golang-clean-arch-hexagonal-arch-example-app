package http_servers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Init() http.Handler {
	fmt.Println("Http router | Init | Implementation: Julienschmidt Httprouter")

	r := httprouter.New()
	r.GET("/", homeRoute)
	r.GET("/hello", helloRoute)
	r.GET("/show/:my_param", showMyParamRoute)

	return r
}

func homeRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("welcome from Julienschmidt Httprouter"))
}

func helloRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("hello world from Julienschmidt Httprouter"))
}

func showMyParamRoute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param_value := ps.ByName("my_param")
	w.Write([]byte(fmt.Sprintf("show my param from Julienschmidt Httprouter: %s", param_value)))
}
