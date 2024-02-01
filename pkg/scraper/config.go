package scraper

import "github.com/amikus123/go-web-scraper/db"

type ScraperSelectors struct {
	Main string
	Href string
	Img  string
	Text string
}
type ScraperConfig struct {
	SourceID  int64
	Url       string
	Selectors []db.Selector
}
