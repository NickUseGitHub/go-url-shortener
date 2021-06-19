package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	routes := mux.NewRouter()

	routes.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	}).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", routes))
}