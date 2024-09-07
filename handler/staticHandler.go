package handler

import "net/http"

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/static" {
		ShowErrorPage(w, "Page not found", http.StatusNotFound)
		return
	}
}
