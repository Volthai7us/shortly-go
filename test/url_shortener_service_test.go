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

	// Test shorten another URL
	anotherURL := "https://www.example.com"
	anotherShortURL, err := service.Shorten(anotherURL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if anotherShortURL == "" {
		t.Errorf("Expected shortened URL, got an empty string")
	}

	// Assert that the second URL was stored correctly
	url, found = repo.Find(anotherShortURL)
	if !found {
		t.Errorf("URL %s was not stored correctly", anotherShortURL)
	}

	if url.GetOriginalURL() != anotherURL {
		t.Errorf("Expected original URL to be %s, got %s", anotherURL, url.GetOriginalURL())
	}
}

func TestFind(t *testing.T) {
	// Setup
	repo := repositories.NewURLRepository()
	service := services.NewURLShortenerService(repo)

	originalURL := "https://www.test.com"
	shortURL, _ := service.Shorten(originalURL)

	// Test finding the original URL by its short version
	foundURL, err := service.Find(shortURL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if foundURL != originalURL {
		t.Errorf("Expected original URL to be %s, got %s", originalURL, foundURL)
	}

	// Test finding the original URL by a non-existed short version
	_, err = service.Find("non-existed")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestNumberOfURLs(t *testing.T) {
	// Setup
	repo := repositories.NewURLRepository()
	service := services.NewURLShortenerService(repo)

	// Test number of URLs when the repository is empty
	num := service.NumberOfURLs()
	if num != 0 {
		t.Errorf("Expected number of URLs to be 0, got %d", num)
	}

	// Test number of URLs after storing a new one
	service.Shorten("https://www.test.com")
	num = service.NumberOfURLs()
	if num != 1 {
		t.Errorf("Expected number of URLs to be 1, got %d", num)
	}

	for i := 0; i < 10; i++ {
		service.Shorten("https://www.test.com")
	}

	num = service.NumberOfURLs()
	
	if num != 1 {
		t.Errorf("Expected number of URLs to be 1, got %d", num)
	}

	for i := 0; i < 10; i++ {
		service.Shorten("https://www.test.com " + string(rune(i)))
	}

	num = service.NumberOfURLs()

	if num != 11 {
		t.Errorf("Expected number of URLs to be 11, got %d", num)
	}

	
}