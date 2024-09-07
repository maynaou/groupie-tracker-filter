package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie_tracker/handler"
)

func main() {
	darkBlue := "\033[38;5;17m" // Dark Blue (#0A192F)
	cyan := "\033[38;5;39m"     // Cyan (#64ffda)
	reset := "\033[0m"
	err1 := handler.FetchArtistsAndRelationsData()
	if err1 != nil {
		log.Fatalf("Error fetching data: %v", err1)
	}
	http.HandleFunc("/static", handler.StaticHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/artist", handler.ArtistHandler)
	http.HandleFunc("/details", handler.DetailsHandler)

	fmt.Println(darkBlue + "/-------------------------------------------------------------------\\" + reset)
	fmt.Println(darkBlue + "|" + cyan + "                Server started on http://localhost:7000            " + darkBlue + " |" + reset)
	fmt.Println(darkBlue + "\\-------------------------------------------------------------------/" + reset)

	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
