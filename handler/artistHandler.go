package handler

import (
	"net/http"
	"strconv"
	"strings"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ShowErrorPage(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	creationStart := r.FormValue("creationStart")
	creationEnd := r.FormValue("creationEnd")
	albumStart := r.FormValue("albumStart")
	albumEnd := r.FormValue("albumEnd")
	members := r.FormValue("members")
	MaxMinMember := r.Form["members"]
	locationsearch := r.FormValue("locationSearch")

	var results []API
	query := r.URL.Query().Get("q")
	// if query != "" {
	results = searchArtists(ApiObjects, creationStart, creationEnd, albumStart, albumEnd, members, locationsearch, MaxMinMember)
	// } else {
	// 	results = ApiObjects
	// }

	err := templates.ExecuteTemplate(w, "artist.html", map[string]interface{}{
		"Artists": results,
		"Query":   query,
	})
	if err != nil {
		ShowErrorPage(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func searchArtists(data []API, creationStart, creationEnd, albumStart, albumEnd string, members string, locationsearch string, MaxMinMember []string) []API {
	var results []API
	creationStartInt, _ := strconv.Atoi(creationStart)
	creationEndInt, _ := strconv.Atoi(creationEnd)
	albumStartInt, _ := strconv.Atoi(albumStart)
	albumEndInt, _ := strconv.Atoi(albumEnd)
	membersInt, _ := strconv.Atoi(members)

	locations := strings.Split(locationsearch, ", ")
	locationsFinal := strings.Join(locations, "-")
	b := true
	for _, artist := range data {
		year := extractYear(artist.FirstAlbum)
		include := true
		// Filter based on creation date
		if creationStart != "" && creationEnd != "" {
			if !(creationStartInt <= artist.CreationDate && artist.CreationDate <= creationEndInt) {
				include = false
			}
		}

		// Filter based on album release year
		if albumStart != "" && albumEnd != "" {
			if !(albumStartInt <= year && year <= albumEndInt) {
				include = false
			}
		}

		if MaxMinMember != nil && len(MaxMinMember) != 1 {
			b = false
			for i := 0; i < len(MaxMinMember); i++ {
				nb, _ := strconv.Atoi(MaxMinMember[i])
				membersInt = nb
				if members != "" {
					if len(artist.Members) == membersInt && albumStartInt <= year && year <= albumEndInt {
						results = append(results, artist)
					} else if locationsFinal != "" {
						for location := range artist.Relation.DatesLocations {
							if len(artist.Members) == membersInt && strings.Contains(strings.ToLower(location), strings.ToLower(locationsFinal)) {
								results = append(results, artist)
							}
						}
					}
				}
			}
		}

		// Filter based on number of members
		if members != "" {
			if len(artist.Members) != membersInt {
				include = false
			}
		}

		// Filter based on location search
		if locationsearch != "" {
			matchFound := false
			for location := range artist.Relation.DatesLocations {
				if strings.Contains(strings.ToLower(location), strings.ToLower(locationsFinal)) {
					matchFound = true
					break
				}
			}
			if !matchFound {
				include = false
			}
		}
		// If the artist meets all the criteria, add them to results
		if include && b {
			results = append(results, artist)
		}
	}

	return results
}

func extractYear(album string) int {
	var words int
	parts := strings.Split(album, "-")
	if len(parts) == 3 {
		words, _ = strconv.Atoi(parts[2])
	}
	return words
}
