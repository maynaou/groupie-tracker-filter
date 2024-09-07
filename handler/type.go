package handler

import "html/template"

type API struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	RelationsURL string   `json:"relations"` // URL for relations data
	Relation     Relation
}

type Relation struct {
	DatesLocations map[string][]string `json:"DatesLocations"`
}

var (
	templates  = template.Must(template.ParseGlob("templates/*.html"))
	ApiObjects []API
)
