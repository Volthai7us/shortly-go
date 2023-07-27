package controllers

import (
	"net/http"
	"shortly/app/services"
	"strings"
)

type URLController struct {
	service *services.URLShortenerService
}

func NewURLController(service *services.URLShortenerService) *URLController {
	return &URLController{
		service: service,
	}
}

// simple hello world
func (c *URLController) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

// create a new short url
func (c *URLController) Create(w http.ResponseWriter, r *http.Request) {
	//TODO: implement
}

// redirect to the original url
func (c *URLController) Redirect(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	parts := strings.Split(url, "/")	
	if len(parts) < 2 {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		path := parts[len(parts)-1]
		http.Redirect(w, r, "https://www.google.com/search?q="+path, http.StatusFound)
	}
	//TODO: implement
}
