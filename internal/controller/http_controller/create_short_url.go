package http_controller

import (
	"encoding/json"
	"net/http"

	"url/internal/usecase"
)

type createShortUrlRequest struct {
	LongUrl string `json:"long_url"`
}

type createShortUrlResponse struct {
	ShortUrl string `json:"short_url"`
}

func (c *Controller) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req createShortUrlRequest
	err := decoder.Decode(&req)
	if err != nil {
		InternalServerErrorRespond(w, err)

		return
	}

	if validationError := validateCreateShortUrlRequest(req); validationError != nil {
		ValidationErrorRespond(w, validationError)

		return
	}

	shortUrl, err := c.urlsUseCase.CreateUrl(usecase.CreateUrlReq{
		LongURL: req.LongUrl,
	})
	if err != nil {
		InternalServerErrorRespond(w, err)

		return
	}

	Respond(w, http.StatusOK, createShortUrlResponse{ShortUrl: c.baseUrl + shortUrl.ShortUrl})
}

func validateCreateShortUrlRequest(req createShortUrlRequest) *ValidationError {
	if req.LongUrl == "" {
		return NewValidationError("url is required", "long_url")
	}

	return nil
}
