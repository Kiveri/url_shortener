package http_controller

import (
	"url/internal/domain/model"
	"url/internal/usecase"
)

type (
	HttpController struct {
		urlsUseCase urlsUseCase
		baseUrl     string
	}

	urlsUseCase interface {
		CreateUrl(req usecase.CreateUrlReq) (*model.URL, error)
		FindUrl(req usecase.FindUrlReq) (*model.URL, error)
	}
)

func NewController(urlsUseCase urlsUseCase, baseUrl string) *HttpController {
	return &HttpController{
		urlsUseCase: urlsUseCase,
		baseUrl:     baseUrl,
	}
}
