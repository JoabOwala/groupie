package data

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupie-tracker/models"
)

// GetArtists fetches the list of artists from the API

func GetArtists() ([]models.Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []models.Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, err
	}

	return artists, nil
}

// GetArtistByID fetches a single artist's details by their ID from the API

func GetArtistByID(id int) (*models.Artist, []string, []string, map[string][][]string, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	defer resp.Body.Close()

	var artist models.Artist
	if err := json.NewDecoder(resp.Body).Decode(&artist); err != nil {
		return nil, nil, nil, nil, err
	}

	dates, locations, err := GetArtistConcertDetails(id)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	relations, err := GetArtistRelations()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	artistRelations := make(map[string][][]string)
	for _, relation := range relations {
		if relation.ID == id {
			// Flatten map[string][]string into [][]string for each artist
			var datesLocations [][]string
			for location, dates := range relation.DatesLocations {
				// Combine location and dates into a single slice
				for _, date := range dates {
					datesLocations = append(datesLocations, []string{location, date})
				}
			}
			// Convert the artist ID (int) to a string and add to artistRelations
			artistRelations[fmt.Sprintf("%d", relation.ID)] = datesLocations
		}
	}
	return &artist, dates, locations, artistRelations, nil
}

// GetArtistConcertDetails fetches concert dates and locations for the artist from the API
func GetArtistConcertDetails(artistID int) ([]string, []string, error) {
	// Fetch concert dates for the specific artist
	concertURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%d", artistID)
	datesResp, err := http.Get(concertURL)
	if err != nil {
		return nil, nil, err
	}
	defer datesResp.Body.Close()

	var concertDetails models.ConcertDetails
	if err := json.NewDecoder(datesResp.Body).Decode(&concertDetails); err != nil {
		return nil, nil, err
	}

	// Fetch locations for the specific artist
	locationURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%d", artistID)
	locationsResp, err := http.Get(locationURL)
	if err != nil {
		return nil, nil, err
	}
	defer locationsResp.Body.Close()

	var locationDetails models.LocationDetails
	if err := json.NewDecoder(locationsResp.Body).Decode(&locationDetails); err != nil {
		return nil, nil, err
	}

	// Ensure that the lengths of dates and locations match
	minLen := len(concertDetails.Dates)
	if len(locationDetails.Locations) < minLen {
		minLen = len(locationDetails.Locations)
	}

	dates := concertDetails.Dates[:minLen]
	locations := locationDetails.Locations[:minLen]

	return dates, locations, nil
}

// GetArtistRelations fetches the relations for the artists from the API

func GetArtistRelations() ([]models.RelationData, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var relations models.Relations
	if err := json.NewDecoder(resp.Body).Decode(&relations); err != nil {
		return nil, err
	}

	return relations.RelationsList, nil
}
