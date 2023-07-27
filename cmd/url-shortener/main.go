package main

import (
	"fmt"
	"shortly/app/entities"
	"shortly/app/repositories"
)

func main() {
	urlRepository := repositories.NewURLRepository()

	url := entities.NewURL("https://www.google.com", "abc")
	urlRepository.Store(url)

	url, found := urlRepository.Find("abc")
	if found {
		fmt.Printf("URL: %s\n", url.GetOriginalURL())
	} else {
		fmt.Println("URL not found")
	}
}