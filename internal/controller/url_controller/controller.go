package url_controller

import (
	"url/internal/usecase"
)

type (
	Controller struct {
		urlsUseCase urlsUseCase
	}

	urlsUseCase interface {
		CreateShortUrl(req usecase.CreateShortURLReq) (string, error)
		FindByShortUrl(req usecase.FindLongReq) (string, error)
	}
)

func NewController(urlsUseCase urlsUseCase) *Controller {
	return &Controller{
		urlsUseCase: urlsUseCase,
	}
}
