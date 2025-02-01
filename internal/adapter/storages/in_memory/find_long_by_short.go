package in_memory

import (
	"fmt"
)

func (r *Repo) FindByShortUrl(shortUrl string) (string, error) {
	longURL, exists := r.data[shortUrl]
	if !exists {
		return "", fmt.Errorf("hash %s not found", shortUrl)
	}

	return longURL, nil
}
