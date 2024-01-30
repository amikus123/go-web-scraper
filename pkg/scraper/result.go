package scraper

type ScrapedNewsItem struct {
	Title    string
	ImageSrc string
	Href     string
	// to verify
	InViewport bool
}

type ScrapedData struct {
	SourceID   int64
	Items      *[]ScrapedNewsItem
	Screenshot string
}
