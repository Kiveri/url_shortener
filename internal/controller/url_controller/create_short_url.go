package url_controller

import (
	"encoding/json"
	"net/http"

	"url/internal/controller"
	"url/internal/usecase"
)

type createShortUrlRequest struct {
	LongUrl string `json:"long_url"`
}

func (c *Controller) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req createShortUrlRequest
	err := decoder.Decode(&req)
	if err != nil {
		controller.InternalServer(w, err)

		return
	}

	if validationError := validateCreateShortUrlRequest(req); validationError != nil {
		controller.ValidationErrorRespond(w, validationError)

		return
	}

	shortUrl, err := c.urlsUseCase.CreateShortUrl(usecase.CreateShortURLReq{
		LongURL: req.LongUrl,
	})
	if err != nil {
		controller.InternalServer(w, err)

		return
	}

	controller.Validation(w, http.StatusOK, shortUrl)
}

func validateCreateShortUrlRequest(req createShortUrlRequest) *controller.ValidationError {
	if req.LongUrl == "" {
		return controller.NewValidationError("url is required", "long_url")
	}

	return nil
}
