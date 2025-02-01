package in_memory

import (
	"fmt"
)

func (r *Repo) FindByShortUrl(shortUrl string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	longURL, exists := r.data[shortUrl]
	if !exists {
		return "", fmt.Errorf("hash %s not found", shortUrl)
	}

	return longURL, nil
}
