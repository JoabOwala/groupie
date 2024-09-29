package handlers

import (
	"html/template"
	"log"
	"net/http"

	"groupie-tracker/data"
	"groupie-tracker/models"
)

type AllArtist struct {
	MyData []models.Artist
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := data.GetArtists()
	temp, _ := template.ParseFiles("templates/500.html")
	if err != nil {
		log.Println("Error fetching artists:", err)
		temp.Execute(w, nil)
		// http.Error(w, "Unable to fetch artists", http.StatusInternalServerError)
		return
	}

	artistInfo := AllArtist{
		MyData: artists,
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err = tmpl.Execute(w, artistInfo)
	if err != nil {
		log.Println("Error rendering template:", err)
		// http.Error(w, "Unable to render template", http.StatusInternalServerError)
		return
	}
}
