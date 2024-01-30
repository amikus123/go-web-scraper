package main

import (
	"fmt"

	"github.com/amikus123/go-web-scraper/db"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	DB := db.Connect()

	defer DB.Close()

	// scraper := scraper.Scraper{
	// Screenshoter: screenshoter.Screenshoter{},
	// Config:       scraper.FrondaConfig,
	// }

	// res := scraper.Scrape()

	// rep := db.ScrapeResultRepository{DB: DB}
	srcRep := db.SourceRepository{DB: DB}

	res2, err := srcRep.Get()
	fmt.Println(res2, err)
	// rep.Save(res)
	// connect to d
	fmt.Println("end")

}
