package db

import (
	"database/sql"
	"fmt"

	"github.com/amikus123/go-web-scraper/pkg/scraper"
)

type NewsItem struct {
	ID         int64
	Text       string
	Href       string
	Screenshot string
}

type NewsItemRepository struct {
	DB *sql.DB
}

func (*NewsItemRepository) Save(data *scraper.ScrapedData) (int64, error) {

	result, err := db.Exec("INSERT INTO scrape_result (source_id, screenshot) VALUES (?, ?)", data.SourceID, data.Screenshot)

	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
