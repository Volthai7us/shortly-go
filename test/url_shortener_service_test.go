package test

import (
	"shortly/app/repositories"
	"shortly/app/services"
	"testing"
)

func TestShorten(t *testing.T) {
	// Setup
	repo := repositories.NewURLRepository()
	service := services.NewURLShortenerService(repo)

	originalURL := "https://www.test.com"
	shortURL, err := service.Shorten(originalURL)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if shortURL == "" {
		t.Errorf("Expected shortened URL, got an empty string")
	}

	// Assert that the URL was stored correctly
	url, found := repo.Find(shortURL)
	if !found {
		t.Errorf("URL %s was not stored correctly", shortURL)
	}

	if url.GetOriginalURL() != originalURL {
		t.Errorf("Expected original URL to be %s, got %s", originalURL, url.GetOriginalURL())
	}
}
