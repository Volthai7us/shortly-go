package services

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"shortly/app/entities"
	"shortly/app/repositories"
)

type URLShortenerService struct {
	urlRepository *repositories.URLRepository
}

func NewURLShortenerService(urlRepository *repositories.URLRepository) *URLShortenerService {
	return &URLShortenerService{
		urlRepository: urlRepository,
	}
}

func (s *URLShortenerService) Shorten(originalURL string) (string, error) {
	url, found := s.urlRepository.FindByOriginalURL(originalURL)
	if found {
		return url.GetShortURL(), nil
	}

	shortURL := s.generateShortURL(originalURL)
	fmt.Println("shortURL: ", shortURL)
	url_struct := entities.NewURL(originalURL, shortURL)
	s.urlRepository.Store(url_struct)

	return shortURL, nil
}

func (s *URLShortenerService) Find(shortURL string) (string, error) {
	url, found := s.urlRepository.Find(shortURL)
	if found {
		return url.OriginalURL, nil
	}
	return "", errors.New("URL not found")
}

func (s *URLShortenerService) NumberOfURLs() int {
	return s.urlRepository.Count()
}

func (s *URLShortenerService) generateShortURL(originalURL string) string {
	randomString := sha256.Sum256([]byte(originalURL))
	hexString := fmt.Sprintf("%x", randomString)
	return hexString[0:8]
}