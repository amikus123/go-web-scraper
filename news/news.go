package news

import "time"

type NewsImportanceLevel int

const (
	High NewsImportanceLevel = iota
	Medium
	Low
)

type NewsItem struct {
	Title     string
	ImageSrc  string
	Href      string
	ScrapedAt time.Time
	// to verify
	InViewport      bool
	ImportanceLevel NewsImportanceLevel
}
