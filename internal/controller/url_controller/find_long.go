package url_controller

import (
	"errors"
	"net/http"

	"url/internal/adapter/storage"
	"url/internal/controller"
	"url/internal/usecase"
)

type findLongUrlResponse struct {
	LongUrl string `json:"long_url"`
}

func (c *Controller) FindLongUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Path[1:]
	longUrl, err := c.urlsUseCase.FindUrl(usecase.FindUrlReq{
		ShortUrl: shortUrl,
	})
	if err != nil {
		if errors.Is(err, storage.NotFound) {
			controller.NotFoundErrorRespond(w, controller.NewNotFoundError("not found"))

			return
		}

		controller.InternalServerErrorRespond(w, err)

		return
	}

	controller.Respond(w, http.StatusOK, findLongUrlResponse{LongUrl: longUrl.LongUrl})
}
