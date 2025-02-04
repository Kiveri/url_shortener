package in_memory

import (
	"url/internal/domain/model"
)

func (r *Repo) CreateUrl(url *model.URL) (*model.URL, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.urls[url.ShortUrl]; ok {
		return url, nil
	} else {
		r.urls[url.ShortUrl] = url.LongUrl
	}

	return url, nil
}
