package db

import (
	"database/sql"
	"fmt"
)

type Selector struct {
	ID       int64
	SourceID int64
	Main     string
	Text     string
	Img      string
	Href     string
}

type SelectorRepository struct {
	DB *sql.DB
}

func (*SelectorRepository) Save(selector Selector) (int64, error) {

	sqlStatement := `
	INSERT INTO selector
	(source_id,main,  text, img, href)
	VALUES (?, ?, ?, ?, ?)`

	result, err := db.Exec(sqlStatement,
		selector.SourceID, selector.Main, selector.Text, selector.Img, selector.Href)

	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
