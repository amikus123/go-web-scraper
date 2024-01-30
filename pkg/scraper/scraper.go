package scraper

import (
	"context"
	"log"

	"github.com/amikus123/go-web-scraper/pkg/screenshoter"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type Scraper struct {
	ctx          context.Context
	Screenshoter screenshoter.ScreenshoterBehaviour
	// config
	Config ScraperConfig
}

type ScraperBehaviour interface {
	Scrape()
	takeScreenshot()
	scrapeHeadings() []ScrapedNewsItem
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

func (s *Scraper) takeScreenshot() *[]byte {
	buf, err := s.Screenshoter.TakeScreenshot(s.ctx, s.Config.Url)

	if err != nil {
		panic(err)
	}

	return &buf
}

func (s *Scraper) scrapeHeadings() *[]ScrapedNewsItem {

	var res []ScrapedNewsItem

	// navigate to the target web page and select the HTML elements of interest
	var nodes []*cdp.Node
	selectors := s.Config.Selectors

	chromedp.Run(s.ctx,
		chromedp.Navigate(s.Config.Url),
		chromedp.Nodes(selectors.Main, &nodes, chromedp.ByQueryAll),
	)

	var title, imageSrc, href string
	for _, node := range nodes {

		chromedp.Run(s.ctx,
			chromedp.AttributeValue(selectors.Href, "href", &title, nil, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.AttributeValue(selectors.Img, "src", &imageSrc, nil, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(selectors.Text, &href, chromedp.ByQuery, chromedp.FromNode(node)),
		)
		newsItem := ScrapedNewsItem{}

		newsItem.Title = title
		newsItem.Href = href
		newsItem.ImageSrc = imageSrc

		res = append(res, newsItem)
	}
	return &res
}

func (s *Scraper) Scrape() *ScrapedData {
	cancel := s.init()

	defer cancel()
	// screenshot := s.takeScreenshot()
	headings := s.scrapeHeadings()

	// save screenshot and return link to it

	return &ScrapedData{
		SourceID:   s.Config.SourceID,
		Items:      headings,
		Screenshot: "",
	}
}
