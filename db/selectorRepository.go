package db

import (
	"database/sql"
	"fmt"
)

type Selector struct {
	ID           int64
	SourceID     int64
	MainSelector string
	TextSelector string
	ImgSelector  string
	HrefSelector string
}

type SelectorRepository struct {
	DB *sql.DB
}

func (*SelectorRepository) Save() (int64, error) {

	sqlStatement := `
	INSERT INTO scrape_result 
	(source_id, screenshot) VALUES (?, ?)`

	result, err := db.Exec(sqlStatement, 1, 1)

	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
