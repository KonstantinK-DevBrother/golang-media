package main

import (
	"golang-media/src/web"
	"log"
	"net/http"
	"runtime"

	"github.com/gorilla/mux"
)

func main() {
	runtime.GOMAXPROCS(1)

	r := mux.NewRouter()
	r.HandleFunc("/video/{filename}", web.Metrics(web.HandleVideo))

	r.HandleFunc("/", web.HandleHTML)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))
	r.Use(mux.CORSMethodMiddleware(r))

	log.Fatal(http.ListenAndServe(":3333", r))
}
