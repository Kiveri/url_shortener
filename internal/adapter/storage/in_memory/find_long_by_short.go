package in_memory

import (
	"url/internal/adapter/storage"
	"url/internal/domain/model"
)

func (r *Repo) FindByShortUrl(shortUrl string) (*model.URL, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	longUrl, exists := r.urls[shortUrl]
	if !exists {
		return nil, storage.NotFound
	}
	url := model.NewURL(shortUrl, longUrl)

	return url, nil
}
