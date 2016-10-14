package handlers

import (
	"html/template"
	"log"
	"net/http"

	"server/g"

	"github.com/gorilla/mux"
)

// Page struct of blog
type Page struct {
	Title   string
	Content string
	Date    string
}

func servePage(w http.ResponseWriter, r *http.Request, field string) {
	vars := mux.Vars(r)
	filter := vars["id"]
	thisPage := Page{}
	err := g.Database.QueryRow("select page_title, page_content, page_date from pages where "+field+"=?", filter).Scan(&thisPage.Title, &thisPage.Content, &thisPage.Date)
	if err != nil {
		log.Println("Couldn't get page :", filter)
		// log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	// html := `<head><title>` + thisPage.title + `</title></head><body><h1>` + thisPage.title + `</h1><div>` + thisPage.content + `</div>`
	// fmt.Fprintln(w, html)
	t, err := template.ParseFiles("tpl/page.tpl")
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("ok")
	}
	t.Execute(w, thisPage)
}

// ServePageByID handle the page request by id
func ServePageByID(w http.ResponseWriter, r *http.Request) {
	servePage(w, r, "id")
}

// ServePageByGUID handle the page request by guid
func ServePageByGUID(w http.ResponseWriter, r *http.Request) {
	servePage(w, r, "page_guid")
}

// BUG(src): Mapping between XML elements and data structures is inherently flawed:
// an XML element is an order-dependent collection of anonymous

// BUG|TODO:  heheh
