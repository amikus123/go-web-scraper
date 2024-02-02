package web

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/amikus123/go-web-scraper/db"
)

func getTemplateURL(template string) string {

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return wd + "/web/" + template + ".html"
}

func StartWebServer(DB *sql.DB) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleIndex(w, r, DB)
	})

	http.HandleFunc("/source", func(w http.ResponseWriter, r *http.Request) {
		handleSource(w, r, DB)
	})

	fmt.Println("started web server")

}

func handleIndex(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	srcRep := db.SourceRepository{DB: DB}

	tmplURL := getTemplateURL("index")
	tmpl := template.Must(template.ParseFiles(tmplURL))

	// handle incorrect routes
	if r.URL.Path != "/" {
		fmt.Println("redirect", r.URL)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	sources, err := srcRep.Get()

	if err != nil {
		panic(err)
	}

	tmplMap := map[string][]db.Source{
		"Sources": sources,
	}

	tmpl.Execute(w, tmplMap)

}

type SourcePageData struct {
	Source   *db.Source
	AddedNew bool
}

func handleSource(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	srcRep := db.SourceRepository{DB: DB}
	selRep := db.SelectorRepository{DB: DB}

	tmplURL := getTemplateURL("source")
	tmpl := template.Must(template.ParseFiles(tmplURL))

	query := r.URL.Query()

	paramSourceID := query.Get("id")

	sourceID, err := strconv.ParseInt(paramSourceID, 10, 64)

	if err != nil {
		panic(err)
	}

	// handle POSt
	if r.Method == http.MethodPost {
		time.Sleep(1 * time.Second)
		main := r.PostFormValue("selector-main")
		text := r.PostFormValue("selector-text")
		img := r.PostFormValue("selector-img")
		href := r.PostFormValue("selector-href")

		selRep.Save(db.Selector{
			Main:     main,
			Text:     text,
			Img:      img,
			Href:     href,
			SourceID: sourceID,
		})

		source, err := srcRep.GetWithSelectorsByID(sourceID)

		if err != nil {
			panic(err)
		}

		tmplMap := SourcePageData{
			Source:   source,
			AddedNew: true,
		}

		tmpl.Execute(w, tmplMap)
		return
	}

	// handle GET
	if r.Method == http.MethodGet {
		source, err := srcRep.GetWithSelectorsByID(sourceID)

		if err != nil {
			panic(err)
		}
		tmplMap := SourcePageData{
			Source:   source,
			AddedNew: false,
		}

		tmpl.Execute(w, tmplMap)
		return
	}
	// handle DELETE

	// handle UPDATE

}
