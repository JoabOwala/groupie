package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/data"
	"groupie-tracker/models"
)

// ArtistDetailsHandler fetches and displays a single artist's detail
func ArtistDetails(w http.ResponseWriter, r *http.Request) {
	// Extract the artistID from the URL
	path := r.URL.Path
	idStr := strings.TrimPrefix(path, "/artists/") // Extract the ID part from the URL
	artistID, err := strconv.Atoi(idStr)           // Convert ID from string to integer
	if err != nil {
		log.Println("Invalid artist ID:", err)
		BadRequestHandler(w, r) // Ensure you have this handler defined
		return
	}

	// Fetch the artist's details, concert dates, locations, and relations
	artist, dates, locations, artistRelations, err := data.GetArtistByID(artistID)
	if err != nil {
		log.Println("Error fetching artist details:", err)
		InternalServerHandler(w, r) // Ensure you have this handler defined
		return
	}

	// Combine dates and locations into a single data structure
	combinedDetails := make([][]string, len(dates))
	for i := range dates {
		// Handle cases where dates or locations might be empty
		loc := ""
		if i < len(locations) {
			loc = locations[i]
		}
		combinedDetails[i] = []string{dates[i], loc}
	}

	// Pass artist, concert details, and relations to the template
	tmplData := struct {
		Artist          *models.Artist
		ConcertDetails  [][]string
		ArtistRelations map[string][][]string
	}{
		Artist:          artist,
		ConcertDetails:  combinedDetails,
		ArtistRelations: artistRelations,
	}

	tmpl := template.Must(template.ParseFiles("templates/artist_details.html"))
	if err := tmpl.Execute(w, tmplData); err != nil {
		log.Println("Error rendering template:", err)
		InternalServerHandler(w, r) // Ensure you have this handler defined
	}
}
