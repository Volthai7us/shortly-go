package repositories

import "shortly/app/entities"

// in-memory repository
type URLRepository struct {
    urls map[string]*entities.URL
}

// new
func NewURLRepository() *URLRepository {
    return &URLRepository{
        urls: make(map[string]*entities.URL),
    }
}

// store
func (r *URLRepository) Store(url *entities.URL) {
    r.urls[url.GetShortURL()] = url
}

// find
func (r *URLRepository) Find(shortURL string) (*entities.URL, bool) {
    url, found := r.urls[shortURL]
    return url, found
}

// find by original url
func (r *URLRepository) FindByOriginalURL(originalURL string) (*entities.URL, bool) {
    for _, url := range r.urls {
        if url.GetOriginalURL() == originalURL {
            return url, true
        }
    }
    return nil, false
}

// count
func (r *URLRepository) Count() int {
    return len(r.urls)
}
