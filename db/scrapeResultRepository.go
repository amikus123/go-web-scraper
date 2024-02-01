package db

import (
	"database/sql"
	"fmt"
)

type ScrapeResult struct {
	ID         int64
	SourceID   int64
	Screenshot string
	Items      []NewsItem
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

func (*ScrapeResultRepository) Get(scrapeResultId int64) (*ScrapeResult, error) {

	rows, err := db.Query("SELECT id, source_id, screenshot FROM scrape_result WHERE id=?", scrapeResultId)

	if err != nil {
		return nil, fmt.Errorf("addAlbum: %v", err)
	}
	defer rows.Close()

	res := ScrapeResult{}
	for rows.Next() {
		var ID, sourceID int64
		var screenshot string

		if err := rows.Scan(&ID, &sourceID, &screenshot); err != nil {
			return nil, err
		}

		res.ID = ID
		res.Screenshot = screenshot
		res.SourceID = sourceID
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &res, nil
}

func (*ScrapeResultRepository) GetWithNewsItems(scrapeResultID int64) ([]ScrapeResult, error) {

	rows, err := db.Query("SELECT scrape_result.id, scrape_result.screenshot, scrape_result.source_id, news_item.id, news_item.text,  news_item.href, news_item.img FROM scrape_result LEFT JOIN news_item ON scrape_result.id = news_item.scrape_result_id WHERE scrape_result.id=?", scrapeResultID)

	if err != nil {
		return nil, fmt.Errorf("addAlbum: %v", err)
	}
	defer rows.Close()

	var scrResultsMap = make(map[int64]*ScrapeResult)

	for rows.Next() {
		var scrapedResultID, newsItemID int64
		var text, screenshot, href, img string
		var sourceID sql.NullInt64
		fmt.Println()
		if err := rows.Scan(&scrapedResultID, &screenshot, &sourceID, &newsItemID, &text, &href, &img); err != nil {
			return nil, err
		}
		scrResult, exists := scrResultsMap[scrapedResultID]

		if !exists {
			scrResult = &ScrapeResult{
				ID:         scrapedResultID,
				SourceID:   1,
				Screenshot: screenshot,
				Items:      make([]NewsItem, 0),
			}
			scrResultsMap[scrapedResultID] = scrResult

		}
		newsItem := NewsItem{
			ID:             newsItemID,
			Text:           text,
			Href:           href,
			Img:            img,
			ScrapeResultID: scrapedResultID,
		}
		scrResult.Items = append(scrResult.Items, newsItem)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	var scrResults []ScrapeResult

	for _, val := range scrResultsMap {
		scrResults = append(scrResults, *val)
	}
	return scrResults, nil
}
