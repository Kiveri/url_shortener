package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repo) FindByShortUrl(shortUrl string) (string, error) {
	query := `SELECT long_url FROM urls WHERE short_url = $1`
	var longURL string

	err := r.cluster.Conn.QueryRow(context.Background(), query, shortUrl).Scan(&longURL)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", NotFound
		}

		return "", fmt.Errorf("FindClient: %w", err)
	}

	return longURL, nil
}
