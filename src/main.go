package main

import (
	"encoding/json"
	"github.com/dchest/uniuri"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	routes := mux.NewRouter()

	routes.HandleFunc("/shorten-url", func(w http.ResponseWriter, r *http.Request) {
		sUrl := ShortenUrl{}

		json.NewDecoder(r.Body).Decode(&sUrl)

		s := uniuri.New()
		w.Write([]byte(s))
	}).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8000", routes))
}
