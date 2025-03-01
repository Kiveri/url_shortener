package grpc_controller

import (
	"context"
	"fmt"

	urlshortener "url/internal/pb/url_shortener"
	"url/internal/usecase"
)

func (c *Controller) ShortenLink(ctx context.Context, req *urlshortener.ShortenRequest) (*urlshortener.ShortenResponse, error) {
	shortUrl, err := c.urlsUseCase.CreateUrl(usecase.CreateUrlReq{
		LongURL: req.LongUrl,
	})
	if err != nil {
		return nil, fmt.Errorf("CreateShortUrl err: %w", err)
	}

	return &urlshortener.ShortenResponse{ShortUrl: c.baseUrl + shortUrl.ShortUrl}, nil
}
