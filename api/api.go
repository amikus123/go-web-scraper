package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/amikus123/go-web-scraper/db"
	"github.com/amikus123/go-web-scraper/pkg/scraper"
	"github.com/amikus123/go-web-scraper/pkg/screenshoter"
)

func StartAPIServer(DB *sql.DB) {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		handleScrape(DB)
		fmt.Fprintf(w, "Scraped!")
	}

	http.HandleFunc("/api/scrape", h1)
	fmt.Println("started api server")
}

func handleScrape(DB *sql.DB) {
	srcRep := db.SourceRepository{DB: DB}
	scrResRep := db.ScrapeResultRepository{DB: DB}
	newsItemRep := db.NewsItemRepository{DB: DB}

	srcs, err := srcRep.GetRecentlyNotUsed()

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

		if err = newsItemRep.Save(*scrapedData.Items, id); err != nil {
			panic(err)

		}

		src.LastScrapeAt = time.Now()

		if err = srcRep.Update(src); err != nil {
			panic(err)
		}

	}
}
