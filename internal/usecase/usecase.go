package usecase

import (
	"url/internal/domain/model"
)

type urlRepo interface {
	CreateShortUrl(url *model.URL) (*model.URL, error)
	FindByShortUrl(shortUrl string) (*model.URL, error)
}

type UrlUseCase struct {
	urlRepo urlRepo
}

func NewUseCase(urlRepo urlRepo) *UrlUseCase {
	return &UrlUseCase{
		urlRepo: urlRepo,
	}
}
