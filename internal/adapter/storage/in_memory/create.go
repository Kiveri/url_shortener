package in_memory

func (r *Repo) CreateShortUrl(shortUrl, longUrl string) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.urls[shortUrl]; ok {
		return shortUrl, nil
	} else {
		r.urls[shortUrl] = longUrl
	}

	return shortUrl, nil
}
