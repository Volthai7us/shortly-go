package services

import (
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
	//TODO: implement this method
	return originalURL[0:6]
}