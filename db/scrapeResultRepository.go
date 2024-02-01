package db

import (
	"database/sql"
	"fmt"
)

type ScrapeResult struct {
	ID         int64
	SourceID   int64
	Screenshot string
}

type ScrapeResultRepository struct {
	DB *sql.DB
}

func (*ScrapeResultRepository) Save(sourceId int64, screenshotUrl string) (int64, error) {

	result, err := db.Exec("INSERT INTO scrape_result (source_id, screenshot) VALUES (?, ?)", sourceId, screenshotUrl)

	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()

	return id, err
}
