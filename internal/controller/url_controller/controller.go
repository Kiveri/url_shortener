package url_controller

import (
	"url/internal/domain/model"
	"url/internal/usecase"
)

type (
	Controller struct {
		urlsUseCase urlsUseCase
		baseUrl     string
	}

	urlsUseCase interface {
		CreateShortUrl(req usecase.CreateShortURLReq) (*model.URL, error)
		FindByShortUrl(req usecase.FindLongReq) (*model.URL, error)
	}
)

func NewController(urlsUseCase urlsUseCase, baseUrl string) *Controller {
	return &Controller{
		urlsUseCase: urlsUseCase,
		baseUrl:     baseUrl,
	}
}
