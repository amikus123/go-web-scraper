package main

import (
	"time"

	"github.com/amikus123/go-web-scraper/db"
	"github.com/amikus123/go-web-scraper/pkg/scraper"
	"github.com/amikus123/go-web-scraper/pkg/screenshoter"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	DB := db.Connect()

	defer DB.Close()

	srcRep := db.SourceRepository{DB: DB}
	scrResRep := db.ScrapeResultRepository{DB: DB}
	newsItemRep := db.NewsItemRepository{DB: DB}

	srcs, err := srcRep.GetSourcesToScrape()

	if err != nil {
		panic(err)
	}

	for _, src := range srcs {
		scraper := scraper.Scraper{
			Screenshoter: screenshoter.Screenshoter{},
			Config: scraper.ScraperConfig{
				SourceID:  src.ID,
				Url:       src.Url,
				Selectors: src.Selectors,
			},
		}
		scrapedData := scraper.Scrape()

		id, err := scrResRep.Save(scrapedData.SourceID, "")

		if err != nil {
			panic(err)
		}

		err = newsItemRep.Save(*scrapedData.Items, id)

		if err != nil {
			panic(err)
		}
		src.LastScrapeAt = time.Now()

		err = srcRep.Update(src)

		if err != nil {
			panic(err)
		}
	}

}
