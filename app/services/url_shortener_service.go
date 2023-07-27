package services

import (
	"errors"
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
	url, found := s.urlRepository.Find(originalURL)
	if found {
		return url.ShortURL, errors.New("URL already exists")
	}

	shortURL := s.generateShortURL(originalURL)
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
	return originalURL[0:6]
}