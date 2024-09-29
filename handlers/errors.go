package handlers

import (
	"html/template"
	"net/http"
)

// handle 404 error
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/404.html")
	if err != nil {
		http.Error(w, "404 Page Not Found ", http.StatusNotFound)
		return
	}
	tpl.Execute(w, nil)
}

// handle 405 error
func FourohFiveHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/405.html")
	if err != nil {
		http.Error(w, "405 Bad Method ", http.StatusMethodNotAllowed)
		return
	}
	tpl.Execute(w, nil)
}

// handle 400 error
func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/400.html")
	if err != nil {
		http.Error(w, "400 Bad Request ", http.StatusBadRequest)
		return
	}
	tpl.Execute(w, nil)
}

// handle 500 error
func InternalServerHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/500.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error ", http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, nil)
}
