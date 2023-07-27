package entities

type URL struct {
	OriginalURL string
	ShortURL 	string
}

func NewURL(originalURL string, shortURL string) *URL {
	return &URL{
		OriginalURL: originalURL,
		ShortURL: shortURL,
	}
}

func (u *URL) GetOriginalURL() string {
	return u.OriginalURL
}

func (u *URL) GetShortURL() string {
	return u.ShortURL
}