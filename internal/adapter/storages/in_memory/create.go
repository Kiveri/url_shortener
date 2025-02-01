package in_memory

func (r *Repo) CreateShortUrl(shortUrl, longUrl string) (string, error) {
	if _, ok := r.data[shortUrl]; ok {
		return shortUrl, nil
	} else {
		r.data[shortUrl] = longUrl
	}
	return shortUrl, nil
}
