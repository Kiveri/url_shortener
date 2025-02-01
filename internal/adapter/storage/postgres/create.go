package postgres

import (
	"context"
	"fmt"
)

func (r *Repo) CreateShortUrl(shortUrl, longURL string) (string, error) {
	query := `INSERT INTO urls (short_url, long_url) VALUES ($1, $2) ON CONFLICT (short_url) DO NOTHING;`

	_, err := r.cluster.Conn.Exec(context.Background(), query, shortUrl, longURL)
	if err != nil {
		return "", fmt.Errorf("error inserting hash and long_url: %v", err)
	}

	return shortUrl, nil
}
