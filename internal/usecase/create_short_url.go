package usecase

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

type CreateShortURLReq struct {
	LongURL string
}

func (u *UrlUseCase) CreateShortUrl(req CreateShortURLReq) (string, error) {
	hash := sha256.Sum256([]byte(req.LongURL))
	shortUrl := base64.StdEncoding.EncodeToString(hash[:])[:10]
	shortUrl = strings.ReplaceAll(shortUrl, "+", "_")
	shortUrl = strings.ReplaceAll(shortUrl, "/", "_")
	shortUrl = strings.ReplaceAll(shortUrl, "=", "_")

	_, err := u.urlRepo.CreateShortUrl(shortUrl, req.LongURL)
	if err != nil {
		return "", fmt.Errorf("urlRepo.Save: %w", err)
	}
	return shortUrl, nil
}
