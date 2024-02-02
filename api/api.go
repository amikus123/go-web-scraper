package api

import (
	"fmt"
	"net/http"
)

func StartAPIServer() {

	// handler function #1 - returns the index.html template, with film data
	h1 := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HeloWorld")
	}

	http.HandleFunc("/api", h1)

}
