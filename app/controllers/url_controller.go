package controllers

import (
	"fmt"
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
	url := r.URL.Query().Get("url")
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	shortURL, _ := c.service.Shorten(url)
	w.Write([]byte(shortURL))

	fmt.Println(c.service.NumberOfURLs())
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
		originalURL, err := c.service.Find(path)
		fmt.Println(originalURL)
		if err == nil {
			if !strings.HasPrefix(originalURL, "http") {
				originalURL = "http://" + originalURL
			}

			http.Redirect(w, r, originalURL, http.StatusFound)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
