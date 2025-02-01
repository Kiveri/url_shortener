package usecase

import (
	"fmt"
)

type FindLongReq struct {
	ShortUrl string
}

func (u *UrlUseCase) FindByShortUrl(req FindLongReq) (string, error) {
	longURL, err := u.urlRepo.FindByShortUrl(req.ShortUrl)
	if err != nil {
		return "", fmt.Errorf("urlRepo.FindByHash: %w", err)
	}
	return longURL, nil
}
