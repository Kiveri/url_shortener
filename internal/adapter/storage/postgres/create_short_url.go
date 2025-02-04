package postgres

import (
	"context"
	"fmt"

	"url/internal/domain/model"
)

func (r *Repo) CreateUrl(url *model.URL) (*model.URL, error) {
	query := `INSERT INTO urls (short_url, long_url) VALUES ($1, $2) ON CONFLICT (short_url) DO NOTHING;`

	_, err := r.cluster.Conn.Exec(context.Background(), query,
		url.ShortUrl,
		url.LongUrl,
	)

	if err != nil {
		return nil, fmt.Errorf("error inserting hash and long_url: %v", err)
	}

	return url, nil
}
