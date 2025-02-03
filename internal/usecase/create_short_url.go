package usecase

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"url/internal/domain/model"
)

type CreateShortURLReq struct {
	LongURL string
}

func (u *UrlUseCase) CreateShortUrl(req CreateShortURLReq) (*model.URL, error) {
	var b strings.Builder
	var allowedChars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	hasher := sha256.New()
	hasher.Write([]byte(req.LongURL))
	hashBytes := hasher.Sum(nil)
	hashHex := hex.EncodeToString(hashBytes)

	for _, char := range hashHex {
		if strings.ContainsRune(allowedChars, char) {
			b.WriteRune(char)
		}
		if b.Len() < 10 {
			b.WriteRune('0')
		}
		if b.Len() == 10 {
			break
		}
	}

	hash := b.String()

	//shortUrl := hash + u.domainUrl

	url := model.NewURL(hash, req.LongURL)
	url, err := u.urlRepo.CreateShortUrl(url)
	if err != nil {
		return nil, fmt.Errorf("urlRepo.CreateShortUrl: %w", err)
	}

	return url, nil

	//hasher := sha256.Sum256([]byte(req.LongURL))
	//hash := base64.StdEncoding.EncodeToString(hasher[:])[:10]
	//hash = strings.ReplaceAll(hash, "+", "_")
	//hash = strings.ReplaceAll(hash, "/", "_")
	//hash = strings.ReplaceAll(hash, "=", "_")
	//var b strings.Builder
	//b.WriteString(hash)
	//b.WriteString(u.domainUrl)
	//shortUrl := b.String()
	//
	//url := model.NewURL(shortUrl, req.LongURL)
	//
	//url, err := u.urlRepo.CreateShortUrl(url)
	//if err != nil {
	//	return nil, fmt.Errorf("urlRepo.Save: %w", err)
	//}
	//return url, nil
}
