package usecase

import (
	"fmt"

	"url/internal/domain/model"
)

type FindLongReq struct {
	ShortUrl string
}

func (u *UrlUseCase) FindByShortUrl(req FindLongReq) (*model.URL, error) {
	//prefix := u.domainUrl
	//hash := req.ShortUrl[len(prefix):]

	url, err := u.urlRepo.FindByShortUrl(req.ShortUrl)
	if err != nil {
		return nil, fmt.Errorf("urlRepo.FindByHash: %w", err)
	}
	return url, nil
}
