package headless

import (
	"context"
	"log"
	"os"
	"web-scraper/screenshoter"

	"github.com/chromedp/chromedp"
)

type NewsItem struct {
	title string
}

type Scraper struct {
	Url          string
	ctx          context.Context
	Screenshoter screenshoter.ScreenshoterBehaviour
	// config
}

type ScraperBehaviour interface {
	Scrape()
	takeScreenshot()
	scrapeHeadings()
	init()
}

func (s *Scraper) init() context.CancelFunc {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	s.ctx = ctx
	return cancel
}

func (s *Scraper) takeScreenshot() {
	buf, err := s.Screenshoter.TakeScreenshot(s.ctx, s.Url)

	if err != nil {
		panic(err)
	}

	// if err := os.WriteFile("fullScreenshot.png", buf, 0o644); err != nil {
	// 	log.Fatal(err)
	// }

}

func (s *Scraper) Scrape() {

	cancel := s.init()

	defer cancel()
	s.takeScreenshot()
}
