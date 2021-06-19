package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	routes := mux.NewRouter()

	routes.HandleFunc("/shorten-url", func (w http.ResponseWriter, r *http.Request) {
		sUrl := ShortenUrl{}

		json.NewDecoder(r.Body).Decode(&sUrl)

		json.NewEncoder(w).Encode(sUrl)
	}).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8000", routes))
}