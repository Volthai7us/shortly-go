package services

import (
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

func (s *URLShortenerService) Shorten(originalURL string) string {
	url, found := s.urlRepository.Find(originalURL)
	if found {
		return url.ShortURL
	}

	shortURL := s.generateShortURL(originalURL)
	url_struct := entities.NewURL(originalURL, shortURL)
	s.urlRepository.Store(url_struct)

	return shortURL
}

func (s *URLShortenerService) Find(shortURL string) string {
	url, found := s.urlRepository.Find(shortURL)
	if found {
		return url.OriginalURL
	}
	return ""
}

func (s *URLShortenerService) NumberOfURLs() int {
	return s.urlRepository.Count()
}

func (s *URLShortenerService) generateShortURL(originalURL string) string {
	return originalURL[0:6]
}