package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"server/g"
	"server/lib/handlers"

	"github.com/gorilla/mux"
)

// Pages deal with api request about pages
func Pages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageGUID := vars["guid"]
	fmt.Println(pageGUID)
	thisPage := handlers.Page{}
	err := g.Database.QueryRow("select id, page_title, page_content, page_date from pages where page_guid=?", pageGUID).Scan(&thisPage.ID, &thisPage.Title, &thisPage.RawContent, &thisPage.Date)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	thisPage.Content = template.HTML(thisPage.RawContent)

	output, err := json.Marshal(thisPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// fmt.Println(string(output))

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, (string)(output))
}
