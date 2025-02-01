package in_memory

func (r *Repo) CreateShortUrl(shortUrl, longUrl string) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[shortUrl]; ok {
		return shortUrl, nil
	} else {
		r.data[shortUrl] = longUrl
	}
	return shortUrl, nil
}
