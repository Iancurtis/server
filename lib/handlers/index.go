package handlers

import "net/http"

//RedirIndex redirect index to other url
func RedirIndex(w http.ResponseWriter, r *http.Request) {
	// http.Redirect(w, r, "/pages", http.StatusMovedPermanently)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}
