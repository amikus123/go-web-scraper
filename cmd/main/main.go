package main

import (
	"log"
	"net/http"

	"github.com/amikus123/go-web-scraper/api"
	"github.com/amikus123/go-web-scraper/db"
	"github.com/amikus123/go-web-scraper/web"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	DB := db.Connect()
	defer DB.Close()

	web.StartWebServer(DB)
	api.StartAPIServer(DB)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
