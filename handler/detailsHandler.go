package handler

import (
	"net/http"
	"strconv"
)

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ShowErrorPage(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the artist ID from the query parameters
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 || id > len(ApiObjects) {
		ShowErrorPage(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	// Fetch the artist details (relations already fetched and stored)
	artist := ApiObjects[id-1]

	// Render the details template using the already fetched artist and relations data
	err = templates.ExecuteTemplate(w, "details.html", map[string]interface{}{
		"Artist":    artist,
		"Relations": artist.Relation, // Use the stored Relation data
	})
	if err != nil {
		ShowErrorPage(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
