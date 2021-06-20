package main

import (
	"encoding/json"
	"github.com/NickUseGitHub/url-shortener/src/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	routes := mux.NewRouter()

	routes.HandleFunc("/shorten-url", func(w http.ResponseWriter, r *http.Request) {
		sUrl := models.ShortenUrl{}
		json.NewDecoder(r.Body).Decode(&sUrl)

		shortenUrl := sUrl.GetShortenUrl()
		w.Write([]byte(shortenUrl))
	}).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8000", routes))
}
