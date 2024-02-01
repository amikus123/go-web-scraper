package db

import (
	"database/sql"
	"time"
)

type Source struct {
	ID        int64
	Url       string
	Name      string
	Selectors []Selector
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SourceRepository struct {
	DB *sql.DB
}

func (*SourceRepository) Get() ([]Source, error) {

	rows, err := db.Query("SELECT source.id, url, name, selector.id, main_selector,  text_selector, href_selector, img_selector FROM source LEFT JOIN selector ON source.id = selector.source_id ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

	if err = rows.Err(); err != nil {
		return nil, err
	}

	var sources []Source

	for _, val := range sourceMap {
		sources = append(sources, *val)
	}
	return sources, nil
}
