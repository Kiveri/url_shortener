package url_controller

import (
	"errors"
	"net/http"

	"url/internal/adapter/storage/postgres"
	"url/internal/controller"
	"url/internal/usecase"
)

func (c *Controller) FindLongUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.PathValue("short_url")
	longUrl, err := c.urlsUseCase.FindByShortUrl(usecase.FindLongReq{
		ShortUrl: shortUrl,
	})
	if err != nil {
		if errors.Is(err, postgres.NotFound) {
			controller.ValidationErrorRespond(w, controller.NewValidationError("long url not found", "short_url"))

			return
		}

		controller.InternalServer(w, err)

		return
	}

	controller.Validation(w, http.StatusOK, longUrl)
}
