package grpc_controller

import (
	"url/internal/domain/model"
	urlshortener "url/internal/pb/url_shortener"

	"url/internal/usecase"
)

type (
	Controller struct {
		urlshortener.UnimplementedUrlShortenerServer
		urlsUseCase urlsUseCase
		baseUrl     string
	}

	urlsUseCase interface {
		CreateUrl(req usecase.CreateUrlReq) (*model.URL, error)
		FindUrl(req usecase.FindUrlReq) (*model.URL, error)
	}
)

func NewController(urlsUseCase urlsUseCase, baseUrl string) *Controller {
	return &Controller{
		urlsUseCase: urlsUseCase,
		baseUrl:     baseUrl,
	}
}
