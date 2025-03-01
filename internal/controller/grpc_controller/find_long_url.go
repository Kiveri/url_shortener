package grpc_controller

import (
	"context"
	"fmt"

	urlshortener "url/internal/pb/url_shortener"
	"url/internal/usecase"
)

func (c *Controller) GetLink(ctx context.Context, req *urlshortener.ResolveRequest) (*urlshortener.ResolveResponse, error) {
	longUrl, err := c.urlsUseCase.FindUrl(usecase.FindUrlReq{
		ShortUrl: req.ShortUrl,
	})
	if err != nil {
		return nil, fmt.Errorf("FindLongUrl err: %w", err)
	}

	return &urlshortener.ResolveResponse{LongUrl: longUrl.LongUrl}, nil
}
