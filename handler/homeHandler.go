package handler

import (
	"encoding/json"
	"io"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ShowErrorPage(w, "Page not found", http.StatusNotFound)
		return
	}
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		ShowErrorPage(w, "Error fetching data", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		ShowErrorPage(w, "Error reading data", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(data, &ApiObjects)
	if err != nil {
		ShowErrorPage(w, "Error unmarshalling data", http.StatusInternalServerError)
		return
	}

	err = templates.ExecuteTemplate(w, "hpage.html", ApiObjects)
	if err != nil {
		ShowErrorPage(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
