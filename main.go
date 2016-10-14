package main

import (
	"database/sql"
	"net/http"

	"server/g"

	"server/lib/handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/cms")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	g.Database = db

	route := mux.NewRouter()
	route.HandleFunc("/page/{id:[0-9a-zA\\-]+}", handlers.ServePageByGUID)
	http.Handle("/", route)

	http.ListenAndServe(":8080", nil)

}
