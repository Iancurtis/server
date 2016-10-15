package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/g"
	"strconv"
)

// CommentPut add a comment
func CommentPut(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error())
		return
	}

	pageID := r.FormValue("page_id")
	name := r.FormValue("name")
	email := r.FormValue("email")
	content := r.FormValue("content")

	if pageID == "" || name == "" || content == "" || email == "" {
		// TODO: make a json lib
		log.Println("Field cannot be nil")
		return
	}

	res, err := g.Database.Exec("insert into comments set page_id=?, comment_name=?, comment_email=?, comment_text=?", pageID, name, email, content)
	if err != nil {
		log.Println(err.Error())
		return
	}

	id, err := res.LastInsertId()

	var isAdded bool

	if err != nil {
		isAdded = false
	} else {
		isAdded = true
	}

	//jsonResponse := JSONResponse{}
	response := JSONResponse{map[string]string{}}
	response.Fields["id"] = strconv.FormatInt(id, 10) //string(id)
	response.Fields["isAdded"] = strconv.FormatBool(isAdded)

	jsonResp, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonResp))

}
