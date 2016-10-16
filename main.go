package main

import (
	"database/sql"
	"net/http"

	"server/g"

	"server/lib/api"
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
	//API
	route.HandleFunc("/api/pages", api.Pages)
	route.HandleFunc("/api/pages/{guid:[0-9a-zA-Z-]+}", api.Pages)

	route.HandleFunc("/api/comments", api.CommentPut).Methods("post")

	route.HandleFunc("/pages/{id:[0-9a-zA\\-]+}", handlers.ServePageByGUID)
	route.HandleFunc("/", handlers.RedirIndex)
	route.HandleFunc("/pages", handlers.ServePages)
	http.Handle("/", route)
	//静态文件系统
	http.Handle("/assets/", http.FileServer(http.Dir("./")))

	go func() {
		err = http.ListenAndServeTLS("127.0.0.1:8080", "certificate.pem", "key.pem", nil)
		if err != nil {
			panic(err.Error())
		}
	}()
	go func() {
		err = http.ListenAndServe("127.0.0.1:8090", handlers.RedirectTo("https://127.0.0.1:8080"))
		if err != nil {
			panic(err.Error())
		}
	}()
	select {}
}
