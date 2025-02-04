package postgres

import (
	"context"
	"errors"
	"fmt"

	"url/internal/adapter/storage"
	"url/internal/domain/model"

	"github.com/jackc/pgx/v5"
)

func (r *Repo) FindUrl(shortUrl string) (*model.URL, error) {
	var url model.URL

	query := `SELECT long_url FROM urls WHERE short_url = $1`

	err := r.cluster.Conn.QueryRow(context.Background(), query,
		shortUrl,
	).Scan(&url.LongUrl)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.NotFound
		}

		return nil, fmt.Errorf("FindUrl: %w", err)
	}

	return &url, nil
}
