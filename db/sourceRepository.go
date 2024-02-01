package db

import (
	"database/sql"
	"time"
)

type Source struct {
	ID           int64
	Url          string
	Name         string
	Selectors    []Selector
	LastScrapeAt time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type SourceRepository struct {
	DB *sql.DB
}

func (s *SourceRepository) Get() ([]Source, error) {

	sqlStatement := `
	SELECT source.id, url, name, selector.id, main_selector,
	text_selector, href_selector, img_selector
	FROM source LEFT JOIN selector
	ON source.id = selector.source_id;`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return s.groupSourcesWithSelectors(rows)
}
func (s *SourceRepository) GetSourcesToScrape() ([]Source, error) {

	sqlStatement := `
	SELECT source.id, url, name, selector.id, main_selector,
	text_selector, href_selector, img_selector
	FROM source LEFT JOIN selector
	ON source.id = selector.source_id
	WHERE last_scrape_at IS NULL 
	OR last_scrape_at <= DATE_SUB(NOW(), INTERVAL 23 HOUR);`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return s.groupSourcesWithSelectors(rows)
}

func (*SourceRepository) groupSourcesWithSelectors(rows *sql.Rows) ([]Source, error) {
	var sourceMap = make(map[int64]*Source)
	for rows.Next() {
		var sourceID, selectorID int64
		var sourceName, sourceUrl, mainSelector, imgSelector, hrefSelector, textSelector string

		if err := rows.Scan(&sourceID, &sourceUrl, &sourceName, &selectorID, &mainSelector, &textSelector, &hrefSelector, &imgSelector); err != nil {
			return nil, err
		}
		source, exists := sourceMap[sourceID]
		if !exists {
			source = &Source{
				ID:        sourceID,
				Name:      sourceName,
				Url:       sourceUrl,
				Selectors: make([]Selector, 0),
			}
			sourceMap[sourceID] = source

		}
		selector := Selector{
			ID:           selectorID,
			SourceID:     sourceID,
			MainSelector: mainSelector,
			TextSelector: textSelector,
			ImgSelector:  imgSelector,
			HrefSelector: hrefSelector,
		}
		source.Selectors = append(source.Selectors, selector)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	var sources []Source

	for _, val := range sourceMap {
		sources = append(sources, *val)
	}
	return sources, nil
}

func (*SourceRepository) Update(s Source) error {
	sqlStatement := `
	UPDATE source
	SET name = ?, url = ?, last_scrape_at = ?
	WHERE id = ?;`
	_, err := db.Exec(sqlStatement, s.Name, s.Url, s.LastScrapeAt, s.ID)

	return err
}
