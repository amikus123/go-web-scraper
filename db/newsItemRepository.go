package db

import (
	"database/sql"
	"fmt"
	"strings"
)

type NewsItem struct {
	ID             int64
	Text           string
	Href           string
	Img            string
	ScrapeResultID int64
}

type NewsItemRepository struct {
	DB *sql.DB
}

func (*NewsItemRepository) Save(data []NewsItem, scrapeResultID int64) error {

	sqlStr := "INSERT INTO news_item (scrape_result_id, text, href, img) VALUES"
	vals := []interface{}{}

	for _, row := range data {
		sqlStr += "(?, ?, ?, ?),"
		vals = append(vals, scrapeResultID, row.Text, row.Href, row.Img)
	}
	sqlStr = strings.TrimSuffix(sqlStr, ",")
	// prepare the statement

	stmt, err := db.Prepare(sqlStr)

	if err != nil {
		panic(err)
	}
	// format all vals at once
	_, err = stmt.Exec(vals...)

	if err != nil {
		return fmt.Errorf("addAlbum: %v", err)
	}

	return nil
}
