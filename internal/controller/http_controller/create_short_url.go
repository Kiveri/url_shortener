package http_controller

import (
	"encoding/json"
	"net/http"

	"url/internal/controller"
	"url/internal/usecase"
)

type createShortUrlRequestHttp struct {
	LongUrl string `json:"long_url"`
}

type createShortUrlResponseHttp struct {
	ShortUrl string `json:"short_url"`
}

func (c *HttpController) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req createShortUrlRequestHttp
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

	Respond(w, http.StatusOK, createShortUrlResponseHttp{ShortUrl: c.baseUrl + shortUrl.ShortUrl})
}

func validateCreateShortUrlRequest(req createShortUrlRequestHttp) *controller.ValidationError {
	if req.LongUrl == "" {
		return controller.NewValidationError("url is required", "long_url")
	}

	return nil
}
