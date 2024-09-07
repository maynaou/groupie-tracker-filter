package handler

import "net/http"

func ShowErrorPage(w http.ResponseWriter, errMsg string, statusCode int) {
	data := struct {
		Err string
	}{
		Err: errMsg,
	}
	w.WriteHeader(statusCode)
	if err := templates.ExecuteTemplate(w, "error.html", data); err != nil {
		http.Error(w, "Error displaying the error page", http.StatusInternalServerError)
	}
}
