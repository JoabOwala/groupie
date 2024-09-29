package models

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Image        string   `json:"image"`
}

// ConcertDetails for a specific artist
type ConcertDetails struct {
	Dates []string `json:"dates"`
}

// LocationDetails for a specific artist
type LocationDetails struct {
	Locations []string `json:"locations"`
}

// RelationData contains information about the relations between artists
type RelationData struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// If the response is an array, use a slice instead of a map
type Relations struct {
	RelationsList []RelationData `json:"index"` // Adjust field name if needed
}
