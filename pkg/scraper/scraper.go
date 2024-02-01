package scraper

import (
	"context"
	"log"
	"strings"

	"github.com/amikus123/go-web-scraper/db"
	"github.com/amikus123/go-web-scraper/pkg/screenshoter"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type Scraper struct {
	ctx          context.Context
	Screenshoter screenshoter.ScreenshoterBehaviour
	Config       ScraperConfig
}

type ScrapedData struct {
	SourceID   int64
	Items      *[]db.NewsItem
	Screenshot string
}

type ScraperBehaviour interface {
	Scrape()
	takeScreenshot()
	scrapeHeadings() []db.NewsItem
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

func (s *Scraper) scrapeHeadings() *[]db.NewsItem {

	var res []db.NewsItem

	// navigate to the target web page and select the HTML elements of interest
	var nodes []*cdp.Node
	selectors := s.Config.Selectors[0]

	chromedp.Run(s.ctx,
		chromedp.Navigate(s.Config.Url),
		chromedp.Nodes(selectors.MainSelector, &nodes, chromedp.ByQueryAll),
	)

	var title, imageSrc, href string
	for _, node := range nodes {

		chromedp.Run(s.ctx,
			chromedp.AttributeValue(selectors.HrefSelector, "href", &href, nil, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.AttributeValue(selectors.ImgSelector, "src", &imageSrc, nil, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(selectors.TextSelector, &title, chromedp.ByQuery, chromedp.FromNode(node)),
		)
		newsItem := db.NewsItem{}

		if strings.HasPrefix(href, "/") {
			href = s.Config.Url + href
		}

		newsItem.Text = title
		newsItem.Href = href
		newsItem.Img = imageSrc

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
