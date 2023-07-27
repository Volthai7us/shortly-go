package repositories

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"shortly/app/entities"
)

// in-memory repository
type URLRepository struct {
	urls map[string]*entities.URL
	filePath string
}

// new
func NewURLRepository(filePath string) (*URLRepository, error) {
	urls := make(map[string]*entities.URL)
	
	if(filePath != "") {
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			if os.IsNotExist(err) {
				err = ioutil.WriteFile(filePath, []byte("{}"), 0644)
				if err != nil {
					return nil, err
				}
			} else {
				return nil, err
			}
		} else {
			err = json.Unmarshal(data, &urls)
			if err != nil {
				return nil, err
			}
		}
	}

	return &URLRepository{
		urls: urls,
		filePath: filePath,
	}, nil
}

// store
func (r *URLRepository) Store(url *entities.URL) error {
	r.urls[url.GetShortURL()] = url

	data, err := json.Marshal(r.urls)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

//sync
func (r *URLRepository) Sync() error {
	if r.filePath == "" {
		return nil
	}
	var urls map[string]*entities.URL

	data, err := ioutil.ReadFile(r.filePath)
	if err != nil {
		return err
	}
	
	err = json.Unmarshal(data, &urls)
	if err != nil {
		return err
	}

	r.urls = urls
	return nil
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

// all
func (r *URLRepository) All() []*entities.URL {
	urls := make([]*entities.URL, 0, len(r.urls))
	for _, url := range r.urls {
		urls = append(urls, url)
	}
	return urls
}
