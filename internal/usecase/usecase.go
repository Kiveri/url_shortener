package usecase

import (
	"url/internal/domain/model"
)

type urlRepo interface {
	CreateUrl(url *model.URL) (*model.URL, error)
	FindUrl(shortUrl string) (*model.URL, error)
}

type UrlUseCase struct {
	urlRepo urlRepo
}

func NewUseCase(urlRepo urlRepo) *UrlUseCase {
	return &UrlUseCase{
		urlRepo: urlRepo,
	}
}
