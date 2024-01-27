package headless

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/amikus123/go-web-scraper/news"
	"github.com/amikus123/go-web-scraper/screenshoter"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type Scraper struct {
	Url          string
	ctx          context.Context
	Screenshoter screenshoter.ScreenshoterBehaviour
	// config
	Selectors HeadlessSelectors
}

type ScraperBehaviour interface {
	Scrape()
	takeScreenshot()
	scrapeHeadings() []news.NewsItem
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
	fmt.Println(buf)
	// if err := os.WriteFile("fullScreenshot.png", buf, 0o644); err != nil {
	// 	log.Fatal(err)
	// }

}

func (s *Scraper) scrapeHeadings() []news.NewsItem {

	var res []news.NewsItem

	// navigate to the target web page and select the HTML elements of interest
	var nodes []*cdp.Node
	chromedp.Run(s.ctx,
		chromedp.Navigate(s.Url),
		chromedp.Nodes(s.Selectors.Main, &nodes, chromedp.ByQueryAll),
	)

	var title, imageSrc, href string
	for _, node := range nodes {

		chromedp.Run(s.ctx,
			chromedp.AttributeValue(s.Selectors.Href, "href", &title, nil, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.AttributeValue(s.Selectors.Img, "src", &imageSrc, nil, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(s.Selectors.Text, &href, chromedp.ByQuery, chromedp.FromNode(node)),
		)
		newsItem := news.NewsItem{}

		newsItem.Title = title
		newsItem.Href = href
		newsItem.ImageSrc = imageSrc
		newsItem.ScrapedAt = time.Now()

		res = append(res, newsItem)
	}
	return res
}

func (s *Scraper) Scrape() {
	cancel := s.init()

	defer cancel()
	// s.takeScreenshot()
	res := s.scrapeHeadings()
	fmt.Println(res)
}
