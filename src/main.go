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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("[X]:: DB not connect")
	}

	fmt.Println("[√]:: DB Connected")

	// migration
	isShortenUrlTableExisted := db.Migrator().HasTable(&models.ShortenUrl{})
	if isShortenUrlTableExisted != true {
		db.Migrator().CreateTable(&models.ShortenUrl{})
	}

	routes := mux.NewRouter()

	routes.HandleFunc("/shorten-url", func(w http.ResponseWriter, r *http.Request) {
		sUrlReq := models.ShortenUrl{}
		json.NewDecoder(r.Body).Decode(&sUrlReq)

		var sUrl models.ShortenUrl
		db.Where("url = ?", sUrlReq.Url).Find(&sUrl)

		if sUrl.ID != 0 {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(struct {
				Message string `json:"message"`
			}{
				Message: fmt.Sprintf("'%s' has already had hashedUrl", sUrlReq.Url),
			})
			return
		}

		shortenUrl := models.ShortenUrl{
			HashedUrl: models.GetShortenUrl(),
			Url:       sUrlReq.Url,
		}

		db.Create(&shortenUrl)

		json.NewEncoder(w).Encode(shortenUrl)
	}).Methods(http.MethodPost)

	routes.HandleFunc("/{shortenUrl:.+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		shortenUrl := vars["shortenUrl"]

		w.Write([]byte(fmt.Sprintf("Your shortenUrl is '%s'", shortenUrl)))
	}).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", routes))
}
