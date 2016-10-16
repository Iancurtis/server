package model

import (
	"errors"
	"html/template"
	"log"
	"server/g"
)

// Page struct of blog
type Page struct {
	ID         string
	GUID       string
	Title      string
	RawContent string
	Content    template.HTML
	Date       string
	Comments   []Comment
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

// GetComments get all comments associate with this page.
func (p *Page) GetComments() error {
	//fmt.Printf("comments addr :%p\n", &OutComments)
	rows, err := g.Database.Query("select comment_name, comment_email, comment_text from comments where page_id=?", p.ID)
	if err != nil {
		log.Println(err.Error())
		return errors.New("Failed to get data from databse.")
	}
	defer rows.Close()
	thisComment := Comment{}
	for rows.Next() {
		rows.Scan(&thisComment.Name, &thisComment.Email, &thisComment.CommentText)
		p.Comments = append(p.Comments, thisComment)
	}
	return nil
}
