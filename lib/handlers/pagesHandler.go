package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"server/g"

	"github.com/gorilla/mux"
)

// Page struct of blog
type Page struct {
	ID         string
	GUID       string
	Title      string
	RawContent string
	Content    template.HTML
	Date       string
}

// TruncatedContent get a shortened content
func (p *Page) TruncatedContent() template.HTML {
	for i := range p.Content {
		if i >= 10 {
			return p.Content[:i] + "..."
		}
	}
	return p.Content
}

func servePage(w http.ResponseWriter, r *http.Request, field string) {
	vars := mux.Vars(r)
	filter := vars["id"]
	thisPage := Page{}
	err := g.Database.QueryRow("select id, page_title, page_content, page_date from pages where "+field+"=?", filter).Scan(&thisPage.ID, &thisPage.Title, &thisPage.RawContent, &thisPage.Date)
	thisPage.Content = template.HTML(thisPage.RawContent)
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

// ServePages handle the pages request
// show the pages list
func ServePages(w http.ResponseWriter, r *http.Request) {
	pages := []Page{}
	rows, err := g.Database.Query("select page_guid, page_title, page_content, page_date from pages order by ? desc", "page_date")
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		thisPage := Page{}
		rows.Scan(&thisPage.GUID, &thisPage.Title, &thisPage.RawContent, &thisPage.Date)
		thisPage.Content = template.HTML(thisPage.RawContent)
		pages = append(pages, thisPage)
	}
	t, err := template.ParseFiles("tpl/blog-list.tpl")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	t.Execute(w, pages)
	//fmt.Fprintln(w, pages)
}

// BUG(src): Mapping between XML elements and data structures is inherently flawed:
// an XML element is an order-dependent collection of anonymous

// BUG|TODO:  heheh
