package main

import (
	"fmt"
	"web-scraper/headless"
	"web-scraper/screenshoter"
)

func main() {
	scraper := headless.Scraper{Url: "https://www.fronda.pl/", Screenshoter: screenshoter.Screenshoter{}}
	scraper.Scrape()
	fmt.Println("end")
}
