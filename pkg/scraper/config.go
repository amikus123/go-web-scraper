package scraper

type ScraperSelectors struct {
	Main string
	Href string
	Img  string
	Text string
}
type ScraperConfig struct {
	SourceID  int64
	Url       string
	Selectors *ScraperSelectors
}

var FrondaConfig = ScraperConfig{
	SourceID: 1,
	Url:      "https://www.fronda.pl/",
	Selectors: &ScraperSelectors{Main: ".itemBox",
		Href: "a",
		Img:  "img",
		Text: ".itemBox-title"},
}
