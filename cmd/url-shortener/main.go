package main

import (
	"log"
	"net/http"
	"shortly/app/controllers"
	"shortly/app/repositories"
	"shortly/app/services"
	"time"
)

func main() {
	urlRepository, err := repositories.NewURLRepository("/data/urls.json")
	if err != nil {
		log.Fatalf("Failed to create URL repository: %v", err)
	}

	go func() {
		for {
			time.Sleep(10 * time.Second) // Or whatever interval makes sense
			err := urlRepository.Sync()
			if err != nil {
				log.Println(err) // handle error
			}
		}
	}()

	urlService := services.NewURLShortenerService(urlRepository)
	urlController := controllers.NewURLController(urlService)

	http.HandleFunc("/hello", urlController.Index)
	http.HandleFunc("/create", urlController.Create)
	http.HandleFunc("/urls", urlController.Urls)
	http.HandleFunc("/", urlController.Redirect)

	http.ListenAndServe(":5173", nil)
}
