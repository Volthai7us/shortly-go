package controllers

import (
	"net/http"
	"shortly/app/services"
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
