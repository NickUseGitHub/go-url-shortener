package main

import (
	"encoding/json"
	"fmt"
	"github.com/NickUseGitHub/url-shortener/src/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=shortenurl port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("[X]:: DB not connect")
	}

	fmt.Println("[√]:: DB Connected")

	routes := mux.NewRouter()

	routes.HandleFunc("/shorten-url", func(w http.ResponseWriter, r *http.Request) {
		sUrl := models.ShortenUrl{}
		json.NewDecoder(r.Body).Decode(&sUrl)

		shortenUrl := sUrl.GetShortenUrl()
		w.Write([]byte(shortenUrl))
	}).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8000", routes))
}
