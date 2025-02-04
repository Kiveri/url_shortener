package usecase

import (
	"fmt"

	"url/internal/domain/model"
)

type FindUrlReq struct {
	ShortUrl string
}

func (u *UrlUseCase) FindUrl(req FindUrlReq) (*model.URL, error) {
	url, err := u.urlRepo.FindUrl(req.ShortUrl)
	if err != nil {
		return nil, fmt.Errorf("urlRepo.FindByHash: %w", err)
	}
	return url, nil
}
