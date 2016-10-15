/*
Package handlers implements a library for handling request from clients
*/
package handlers

import (
	"fmt"
	"net/http"
)

// func RedirectHttpToTl(w http.ResponseWriter, r *http.Request) {
// 	http.Redirect(w, req, )
// }

// RedirectTo request to addr
func RedirectTo(addr string) http.HandlerFunc {
	println("addr is", addr)
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("uri:", r.RequestURI)
		fmt.Println("url:", addr+r.RequestURI)
		http.Redirect(w, r, addr+r.RequestURI, 307)
	}
}
