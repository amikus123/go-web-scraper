package main

import (
	"fmt"

	"github.com/amikus123/go-web-scraper/headless"
	"github.com/amikus123/go-web-scraper/screenshoter"
)

func main() {
	scraper := headless.Scraper{
		Url:          "https://www.fronda.pl/",
		Screenshoter: screenshoter.Screenshoter{},
		Selectors:    headless.FrondaSelectors,
	}
	scraper.Scrape()
	fmt.Println("end")
}
