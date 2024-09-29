package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"groupie-tracker/handlers"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run .")
		return
	}

	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))

	artistIDRegex := regexp.MustCompile(`^/artists/(\d+)$`)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			handlers.HomeHandler(w, r)
		} else if artistIDRegex.MatchString(r.URL.Path) {
			// Extract the ID from the URL
			matches := artistIDRegex.FindStringSubmatch(r.URL.Path)
			if len(matches) > 1 {
				idStr := matches[1]
				id, err := strconv.Atoi(idStr)
				if err == nil && id >= 1 && id <= 52 {
					handlers.ArtistDetails(w, r)
					return
				} else {
					w.WriteHeader(404)
					http.ServeFile(w, r, "templates/404.html")
				}
			}
		} else {
			w.WriteHeader(404)
			http.ServeFile(w, r, "templates/404.html")
		}
	})

	port := ":8000"
	log.Printf("Server started on http://localhost%s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
